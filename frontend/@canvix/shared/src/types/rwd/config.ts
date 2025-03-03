import type { ComInfo, Default, DefaultAttr, SubpanelInfoDefault } from '../common/config-type';
import type { DataConfigType, DataConfigTypeRuntime } from '../data-source/service';

export const enum HiddenMode {
  visible,
  implicit,
  unmount,
}

export function isUnmounted(state: any) {
  return state === HiddenMode.unmount;
}

type Basic<T extends object = any> = {
  basic: T;
};

type ComBasic<Data extends DataConfigType = DataConfigType> = {
  data?: Data
  com: ComInfo
  hide: HiddenMode
};

export type Com<Data extends DataConfigType = DataConfigType> = {
  type: 'com';
  lock?: boolean;
} & ComBasic<Data> &
Default;

export type ComRuntime<Attr extends DefaultAttr = DefaultAttr> = Com<DataConfigTypeRuntime> &
Basic & {
  attr: Attr;
  children?: { id: string; type: string }[];
};

export type Subcom<
    Attr extends DefaultAttr = DefaultAttr,
    Data extends DataConfigType = DataConfigType,
  > = {
    type: 'subcom';
  } & ComBasic<Data> &
  Default & {
    attr: Attr;
  };

export type SubcomRuntime<Attr extends DefaultAttr = DefaultAttr> = Subcom<
Attr,
DataConfigTypeRuntime
> & {
  children?: { id: string; type: string }[];
};

export type Container<Data extends DataConfigType = DataConfigType> = {
  type: 'container';
  lock?: boolean;
} & ComBasic<Data> &
Default;

export type ContainerRuntime<Attr extends DefaultAttr = DefaultAttr> =
    Container<DataConfigTypeRuntime> &
    Basic<Record<string, any>> & {
      attr: Attr;
      children?: { id: string; type: string }[];
    };

export type Panel = {
  type: 'panel';
  com?: undefined;
  hide?: HiddenMode;
} & Omit<Default, 'name'>;

export type PanelRuntime = Panel & Basic & Default;

export type ConfigType = Com | Subcom | Container | Panel;

export type ConfigTypeExcludePanel = Exclude<ConfigType, Panel>;

export type ConfigTypeRuntime = ComRuntime | SubcomRuntime | ContainerRuntime | PanelRuntime;

export type SubpanelInfo = Omit<SubpanelInfoDefault, 'viewControlled'>;

export type ComItemType = ConfigType['type'];
