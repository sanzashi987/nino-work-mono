/* eslint-disable react/no-unused-state */
/* eslint-disable react/no-unused-class-component-methods */
import React, { Component } from 'react';
import type { NodeProps } from 'tail-js';
import { EndpointType } from '@canvix/component-factory';
import { InteractionNodeTypeRuntime } from '@canvix/event-core';
import { SourceList, TargetList } from './EndpointList';
import type { EndpointsStatusType, EndpointResType, TailEditorInterface } from '../../types';
import { TailEditorContext } from '../../interface';
// import { toEndpoints, EndpointsStatusType, EndpointResType } from '../../utils/generateNode';

type TemplateState = {
  endpoints: EndpointResType | null;
} & EndpointsStatusType;

class BasicNode<T extends object = object> extends Component<NodeProps<InteractionNodeTypeRuntime> & T> {
  static contextType?: React.Context<any> | undefined = TailEditorContext;

  declare context: TailEditorInterface;

  state: TemplateState = {
    endpoints: null,
    deprecated: false,
    isVertical: false
  };

  componentDidUpdate(lp: NodeProps<InteractionNodeTypeRuntime>, ls: TemplateState) {
    if (lp.selected !== this.props.selected) {
      this.props.setContainerStyle({ zIndex: this.props.selected ? 1 : 'unset' });
    }

    if (ls.endpoints !== this.state.endpoints) {
      this.props.updateNodeHandles();
    }
  }

  getEndpointList() {
    this.context
      .toEndpoints(this.props.node)
      .then((state) => {
        this.setState(() => ({ ...state }));
      })
      .catch((e) => {
        console.log(e);
      });
  }
}

export function renderVertical(id: string, source: EndpointType[], target: EndpointType[]) {
  return (
    <>
      <div className="endpoint-type">事件</div>
      {source?.length ? (
        <SourceList nodeId={id} endpoints={source} />
      ) : (
        <div className="empty-placeholder">暂无事件</div>
      )}
      <div className="endpoint-type">动作</div>
      {target?.length ? (
        <TargetList nodeId={id} endpoints={target} />
      ) : (
        <div className="empty-placeholder">暂无动作</div>
      )}
    </>
  );
}

export default BasicNode;
