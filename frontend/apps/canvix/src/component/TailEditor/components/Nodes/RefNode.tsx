import React from 'react';
import type { NodeProps } from 'tail-js';
import type { InteractionNodeTypeRuntime } from '@canvas/event-core';
import { Tooltip } from 'antd';
import { SyncOutlined } from '@ant-design/icons';
import NormalNode from './NormalNode';

export default abstract class RefNode extends NormalNode {
  getEndpoint = async (props: NodeProps<InteractionNodeTypeRuntime>) => {
    const { id } = props.node;
    const res = await this.context.getRefNodeEndpoints(id, true);
    this.setState({ ...res });
  };
  getScreenId = () => {
    const { id } = this.props.node;
    const com = this.context.findComponentById(id);
    return (com as any).attr?.screenId;
  };
  abstract refreshEndpoints(): Promise<void>;
  //  async() => {
  //   await this.getEndpoint(this.props);
  //   ````event```.emit('refresh-type-panel', null);
  // };

  getTitle = () => {
    if (this.state.isDelete) return this.getName();
    return '双击编辑引用面板';
  };

  renderIcon = () => {
    const { isDelete } = this.state;
    return !isDelete ? (
      <Tooltip title="刷新动作事件">
        <SyncOutlined
          className="refresh-icon"
          onClick={this.refreshEndpoints}
          onDoubleClick={(e) => {
            e.stopPropagation();
          }}
        />
      </Tooltip>
    ) : null;
  };
  handleDoubleClick = () => {
    const screenId = this.getScreenId();
    if (!screenId) return;
    window.open(`/${this.context.baseRoute}/${screenId}`);
  };
  componentDidMount() {
    this.getEndpoint(this.props);
  }
}
