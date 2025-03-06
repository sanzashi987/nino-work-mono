import React, { Component, createRef } from 'react';
import Tail, {
  coordinates,
  MiniMap,
  CoreMethods,
  SelectModeType,
  EdgeTree,
  EdgeBasic,
  Node,
  EdgeAtomState,
  NodeAtomState,
  getConnectedEdgeByNode,
  NodeTemplatesType,
  HandleAttribute,
  EdgePairedResult,
} from 'tail-js';
import { createMemo } from '@canvas/utilities';
import type { InteractionConfigType } from '@canvas/event-core';
import { getCopyEnable, getItemSwitchable } from './helpers';
import { Arrows, ArrowTemplates } from './Arrows';
import { BezierConnectingEdge, EdgeTemplate } from './Edges';
import { NodeTemplatePicker } from './Nodes';

type NodeRemovalOption = {
  id: string;
  edges: string[];
}[];
type State = { scale: number };
export type Props = {
  activeNodes: string[];
  activeEdges: string[];
  // panelStatus: string;
} & InteractionConfigType;

const MiniMapStatic = (
  <MiniMap
    width={200}
    height={140}
    activeColor="var(--canvas-primary-color)"
    nodeColor="var(--canvas-ui-emphasis-bgcolor)"
    viewportFrameColor="var(--canvas-primary-color)"
    style={{
      background: 'var(--canvas-ui-lvl1-bgcolor)',
    }}
  />
);

abstract class TailEditor<T extends Props> extends Component<T, State> {
  ref = createRef<HTMLDivElement>();
  tail = createRef<CoreMethods>();
  state = { scale: 1 };
  mode: SelectModeType = SelectModeType.single;

  memoSwitch = createMemo(getItemSwitchable);
  memoCopy = createMemo(getCopyEnable);

  setScale = (scale: number) => {
    this.setState({ scale });
  };

  setPanelZoom = (scale: number) => {
    this.setScale(scale);
    this.tail.current?.setScale(scale);
  };

  focusNode = (id: string) => this.tail.current?.focusNode(id);
  resetViewCenter = () => this.tail.current?.moveViewCenter(0, 0);
  switchSelectionMode = () => {
    this.mode = this.mode === SelectModeType.single ? SelectModeType.select : SelectModeType.single;
    return this.tail.current?.switchMode(this.mode);
  };

  onEdgeRightClick = (e: React.MouseEvent, edge: EdgeAtomState) => {
    const { activeEdges } = this.props;
    if (!activeEdges.includes(edge.edge.id)) return;
    this.onContextMenu(e, edge.edge.id);
  };

  onNodeRightClick = (e: React.MouseEvent, node: NodeAtomState) => {
    const { activeNodes } = this.props;
    if (!activeNodes.includes(node.node.id)) return;
    this.onContextMenu(e, node.node.id);
  };

  // abstract findReleativesByNodeIds(nodes: string[]): void;
  findReleativesByNodeIds(nodes: string[]) {
    const tree = this.tail.current?.getEdgeTree() ?? (new Map() as EdgeTree);
    return nodes.reduce<NodeRemovalOption>((last, curr) => {
      const payload = { id: curr, edges: [] as string[] };
      last.push(payload);
      const allHandle = tree.get(curr)?.values();
      if (!allHandle) return last;
      payload.edges = getConnectedEdgeByNode(curr, tree);
      return last;
    }, []);
  }

  abstract getNodeTemplates(): NodeTemplatesType;
  abstract onDrop(e: React.DragEvent, offset: coordinates, scale: number): void;
  abstract onEdgeUpdate(id: string, edge: EdgeBasic, pairedStatus: EdgePairedResult | null): void;
  abstract onNodesUpdate(nodes: Node[]): void;
  abstract onEdgeCreate(edge: EdgeBasic, pairedStatus: EdgePairedResult | null): void;
  abstract onEdgePaired(
    sourceHandle: HandleAttribute,
    targetHandle: HandleAttribute,
  ): EdgePairedResult;

  abstract onNodeClick(e: MouseEvent, s: NodeAtomState): void;
  abstract onEdgeClick(e: React.MouseEvent, s: EdgeAtomState): void;
  /**
   * multiple selection `mouseup` event handler, recevies inside
   * node ids as the second param
   */
  abstract onSelect(e: MouseEvent, nodes: string[]): void;
  /**
   * when right clicking on the nodes or edges in the flow chart editor
   * this method will be called
   * @param e
   */
  abstract onContextMenu(e: React.MouseEvent, invokedBy: string): void;
  abstract resetActive(): void;

  render() {
    const { activeNodes, activeEdges, nodes, edges } = this.props;

    return (
      <div ref={this.ref} className="flow-editor-container">
        <Tail
          ref={this.tail}
          nodes={nodes}
          edges={edges}
          activeEdges={activeEdges}
          activeNodes={activeNodes}
          onViewerDrop={this.onDrop}
          nodeTemplatePicker={NodeTemplatePicker}
          nodeTemplates={this.getNodeTemplates()}
          edgeTemplates={EdgeTemplate}
          connectingEdge={BezierConnectingEdge}
          onNodeUpdate={this.onNodesUpdate}
          onEdgeCreate={this.onEdgeCreate}
          onEdgePaired={this.onEdgePaired}
          markerTemplates={ArrowTemplates}
          markers={Arrows}
          onEdgeUpdate={this.onEdgeUpdate}
          onNodeClick={this.onNodeClick}
          onEdgeClick={this.onEdgeClick}
          onSelect={this.onSelect}
          onEdgeContextMenu={this.onEdgeRightClick}
          onNodeContextMenu={this.onNodeRightClick}
          onViewerScale={this.setScale}
          onViewerClick={this.resetActive}
        >
          {MiniMapStatic}
        </Tail>
      </div>
    );
  }
}

export default TailEditor;
