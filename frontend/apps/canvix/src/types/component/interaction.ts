import type { TargetPlatformSpecifier } from './platform';

type FunctionalDesciber = {
  useChannel?: TargetPlatformSpecifier;
  pairType?: string;
};

type DesciberType = FunctionalDesciber & Record<string, any>;

export type FieldsType = {
  type: 'string' | 'boolean' | 'number' | 'array' | 'object' | 'any' | 'null' | 'pair';
  name: string;
  description?: string;
  children?: Record<string, FieldsType>;
  default?: any;
  optional?: boolean;
  pairType?: string;
};

export type AnnotationEndpointType = Omit<EndpointType, 'id' | 'name'>;

export type EndpointType = {
  id: string;
  name: string;
  fields?: FieldsType;
  /**
   * `describer` will be directly injected to the endpiont, and can
   *  passed to the `Handle` component in `tail-js`
   *  */

  describer?: DesciberType;
  description?: string;
};

export type EndpointsType = Record<
string,
Omit<EndpointType, 'id'> & {
  description?: string;
}
>;

type BaseCollection = {
  name: string;
  isPublic: boolean;
} & AnnotationEndpointType;

export type EventCollection = Map<string, BaseCollection>;
export type HandlerCollectionVal = BaseCollection & {
  action(...arg: any[]): void;
};
export type HandlerCollection = Map<string, HandlerCollectionVal>;
