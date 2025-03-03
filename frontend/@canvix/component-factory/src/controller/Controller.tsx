import React, { createRef } from 'react';
import { createMemo, shallowEqual, Identifier } from '@canvas/utilities';
import { Connector, ConnectorProps } from '@canvix/event-core';
import type {
  IDConfig,
  FullUtils,
  ControllerBasicProps,
  BasicStates,
  LoaderRuntimeBasicProps,
  IDComConfig
} from './types';
import { typeToService, ServiceComponent } from '../services';

export function parseConfig(typeToService: Record<string, ServiceComponent>, config: any) {
  return Object.keys(typeToService)
    .sort()
    .filter((key) => !!config[key])
    .map((key) => typeToService[key]);
}

export abstract class ServiceConnector<P extends IDConfig, S extends IDConfig> extends Connector<
P,
S
> {
  protected _ready = false;

  protected transitionRef = createRef<any>();

  protected usedServices: ServiceComponent[] = [];

  private serviceRef: Record<string, any> = {};

  private scopedEmit;

  private scopedIdGetter;

  private scopedSetState;

  runtimeServices: ServiceComponent[] = [];

  constructor(props: ConnectorProps<P>) {
    super(props);
    this.state = { config: props.config, data: {} } as unknown as S;
    const enabledServices = this.getEnabledServices();
    this.usedServices = parseConfig(enabledServices, props.config);
    this.scopedEmit = this.$emit.bind(this);
    this.scopedIdGetter = this.getIdentifier.bind(this);
    this.scopedSetState = this.setState.bind(this);

    this.serviceRef = Object.fromEntries(
      this.usedServices.map((Service) => [Service.serviceName, React.createRef()])
    );
    this.services = new Proxy(this.serviceRef, { get: (target: any, propKey: string) => target[propKey].current });
    this.runtimeServices = this.usedServices;
  }

  getEnabledServices() {
    return typeToService;
  }

  UNSAFE_componentWillReceiveProps(nextProps: ConnectorProps<P>) {
    if (nextProps.config !== this.props.config) {
      this.setState({ config: nextProps.config });
    }
  }

  /**
   * *** Shall be explicitly used in the future render function ***
   */
  renderServices() {
    return this.runtimeServices.map((Service) => (
      <Service
        // @ts-ignores
        selfRef={this.serviceRef[Service.serviceName]}
        key={Service.configKey}
        $emit={this.scopedEmit}
        config={(this.props.config as any)[Service.configKey]}
        setState={this.scopedSetState}
        instanceRef={this.ref}
        transitionRef={this.transitionRef}
        getIdentifier={this.scopedIdGetter}
      />
    ));
  }

  abstract getIdentifier(): Identifier;
}

abstract class Controller<
  Config extends IDComConfig['config'],
  PanelUtils,
  UtilsOption,
  Children,
  OptionProps extends {} = {},
  PanelUtilsRuntime extends {} = {},
> extends ServiceConnector<
  ControllerBasicProps<Config, PanelUtils, UtilsOption, Children, OptionProps>,
  BasicStates<Config>
  > {
  containerRef = createRef<HTMLDivElement | null>();

  private memoProps = createMemo(
    (state: BasicStates<Config>, userProps?: Record<string, string>) => {
      const { getRuntimeConfig, ...other } = this.constantProps;
      //  config attr/basic => runtime attr/basic
      const { attr, basic } = getRuntimeConfig(state.config);
      return {
        ...state,
        config: {
          ...state.config,
          attr,
          basic
        },
        userProps,
        ...other
      } as LoaderRuntimeBasicProps<PanelUtilsRuntime, UtilsOption> & BasicStates<Config>;
    }
  );

  protected constantProps; /* : Omit<
    LoaderRuntimeBasicProps<PanelUtils, UtilsOption>,
    'userProps' | 'config'
  >; */

  getIdentifier() {
    const comInfo = this.props.config.com!;
    return {
      ...comInfo,
      dashboardId: this.props.dashboardId,
      comId: this.props.config.id,
      panelId: (this.props as any).panelId
    };
  }

  abstract createUtils(): FullUtils<PanelUtilsRuntime, UtilsOption>;
  // when the componentloader first render/mount the given view-components
  // but it can be triggered multiple times, as the unmount state will make the
  // component loader unmount as well.
  abstract mounted(): void;

  constructor(
    props: ConnectorProps<
    ControllerBasicProps<Config, PanelUtils, UtilsOption, Children, OptionProps>
    >
  ) {
    super(props);
    const { getRuntimeConfig, ...other } = this.createUtils();
    this.constantProps = {
      utils: {
        ...other,
        containerRef: this.containerRef
      },
      ref: this.ref,
      mounted: this.mounted.bind(this),
      getRuntimeConfig
    } as const;
  }

  shouldComponentUpdate(
    nextProps: ConnectorProps<
    ControllerBasicProps<Config, PanelUtils, UtilsOption, Children, OptionProps>
    >,
    nextState: BasicStates<Config>
  ) {
    return (
      this.props.children !== nextProps.children
      || this.state !== nextState
      || !shallowEqual(this.props.userProps, nextProps.userProps)
    );
  }

  toCleanup = { current: [] as (() => void)[] };

  cleanup() {
    this.toCleanup.current.forEach((cb) => {
      cb();
    });
  }

  emitInstance = () => {
    this.$emit('instance.init', {
      containerRef: this.containerRef,
      transitionRef: this.transitionRef,
      componentRef: this.ref,
      cleanup: this.toCleanup
    });
  };

  componentWillUnmount(): void {
    super.componentWillUnmount();
    this.cleanup();
  }

  getRuntimeProps(): LoaderRuntimeBasicProps<PanelUtilsRuntime, UtilsOption> & BasicStates<Config> {
    return this.memoProps(this.state, this.props.userProps);
  }
}

export default Controller;
export { typeToService };
