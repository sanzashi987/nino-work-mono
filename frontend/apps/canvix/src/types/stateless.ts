import { Responsive } from '@canvas/component-factory';
import type { LogicalType, EdgeType } from '@canvas/event-core';
import { CWidgets } from '@canvas/ui-components';

export type ActiveDescriber = {
  id: string;
  panelId: string;
  /** 父组件id,视图组件激活时需要该字段，用于复制时判断组件是否是同一层级 */
  parentId?: string;
  // type: string;
  type: Responsive.ComItemType | LogicalType | EdgeType;
};

export type CrossPanelLayerItem = {
  id: string;
  nextId: string;
  panelId: string;
  /** Local Path (root is the panel item belongs to) */
  path: number[];
};

// export type LocalDescriber = {
//   globalPath: number[];
//   localPath: number[];
// } & ActiveDescriber;

type CandidateMark = ActiveDescriber & {
  /** from own panel*/
  localPath: number[];
};

export type MarkedCandidates = CandidateMark[];

export type FlattenMarkedLayers = {
  candidates: CandidateMark[];
  /** The components shall apply migration (move filter, core, delta to target panel) */
  crossPanelLayers: CrossPanelLayerItem[];
};

export type Actives = Record<string, ActiveDescriber>;
export type StringObject = Record<string, string>;

export type ScaleType = number;

export enum EditorStatus {
  flex = 'flex',
  interaction = 'interaction',
  route = 'route',
  // sql = 'sql',
}

export const EditorName = {
  flex: '视图编辑器',
  interaction: '交互编辑器',
  route: '路由编辑器',
  sql: 'SQL生成器',
};

export type LayoutStatus = {
  layerPanel: boolean;
  componentPanel: boolean;
  configPanel: boolean;
  historyPanel: boolean;
  previewPanel: boolean;
  configMode: CWidgets.ConfigMode;
};

export type LayoutPanelName = Exclude<keyof LayoutStatus, 'configMode'>;

/** default to `default` */
export type ActiveTheme = 'light' | string;
export type ActiveBreakpoint = 'default' | string;
export type ActivePanel = 'default' | string;
