import { ComWrapperProps, ConnectOuptut } from '../com-config';
import { UnifiedRenderUtil } from '../com-config/controller';
import { InteractionConfigType, LogicalNodeConfig, PanelPropsType } from '../event-core';
import { DynamicPanelMeta, LayerList, RootMetaType } from './meta';
import { Dimension, ThemeMetaType } from './responsive';

export type PanelProps = Omit<PanelPropsType, 'scale'>;

type EventHubProps = {
  eventHub: EventEmitter;
};

export type DynamicEditorPanelProps = PanelProps & EventHubProps;

export type PreviewPanelProps = PanelProps & {
  updateBreakpoint: (width: number) => void;
};

export type ReportPanelMetaType = {
  dimension?: Dimension;
  data?: any;
};

export type UpdateBreakpointProps = {
  updateBreakpoint: (width: number) => void;
};

export type ReportPanelMeta = {
  reportPanelMeta(panelId: string, meta: ReportPanelMetaType): void;
};

export type ReportPanelProps = PanelProps & ReportPanelMeta;
export type ReportPreviewPanelProps = PreviewPanelProps & ReportPanelMeta;

export type EditorPanelProps = PanelProps &
EventHubProps & {
  getLayerTreeItem(panelId: string): TreeItem | undefined;
};

/**
 * @deprecated panel does not have specs
 */
export type ControlledPanelProps = PanelProps & {
  data: any[];
};

export type TempValueHandles = {
  setValue: React.Dispatch<React.SetStateAction<Record<string, any>>>;
  getValue: () => Record<string, any>;
};

export type ThemeContextType = {
  /** 设置当前主题 */
  setCurrentTheme: (theme: string) => void;
  /** 主题配置 */
  themes: ThemeMetaType;
  /** 当前主题 */
  theme: string;
};

export type PanelLogicalUtilsType = {
  getProcessEnv: () => Record<string, any>;
  getVariable: (id: string, global?: boolean) => any;
  setVariable: (id: string, val: any, global?: boolean) => void;
  useVariable: (
    variableId: string,
    global?: boolean,
  ) => readonly [any, (variableId: string, value: any) => void];
  themeContext: React.Context<ThemeContextType | null>;
};

export type LogicalFullUtilsType = {
  $emit(e: string, val?: any): void;
} & PanelLogicalUtilsType;

export type LogicalNodeProps = {
  projectId: string;
  config: LogicalNodeConfig;
  logicalUtils: PanelLogicalUtilsType;
  panelId: string;
};

export type ComponentRendererProps = {
  layers: LayerList;
  // Com: ConnectOuptut;
} & Pick<ComWrapperProps, 'connect' | 'primitiveUtils'>;
export type NodeRendererProps = {
  projectId: string;
  panelId: string;
  nodes: InteractionConfigType['nodes'];
  connect: ComWrapperProps['connect'];
  panelLogicalUtils: PanelLogicalUtilsType;
};

export type NodeRendererType = React.FC<NodeRendererProps>;
export type ComponentRendererType = React.FC<ComponentRendererProps>;

export type RendererBundle = {
  ComponentRenderer: ComponentRendererType;
  NodeRenderer: NodeRendererType;
};

export type CreatePanelFactoryType = (
  this: any,
  ViewWrapper: ConnectOuptut,
  StaticWrapper: ConnectOuptut,
) => UnifiedRenderUtil;

// -> panel.next

export type PanelPropsFromContext = {
  panelMeta: DynamicPanelMeta;
  globalMeta: Pick<
  RootMetaType,
  'projectId' | 'getProcessEnv' | 'getGlobalVariable' | 'setGlobalVariable'
  >;
};

export type PanelOnlyProps = {
  panelId: string;
  data?: any;
  renderBy?: string;
};

export type PanelRuntimeProps = Pick<PanelConfigProps, 'config'> & PanelPropsFromContext;

export type PanelMinimalProps = PanelRuntimeProps & PanelOnlyProps;

export abstract class PanelMinimal extends BasicPanel<
PanelMinimalProps,
PanelState,
ResponsivePanelUtils,
PanelLogicalUtilsType
> {
  declare depotRef: React.RefObject<LocalVariableDepot>;

  declare render: () => React.ReactNode;
}

type RuntimePropsKey = keyof PanelRuntimeProps;

export type PanelInputProps = PanelMinimal['props'];

export type PanelOutputProps = Omit<PanelInputProps, RuntimePropsKey>;

export type MakePanelOutputProps<T extends PanelInputProps> = Omit<T, RuntimePropsKey>;
