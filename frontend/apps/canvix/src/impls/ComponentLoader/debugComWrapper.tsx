import React, { Component, ComponentType } from 'react';
import { eventHub, getComHubEvent } from '@app/statics';
import { uuid } from '@/utils';

type DebugWrapperState = {
  Block: ComponentType | null;
};
const WidthDebugWrapper = (Com: ComponentType) => class DebugWrapper extends Component<any, DebugWrapperState> {
  codeResponse: string;

  constructor(props: any) {
    super(props);
    this.state = { Block: Com };
    const { config } = props;
    const { name, version } = config?.com ?? {};
    const { codeResponse } = getComHubEvent({ name, version });
    this.codeResponse = codeResponse;
  }

  componentDidMount(): void {
    this.addListener();
  }

  componentWillUnmount(): void {
    this.removeListener();
  }

  addListener = () => {
    eventHub.on(this.codeResponse, this.resolverCode);
  };

  removeListener = () => {
    eventHub.off(this.codeResponse, this.resolverCode);
  };

  resolverCode = (res: string) => {
    const { name, version } = this.props?.config?.com ?? {};
    const moduleName = `${name}/${version}`;
    const id = `${moduleName}/${uuid(10)}`;
    new Function(res)();
    window.dub(id, [moduleName], (module: any) => {
      const currentModule = module.default;
      this.setState({ Block: currentModule });
    });
  };

  render(): React.ReactNode {
    const { Block } = this.state;
    return Block && <Block {...this.props} />;
  }
};

export default WidthDebugWrapper;
