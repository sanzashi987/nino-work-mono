import React, { PureComponent, FC } from 'react';
import { Handle } from 'tail-js';
import { EndpointType } from '@canvix/component-factory';
import Style from './index.module.scss';
// import { SourcePoint, TargetPoint } from '../Point';

type EndpointListProps = {
  nodeId: string;
  endpoints: EndpointType[];
};

const { 'child-list': childListStyle } = Style;

type SourceTarget = 'source' | 'target';

class EndpointList extends PureComponent<EndpointListProps> {
  type: SourceTarget = 'source';

  render() {
    const { nodeId, endpoints } = this.props;
    return (
      <ul className="endpoint-list">
        {endpoints.map(({ id, name, describer }) => (
          <li key={id} className="endpoint-item">
            <span className={`desc ${this.type}`} title={name}>
              {name}
            </span>
            <div className="handle-outer">
              <Handle type={this.type} nodeId={nodeId} handleId={id} describer={describer} />
            </div>
          </li>
        ))}
      </ul>
    );
  }
}

class SourceList extends EndpointList {
  type = 'source' as const;
}
class TargetList extends EndpointList {
  type = 'target' as const;
}

type ChildListProps = {
  childList: string[];
};

const ChildList: FC<ChildListProps> = ({ childList }) => (
  <ul className={childListStyle}>
    {childList.map((e, i) => (
      <li className="child-item" title={e} key={e}>
        {e}
      </li>
    ))}
  </ul>
);

export { SourceList, TargetList, ChildList };
