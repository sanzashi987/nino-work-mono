import React from 'react';
import type { InteractionNodeTypeRuntime } from '@canvas/event-core';
import BasicNode, { renderVertical } from './BasicNode';
import Styles from './index.module.scss';
import { SourceList, TargetList, ChildList } from './EndpointList';
import { TailEditorContext, TailEditorInterface } from '../../interface';

const { 'logical-node-wrap': LClass } = Styles;

const defaultColor = {
  foregroundColor: '#2d2e2f', //'var(--canvix-widget-darker-bgcolor)',
  backgroundColor: 'var(--canvix-ui-lvl1-bgcolor)',
};

class LogicalNode extends BasicNode {
  configMemo: InteractionNodeTypeRuntime | null = null;

  updateEndpointList() {
    const memoCf = this.configMemo;
    const cf = this.props.node;
    if (memoCf?.attr !== cf.attr || memoCf?.com !== cf.com) {
      this.configMemo = this.props.node;
      this.getEndpointList();
    }
  }

  renderWithMenuConfig = ({ menuPalette }: TailEditorInterface) => {
    const { source, target, childList = [] } = this.state.endpoints!;
    const verticalLayout = this.state.isVertical && childList.length === 0;
    const { node, selected } = this.props;
    const { name: cn_name, id, disable, com } = node;
    const { foregroundColor, backgroundColor } = menuPalette[com!.category] ?? defaultColor;
    return (
      <div
        className={`${LClass}  ${selected ? ' selected' : ''}`}
        style={{ background: backgroundColor }}
      >
        <h2 className="title" style={{ background: foregroundColor }}>
          <span className="title-text" title={cn_name}>
            {cn_name}
          </span>
        </h2>
        <div className={`body-container ${verticalLayout ? 'vertical' : ''}`}>
          {verticalLayout ? (
            renderVertical(id, source, target)
          ) : (
            <>
              <TargetList nodeId={id} endpoints={target} />
              {childList.length > 0 && <ChildList childList={childList} />}
              <SourceList nodeId={id} endpoints={source} />
            </>
          )}
        </div>
        {disable && <div className="disabled"></div>}
      </div>
    );
  };

  render() {
    this.updateEndpointList();
    if (!this.state.endpoints) return null;
    return <TailEditorContext.Consumer>{this.renderWithMenuConfig}</TailEditorContext.Consumer>;
  }
}

export default LogicalNode;
