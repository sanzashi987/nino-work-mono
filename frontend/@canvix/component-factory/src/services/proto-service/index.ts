import React, { Component } from 'react';
import { HandlerCollection, EventCollection, ServiceProps } from './types';

function defaultGetEndpoint(endpointMap: Map<string, any>, serviceName: string) {
  return [...endpointMap.entries()]
    .filter((e) => e[1].isPublic)
    .map(([name, desc]) => ({
      id: `${serviceName}.${name}`,
      ...desc,
    }));
}

abstract class ProtoService<Props extends ServiceProps = ServiceProps> extends Component<Props> {
  declare static serviceName: string;
  declare static configKey: string;

  public $emit;

  constructor(props: Props) {
    super(props);
    const { events = new Map(), serviceName } = this.constructor as any;
    const testers = [...events.keys()].map((el) => new RegExp(el));
    this.$emit = (e: string, payload?: any) => {
      props.$emit(testers.some((regex) => regex.test(e)) ? `${serviceName}.${e}` : e, payload);
    };
  }

  static getComponentActions = (e: HandlerCollection, s: string, p: Record<string, any>) =>
    defaultGetEndpoint(e, s);

  static getComponentEvents = (e: EventCollection, s: string, p: Record<string, any>) =>
    defaultGetEndpoint(e, s);

  shouldComponentUpdate(nextProps: ServiceProps) {
    return nextProps.config !== this.props.config;
  }

  componentDidMount(): void {
    this.props.selfRef.current = this;
  }

  componentWillUnmount(): void {
    this.props.selfRef.current = null;
  }

  render(): React.ReactNode {
    return null;
  }
}

export default ProtoService;
export { typeToService, updateTypeToService } from './annotations';
export type { ServiceComponent, GetIdentifierType } from './types';
