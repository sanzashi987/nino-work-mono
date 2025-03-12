import React from 'react';
import type { Identifier } from '../com-config';
import { EndpointType, EventCollection, HandlerCollection } from '../component';

type EndpointGetter<T> = (
  handle: T,
  serviceName: string,
  packageJSON: Record<string, any>,
) => EndpointType[];

export interface GetComponentEndpoints {
  getComponentEvents: EndpointGetter<EventCollection>;
  getComponentActions: EndpointGetter<HandlerCollection>;
}

export type GetIdentifierType = () => Identifier;

export type ServiceProps = {
  $emit(e: string, payload?: any): void;
  config: Record<string, any>;
  getIdentifier: GetIdentifierType;
  selfRef: React.RefObject<any>;
  instanceRef: React.RefObject<any>;
  transitionRef: React.RefObject<any>;
  setState: React.Component['setState'];
};

type GetHandleFunction<T> = (
  handle: T,
  serviceName: string,
  packageJSON: Record<string, any>,
) => EndpointType[];

type ServiceStaticValues = {
  getComponentActions: GetHandleFunction<HandlerCollection>;
  getComponentEvents: GetHandleFunction<EventCollection>;
  events: EventCollection;
  handlers: HandlerCollection;
  serviceName: string;
  configKey: string;
};

export type ServiceComponent = React.ComponentType<ServiceProps> & ServiceStaticValues;
