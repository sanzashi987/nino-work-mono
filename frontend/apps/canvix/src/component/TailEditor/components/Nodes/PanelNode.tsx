import NormalNode from './NormalNode';

class SubpanelNode extends NormalNode {
  getTitle = () => {
    if (this.state.isDelete) return this.getName();
    return '双击进入配置动态面板交互';
  };
}

export default SubpanelNode;
