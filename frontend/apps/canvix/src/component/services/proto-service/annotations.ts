import type {
  HandlerCollectionVal,
  EventCollection,
  ServiceComponent,
  AnnotationEndpointType
} from './types';

export const typeToService: Record<string, ServiceComponent> = {};

/**
 * explicitly annotate the service, shall be used with class
 * Usage:
 * ```
 * import { service } from 'proto-service/annotations'
 *
 * @service('DataService', configKey)
 * class DataService extends ProtoService{}
 * ```
 * @param serviceName to avoid class name loss due to the the bundling, and will be the part of endpoint uri,
 * @param configKey the key choose the value from the original passed config
 * @param supportTypes array of regex for supported component type to let controller whether this
 * service should be applied to the different type of component, `*` means applied to all
 */
export function service(serviceName: string, configKey: string, supportTypes = ['*']) {
  return function (ctor: any) {
    Object.defineProperties(ctor, {
      configKey: { value: configKey, writable: false },
      serviceName: { value: serviceName, writable: false }
    });

    // supportTypes.forEach((type) => {
    //   if (!typeToService[type]) {
    //     typeToService[type] = [];
    //   }
    //   typeToService[type].push(ctor);
    // });

    typeToService[configKey] = ctor;
  };
}

/**
 * explicitly annotate the handler, and write to the service static variable `handlers`
 * Usage:
 * ```
 * import { action, service } from 'proto-service/annotations'
 *
 * @service('DataService', configKey)
 * class DataService extends ProtoService{
 *   @action('ActionToBeTirggered', false, dataRawType)
 *   public action(){}
 * }
 * ```
 * @param name The name (mostly in locale language) for the handler
 * @param fields Defines the handler argument in type
 * @param isPublic Flag that defines the handler will be exported as global handler
 */
export function action(name: string, endpoint: AnnotationEndpointType, isPublic = true): any {
  return function (target: any, methodName: string, descriptor: PropertyDescriptor) {
    if (!target.constructor.handlers) target.constructor.handlers = new Map<string, HandlerCollectionVal>();
    target.constructor.handlers.set(`${methodName}`, {
      name,
      action: descriptor.value,
      isPublic,
      ...endpoint
    });
  };
}

/**
 * Factory that generate shortcut to export the event to the static varible `events` in the service class
 * Decorator will automatically read the property name as the value so that the type for the property shall
 * be `string`.
 *
 * Usage:
 * ```
 * import {event,service} from 'protoService'
 *
 * @service('DataService',configKey)
 * class DataService extends ProtoService{
 *   @event('The data finish fetched', false)
 *   public DataUpdated: string
 * }
 *
 * //....
 * this.$emit(this.DataUpdated, payload) // with property name
 * // or (not recommended)
 * this.$emit('DataUpdated', payload) // directly emit the event
 * ```
 * @param name The name (mostly in locale language) in detail for the event
 * @param fields Defines the action argument in type
 * @param isPublic Flag that defines the event will be exported as global action
 */
export function event(name: string, endpoint: AnnotationEndpointType, isPublic = true) {
  return function (target: any, propertyName: string) {
    if (!target.constructor.events) target.constructor.events = new Map() as EventCollection;
    target.constructor.events.set(`${propertyName.toString()}`, { name, isPublic, ...endpoint });
    // cannot be rewrittten afterwards
    Object.defineProperty(target, propertyName, {
      value: propertyName,
      configurable: false,
      writable: false,
      enumerable: false
    });
  };
}

/** 用于画布编辑器以及预览面板中动态改数据服务 */
export function updateTypeToService(configKey: string, serv: ServiceComponent) {
  typeToService[configKey] = serv;
}
