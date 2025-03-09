/* eslint-disable react/no-unused-class-component-methods */
import { Component, createRef, ReactNode } from 'react';
import { ReplaySubject, asyncScheduler } from 'rxjs';
import { troubleshoot } from '@/utils';
import type {
  IDConfig,
  PanelServiceInstance,
  PanelServiceCtor,
  PanelPropsType,
  ChannelMessageType,
  ConnectEntry,
  ConnectorProps,
  ServiceReturnType,
  ChannelObserver,
  PanelConnectType
} from '@/types';

type AsyncChannelType = ReturnType<typeof createChannel>;

function createChannel<T>() {
  return new ReplaySubject<T>(10, undefined, asyncScheduler);
}

export abstract class DuplexChannel<P extends PanelPropsType = PanelPropsType> extends Component<P> {
  protected emitters: Record<string, AsyncChannelType> = {};

  protected receiver = createChannel<ChannelMessageType>();

  protected responsives = new Set<string>();

  protected mutations: Record<string, ServiceReturnType> = {};

  readonly services: Record<string, PanelServiceInstance> = {};

  constructor(props: P) {
    super(props);
    this.receiver.subscribe({ next: this.onMessage });
  }

  componentDidUpdate(prevP: P, prevS: any) {
    this.responsives.forEach(($name) => {
      this.services[$name].updateConfig();
    });
  }

  componentWillUnmount(): void {
    Object.values(this.services).forEach((service) => {
      service.onDestory?.();
    });
  }

  // abstract getServicesCtor(): PanelServiceCtor<any>[];
  abstract getServicesCtor(): any[]; // PanelServiceCtor<any>[];

  private onMessage = (payload: ChannelMessageType) => {
    const { type, data } = payload;
    this.services[type].onSchedule(data);
  };

  readonly connect: ConnectorProps['connect'] = (id, entry) => {
    const connector = this.createConnector(id, entry);
    return [this.mutations, connector];
  };

  // send message to the Channel receiver, like emit message
  readonly post = (payload: ChannelMessageType) => {
    this.receiver.next(payload);
  };

  // send message to a specific component connected to the Channel
  readonly push = (id: string, payload: ChannelMessageType) => {
    if (!this.emitters[id]) {
      this.emitters[id] = createChannel();
    }
    const instance = this.emitters[id];
    Promise.resolve().then(() => {
      instance.next(payload);
    });
  };

  protected createConnector(id: string, entry: ConnectEntry) {
    return () => {
      if (!this.emitters[id]) {
        this.emitters[id] = createChannel();
      }
      const observer: ChannelObserver = {
        next: (payload: ChannelMessageType<any>) => {
          entry(payload);
        }
      };
      const subscription$ = this.emitters[id]!.subscribe(observer);
      return () => {
        subscription$.unsubscribe();
        if (this.emitters[id].observers.length > 0) return;
        this.emitters[id].complete();
        delete this.emitters[id];
      };
    };
  }

  render(): ReactNode {
    return null;
  }
}

export class ConnectorCore {
  troubleshooter: any;

  disconnect = () => {};

  start = () => {};

  /** the ref of `mutations` in the panel connected */
  connected: Record<string, ServiceReturnType> = {};

  connectedKeys: string[] = [];

  constructor(
    private id: string,
    connect: PanelConnectType | undefined,
    public getServices: () => Record<string, any>
  ) {
    if (!connect) return;

    const [connected, start] = connect(id, this.entry);
    this.connected = connected;
    this.start = () => {
      const disconnect = start();
      this.disconnect = () => {
        // clean up
        this.connected = {};
        this.getServices = () => ({});
        this.connectedKeys = [];
        disconnect();
      };
    };
    this.connectedKeys = Object.keys(connected);
  }

  private entry = (payload: ChannelMessageType) => {
    const { type, data } = payload;
    this.connected[type]?.handle.call(this, data);
  };

  @troubleshoot
  $emit(event: string, value: any) {
    // Object.keys(this.connected).forEach((key) => {
    this.connectedKeys.forEach((key) => {
      const service = this.connected[key];
      if (service.supportedEvents?.test(event)) {
        service.emit(this.id, event, value);
      }
    });
  }
}

export class Connector<Config extends IDConfig = IDConfig, S = object> extends Component<
ConnectorProps<Config>,
S
> {
  /** component services */
  services: Record<string, any> = {};

  /** the ref binded to the component/node instance */
  protected ref = createRef<any>();

  core;

  constructor(props: ConnectorProps<Config>) {
    super(props);
    this.core = new ConnectorCore(
      props.config.id,
      props.connect,
      this.getConnectorServices.bind(this)
    );
  }

  componentWillUnmount() {
    this.core.disconnect();
  }

  getConnectorServices() {
    return this.services;
  }

  // for component
  readonly emit = (event: string, value: any) => {
    this.$emit(`instance.${event}`, value);
    // this.core.$emit(`instance.${event}`, value);
  };

  // for controller
  $emit(event: string, value: any) {
    this.core.$emit(event, value);
  }
}

export class DuplexChannelCore {
  protected emitters: Record<string, AsyncChannelType> = {};

  protected receiver = createChannel<ChannelMessageType>();

  protected responsives = new Set<string>();

  protected mutations: Record<string, ServiceReturnType> = {};

  readonly services: Record<string, PanelServiceInstance> = {};

  constructor(scope: any, serviceCtors: PanelServiceCtor<any, any>[]) {
    this.receiver.subscribe({ next: this.onMessage });

    serviceCtors.forEach((ctor) => {
      const { $name, $responsive, $supportedEvents } = ctor;
      // eslint-disable-next-line new-cap
      const instance = new ctor(scope, this.post, this.push);
      this.services[$name] = instance;
      if ($responsive) {
        this.responsives.add($name);
      }
      this.mutations[$name] = {
        emit: instance.emit,
        handle: instance.handle,
        supportedEvents: $supportedEvents
      };
    });
  }

  private onMessage = (payload: ChannelMessageType) => {
    const { type, data } = payload;
    this.services[type].onSchedule(data);
  };

  protected createConnector(id: string, entry: ConnectEntry) {
    return () => {
      if (!this.emitters[id]) {
        this.emitters[id] = createChannel();
      }
      const observer: ChannelObserver = {
        next: (payload: ChannelMessageType<any>) => {
          entry(payload);
        }
      };
      const subscription$ = this.emitters[id]!.subscribe(observer);
      return () => {
        subscription$.unsubscribe();
        if (this.emitters[id].observers.length > 0) return;
        this.emitters[id].complete();
        delete this.emitters[id];
      };
    };
  }

  readonly connect: ConnectorProps['connect'] = (id, entry) => {
    const connector = this.createConnector(id, entry);
    return [this.mutations, connector];
  };

  // send message to the Channel receiver, like emit message
  readonly post = (payload: ChannelMessageType) => {
    this.receiver.next(payload);
  };

  // send message to a specific component connected to the Channel
  readonly push = (id: string, payload: ChannelMessageType) => {
    if (!this.emitters[id]) {
      this.emitters[id] = createChannel();
    }
    const instance = this.emitters[id];
    Promise.resolve().then(() => {
      instance.next(payload);
    });
  };

  updateConfig() {
    this.responsives.forEach(($name) => {
      this.services[$name].updateConfig();
    });
  }

  onDestory(): void {
    Object.values(this.services).forEach((service) => {
      service.onDestory?.();
    });
  }
}
