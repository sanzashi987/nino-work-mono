import {
  ConnectorProps,
  DuplexChannelCore,
  InteractionService,
  PanelServiceCtor
} from '@canvix/event-core';
import { ServiceConnector } from './Controller';
import type { PrimitiveUtils, StaticUtils } from './types';
import { Default, Responsive } from '../services/types';

type PanelConfigType<T> = Default & {
  type: 'panel';
  basic: Record<string, any>;
  hide?: Responsive.HiddenMode;
} & T;

export type PanelConfigProps<T extends object = object> = {
  config: PanelConfigType<T>;
  data?: any;
};

export type PanelState = {
  panelData: any;
  config: PanelConfigProps['config'];
};

const services = [InteractionService];

export abstract class BasicPanel<
  P extends PanelConfigProps,
  S extends PanelState,
  PanelUtils extends object = object,
  UtilsType extends object = object,
  LogicalUtilsType extends object = object,
> extends ServiceConnector<P, S> {
  state: S;

  duplexChannel;

  primitiveUtils: PrimitiveUtils<PanelUtils, UtilsType>;

  logicalUtils: LogicalUtilsType;

  constructor(props: ConnectorProps<P>) {
    super(props);
    this.state = { config: props.config, panelData: props.data ?? null } as S;
    this.duplexChannel = new DuplexChannelCore(this, this.getServicesCtor());

    this.primitiveUtils = { general: { ...this.buildUtils(), ...this.getStaticUtils() } };
    this.logicalUtils = this.getLogicalUtils();
  }

  abstract getLogicalUtils(): LogicalUtilsType;
  abstract buildUtils(): PanelUtils;
  // abstract getPanelStyle(): CSSProperties;
  abstract getStaticUtils(): StaticUtils<UtilsType>;
  abstract render(): React.ReactNode;

  getServicesCtor(): PanelServiceCtor<any, any>[] {
    return services;
  }

  componentDidUpdate(prevProps: Readonly<ConnectorProps<P>>, prevState: S): void {
    this.duplexChannel.updateConfig();
  }

  componentWillUnmount(): void {
    super.componentWillUnmount();
    this.duplexChannel.onDestory();
  }
}
