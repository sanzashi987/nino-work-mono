/* eslint-disable @typescript-eslint/ban-types */
import { ComInfo } from '../com-config';
import { FileType, DataConfigType } from '../services';
import { WithDescriber } from './core';

export type NodeFileType = FileType | LogicalType;

export type LogicalType = 'logical';
export type EdgeType = 'edge';

type AnyObject = Record<string, any>;
export type NodeBaseType<T extends AnyObject = AnyObject> = {
  id: string;
  com?: ComInfo;
  top: number;
  left: number;
  type: NodeFileType;
  disable?: boolean;
  attr?: T;
};

export type InteractionNodeType<
  T extends AnyObject = AnyObject,
  S extends DataConfigType = DataConfigType,
> = NodeBaseType<T> & {
  name?: string;
  data?: S;
};

export type InteractionNodeTypeRuntime<T extends AnyObject = AnyObject> = NodeBaseType<T> & {
  name: string;
};

export type LogicalNodeConfig<T extends AnyObject = AnyObject> = {
  name: string;
  attr: T;
  type: LogicalType;
} & Omit<NodeBaseType<T>, 'com'> &
Required<Pick<NodeBaseType<T>, 'com'>>;

/** Edges */
export declare type EdgeBasic = WithDescriber<{
  source: string;
  sourceNode: string;
  target: string;
  targetNode: string;
}>;

export declare type Edge<T extends Record<string, any> = {}> = {
  id: string;
  disable?: boolean;
  type?: string;
  markerStart?: string;
  markerEnd?: string;
} & EdgeBasic &
T;
// export interface BasicPanel extends DuplexChannel<PanelPropsType> {
//   getPanelId(): string;
// }

type ExportConfig = {
  nodeExport: boolean;
};

export type InteractionConfigType = {
  nodes: Record<string, InteractionNodeType>;
  edges: Record<string, Edge>;
};

export type ConnectableContextType = {
  components: Record<string, ExportConfig>;
  interaction: InteractionConfigType;
};

export type NativeConnectableContextType = {
  getProcessEnv: () => Record<string, any>;
};

export type NativePayloadStructure<Type, Value> = {
  type: Type;
  value?: Value;
};

export type NativeMessageReceived = NativePayloadStructure<
'message',
{
  event: string;
  params: Record<string, any>;
  node: string;
}
>;

export type NativeMessageToSend = NativePayloadStructure<
'message',
{
  instanceId: string;
  node: string;
  handler: string;
  params: Record<string, any>;
}
>;
