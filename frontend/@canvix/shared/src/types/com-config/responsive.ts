import { DataConfigType, DataConfigTypeRuntime } from '../services';
import { ComInfo, ComDefault, DefaultAttr } from './common';

export const enum HiddenMode {
  visible,
  implicit,
  unmount,
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
ComDefault;

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
  ComDefault & {
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
ComDefault;

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
} & Omit<ComDefault, 'name'>;

export type PanelRuntime = Panel & Basic & ComDefault;

export type ConfigType = Com | Subcom | Container | Panel;

export type ConfigTypeExcludePanel = Exclude<ConfigType, Panel>;

export type ConfigTypeRuntime = ComRuntime | SubcomRuntime | ContainerRuntime | PanelRuntime;

export type ComItemType = ConfigType['type'];

export type ConfigTypeSupportedInControllerRuntime =
    | ComRuntime
    | SubcomRuntime
    | ContainerRuntime;

export type ConfigTypeSupportedPackageRuntime = ComRuntime | SubcomRuntime | ContainerRuntime;

export type ViewConfigTypeSupportedInControllerRuntime = ComRuntime | ContainerRuntime;

export type ConfigTypeSupportedInViewControllerRuntime = ComRuntime | ContainerRuntime;

export type ConfigTypeSupportedInController = Com | Subcom | Container;
