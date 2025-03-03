import type {
  ComItemType, ConfigType, Default, DefaultAttr, HiddenMode
} from '@canvix/shared';

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
type PanelInfo = Default;

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
