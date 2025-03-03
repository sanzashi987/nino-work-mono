import type { IDConfig } from '@canvix/shared';
import type { ReactNode } from 'react';

type PanelConnectInputProps = {
  panelId: string;
  /**
   * The responsive panel does not need scale value
   */
  scale?: number;
};

export type ServiceReturnType = {
  emit: (id: string, eventName: string, value: any) => void;
  handle: (payload: any) => void;
  supportedEvents?: RegExp;
};

export type ChannelMessageType<T = any> = {
  type: string;
  data: T;
};

type ChannelSubscriber = (payload: any) => void;

export type ChannelObserver = {
  next: ChannelSubscriber;
  complete?: () => void;
  error?: (error: any) => void;
};

export interface ServiceHostInstance<P> {
  props: P;
}

export interface PanelServiceCtor<T, P> {
  new (
    scope: ServiceHostInstance<P>,
    post: PostMethodType<T>,
    push: PushMethodType<T>,
  ): PanelServiceInstance<T>;
  $responsive: boolean;
  $name: string;
  $supportedEvents: RegExp;
}

export type PostMethodType<T> = (payload: ChannelMessageType<T>) => void;
export type PushMethodType<T> = (id: string, payload: ChannelMessageType<T>) => void;
export interface PanelServiceInstance<T = any> extends ServiceReturnType {
  onSchedule(payload: T): void;
  updateConfig(): void;
  onDestory?(): void;
}

export type ConnectEntry<T = any> = (payload: ChannelMessageType<T>) => void;

export type PanelPropsType = {
  connect?: PanelConnectType;
  data?: any[];
} & PanelConnectInputProps;

export type PanelConnectType<T = any> = (
  id: string,
  entry: ConnectEntry<T>,
) => [Record<string, ServiceReturnType>, ConnectFactoryType];

type ConnectFactoryType = () => () => void;

export type SourceDesciption = {
  source: string; // the description for the endpoint of source node
  sourceNode: string; // id for source node
};

export type WithDescriber<T extends object> = {
  sourceDescriber?: Record<string, any>;
  targetDescriber?: Record<string, any>;
} & T;

export type TargetDescription = WithDescriber<{
  target: string;
  targetNode: string;
}>;

// `output` means the payload is recevied by the consumer `nodes & components`
export type InteractionInputPayloadType = {
  value: Record<string, any>;
} & TargetDescription;

// `output` means the payload is emitted by the consumer `nodes & components`
export type InteractionOutputPayloadType = {
  value: Record<string, any>;
} & SourceDesciption;

export type InteractionPayloadType = InteractionInputPayloadType | InteractionOutputPayloadType;

export type ConnectorProps<Configs extends IDConfig = IDConfig> = {
  connect?: PanelConnectType;
  children?: ReactNode;
} & Configs;
