import type { InteractionConfigType } from '../event-core';
import type {
  ComDefault, ComItemType, ConfigType, DefaultAttr, HiddenMode
} from '../com-config';
import type { LocalVariableCollection } from './variable';
import type { ActiveBreakpoint, ActiveTheme } from '../stateless';
import type { EnvVariables } from '../env';
import { BreakpointMetaType } from './responsive';

/** Filters */
export type FilterType = {
  id: string;
  name: string;
  content: string;
  lock?: boolean;
  createTime?: number;
  updateTime?: number;
};

export type FiltersType = Record<string, FilterType>;
export type FilterDepInfo = Record<string, Array<{ id: string; name: string }>>;

/** Layers */
export type LayerItem = {
  id: string;
  children?: LayerItem[];
  type: ComItemType; // | RWDType;
};

export type LayerList = LayerItem[];

/** Panel Infos */
type ThemeKey = string; // 'dark' | '*';
type ComId = string;
type GlobalBreakpoint = string | '*';
type LocalBreakpoint = GlobalBreakpoint;
export type DeltaKey = `${GlobalBreakpoint}/${LocalBreakpoint}/${ThemeKey}/${ComId}`;
type PanelInfo = ComDefault;

export type PanelBasic<T extends PanelInfo = PanelInfo> = {
  components: Record<string, ConfigType>;
  info: T;
  /**
   * default attr and basic config of components
   */
  core: {
    [comId: string]: {
      attr?: DefaultAttr;
      basic?: Record<string, any>; // RWDCombination;
      hide?: HiddenMode;
    };
  };
  /**
   * Responsive to the themes and the breakpoints,
   * delta attr and basic config of components
   */
  delta: Record<DeltaKey, Record<string, any>>;
  layers: LayerList;
  interaction: InteractionConfigType;
  filters: FiltersType;
  variables: LocalVariableCollection;
};

export type SysMethods = {
  /** system Value getter */
  getProcessEnv: <T>() => EnvVariables<T>;
};

export type DynamicPanelMeta = PanelBasic<PanelInfo>;

/** for project editor & block editor */
export type PanelMetaType = {
  default: DynamicPanelMeta;
} & Record<string, DynamicPanelMeta>;

export type PanelMetaRuntime = DynamicPanelMeta;

export type RootMetaType = {
  panels: PanelMetaType;
  theme: ActiveTheme;
  /** global screen breakpoint */
  breakpoint: ActiveBreakpoint;
  /** won't change after assigned */
  projectId: string;
  breakpoints: BreakpointMetaType;
  /** globalValue getter */
  getGlobalVariable: (id: string) => any;
  setGlobalVariable: (id: string, val: any) => void;
} & SysMethods;
