import React from 'react';
import { ErrorOutline } from "@mui/icons-material";
import BasicNode, { renderVertical } from './BasicNode';
import Styles from './index.module.scss';

const { 'normal-node-wrap': NClass } = Styles;

class NormalNode<T extends {} = {}> extends BasicNode<T> {
  componentDidMount() {
    this.getEndpointList();
  }

  getName() {
    // return this.context.findComponentById(this.props.node.id)?.name || '已删除节点';
    return this.props.node.name || '已删除节点';
  }

  handleDoubleClick = () => {
    const { type, id } = this.props.node;
    const { isDelete } = this.state;
    if (isDelete) return;
    if (type === 'subpanel' || type === 'panel') {
      this.context.switchPanel(id);
    }
  };

  renderIcon = (): React.ReactNode => null;

  getTitle = () => this.getName();

  render() {
    if (!this.state.endpoints) return null;
    const {
      node: { id, disable },
      selected,
    } = this.props;
    const { deprecated, isDelete } = this.state;
    const { source, target } = this.state.endpoints;
    const title = this.getTitle();
    const cn_name = this.getName();
    return (
      <div
        className={`${NClass}  ${selected ? ' selected' : ''}`}
        onDoubleClick={this.handleDoubleClick}
        title={title}
      >
        <div className="node">
          <h2 className="title frnc">
            <i className="font"></i>
            {this.renderIcon()}
            <span className="title-text">{cn_name}</span>
          </h2>
          {!isDelete ? (
            <div className="body-container">{renderVertical(id, source, target)}</div>
          ) : (
            <div className="deleted-node"></div>
          )}
          {disable && <div className="disabled"></div>}
          {(deprecated || isDelete) && (
            <div className="deprecated">
              <ErrorOutline className="error"/>
            </div>
          )}
        </div>
      </div>
    );
  }
}

export default NormalNode;
