/* eslint-disable @typescript-eslint/no-namespace */
import type { ComInfo } from '@canvas/utilities';
import type { Absolute, Box, Flex } from './basic-service/types';
import type { FormConfigType } from './form-service/types';
import type {
  DataConfigType,
  DataConfigTypeRuntime,
  DataConfigPackage,
} from './data-service/types';
import { EndpointType } from './types';

export * from './basic-service/types';
export * from './form-service/types';
export type { EndpointType, /* Ctor, */ FieldsType } from './proto-service/types';
// export type { TransitionInputType } from './instance-service';
export * from './data-service/types';
export type Default = {
  id: string;
  name: string;
};

type DefaultAttr = Record<string, any>;

type SubpanelInfoDefault = {
  controlType?: 'form' | 'data';
  /**
   * @description has higher priority than `height` & `width`
   * when the value is set to `true`,
   * the `width` and `height` property will be set to `100%`
   */
  viewControlled: boolean;
  backgroundColor: string;
  backgroundImage: string | null;
  panelId: string;
  parent: string;
  width: number;
  height: number;
};
/**For Legacy Canvas Editor */
export namespace Legacy {
  type ComBasic<
    Attr extends DefaultAttr = DefaultAttr,
    Data extends DataConfigType = DataConfigType,
  > = {
    attr: Attr;
    data?: Data;
    com: ComInfo;
    supportTheme?: boolean; // 是否支持主题配置
    themeConfig?: Record<string, any>; // 主题配置对象
    assetsConfig?: Record<string, any>; // 资产配置对象
    hide: boolean;
    nodeExport: boolean;
  } & Default;

  export type Com<
    Attr extends DefaultAttr = DefaultAttr,
    Data extends DataConfigType = DataConfigType,
  > = {
    basic: Absolute;
    form?: FormConfigType | null;
    formEnable?: boolean; // 是否支持表单服务
    lock: boolean;
    type: 'com';
  } & ComBasic<Attr, Data>;

  export type ComRuntime<Attr extends DefaultAttr = DefaultAttr> = Com<Attr, DataConfigTypeRuntime>;

  export type Subcom<
    Attr extends DefaultAttr = DefaultAttr,
    Data extends DataConfigType = DataConfigType,
  > = {
    type: 'subcom';
  } & ComBasic<Attr, Data>;

  export type SubcomRuntime<Attr extends DefaultAttr = DefaultAttr> = Subcom<
    Attr,
    DataConfigTypeRuntime
  >;

  export type Group = {
    type: 'group';
    basic: Absolute;
    lock: boolean;
    hide: boolean;
    nodeExport: boolean;
  } & Default;

  export type Subpanel = {
    type: 'subpanel';
    nodeExport: boolean;
  } & Default;

  export type Refpanel = {
    type: 'refPanel';
    basic: Absolute;
  } & Omit<ComBasic<import('../../../../packages/components/ref-panel/type').Attr>, 'data'>;

  export type ConfigType = Com | Subcom | Group | Subpanel;

  export type ConfigMap = {
    com: Com;
    group: Group;
    subcom: Subcom;
    subpanel: Subpanel;
    refPanel: Refpanel;
  };

  export type ConfigTypeRuntime = ComRuntime | SubcomRuntime | Group | Subpanel | Refpanel;
  export type ConfigTypeSupportedInControllerRuntime =
    | ComRuntime
    | SubcomRuntime
    | Group
    | Refpanel;
  export type ConfigRuntimeMap = {
    com: ComRuntime;
    group: Group;
    subcom: SubcomRuntime;
    subpanel: Subpanel;
    refPanel: Refpanel;
  };

  export type SubpanelInfo = SubpanelInfoDefault;

  export type PackageChildrenType = {
    defaultCount?: number;
    viewControlled?: boolean;
    dataControlled?: boolean;
    controlType?: string;
    default: string[];
    supportTypes?: string[];
    supportNames?: string[];
    minCount?: number;
    maxCount?: number;
    addable?: boolean;
    template?: {
      name?: string;
      enableSort?: boolean;
      enableDuplicate?: boolean;
      enableRename?: boolean;
      width?: number;
      height?: number;
    };
  };
  export type CanvasFieldType = BasicCanvasFieldType & {
    supportDynamicPanel?: boolean;
    children?: PackageChildrenType;
    layerOrder?: string[];
  };
  export type PackageJSONType = BasicPackageType & {
    canvas: CanvasFieldType;
  };
  export type ComponentPackageType = BasicPackageType & CanvasFieldType;
}

/** 
 * For Responsive Web Design Editor
 * @deprecated replace by `Responsive` in `@canvas/types`
 * */
export namespace Responsive {
  export enum HiddenMode {
    visible = 0,
    implicit = 1,
    nonexistent = 2,
  }

  export function isUnmountMode(state: any) {
    return state === HiddenMode.nonexistent;
  }

  // type Basic<T extends Box = Box> = {
  type Basic<T extends {} = Record<string, any>> = {
    basic: T;
  };

  type ComBasic<Data extends DataConfigType = DataConfigType> = {
    data?: Data;
    com: ComInfo;
    hide?: HiddenMode;
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
    Default & {};

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
  // export type ConfigType = Com | Subcom | Subpanel;
  export type ConfigTypeExcludePanel = Exclude<ConfigType, Panel>;

  export type ConfigTypeRuntime = ComRuntime | SubcomRuntime | ContainerRuntime | PanelRuntime;

  export type SubpanelInfo = Omit<SubpanelInfoDefault, 'viewControlled'>;

  export type ComItemType = ConfigType['type'];

  // export type ConfigTypeSupportedInControllerRuntime = ComRuntime | SubcomRuntime | GroupRuntime;
  export type ConfigTypeSupportedInControllerRuntime =
    | ComRuntime
    | SubcomRuntime
    | ContainerRuntime; //| GroupRuntime;
  export type ConfigTypeSupportedPackageRuntime = ComRuntime | SubcomRuntime | ContainerRuntime;
  export type ViewConfigTypeSupportedInControllerRuntime = ComRuntime | ContainerRuntime;

  export type ConfigTypeSupportedInViewControllerRuntime = ComRuntime | ContainerRuntime;
  export type ConfigTypeSupportedInController = Com | Subcom | Container;

  export type PackageChildrenAcceptKey = 'type' | 'category' | 'name';
  export type AcceptItem = Partial<Record<PackageChildrenAcceptKey, string[] | string>> & {
    /** 当前类型最大允许数量 */
    maxCount?: number;
  };
  /**
   * package中子组件的accept配置
   * @description 数组间“或”运算、对象间“且”运算
   * [A,B] => A ||  B
   * {C,D} => C && D
   * @example  type为subcom并且category为数组中任意一项
   * [{
   *  type:"subcom",
   *  category: ["map2d-basemap","map2d-layer","map2d-topinfo"]
   * }]
   */
  export type PackageChildrenAccept = AcceptItem[];

  /**
   * package中子组件的配置
   */
  export type PackageChildrenType = {
    /** 用于定义合法的子组件类型 */
    accept: PackageChildrenAccept;
    /**
     * 默认子组件列表
     * @description 数组的每一项为一个对象，当custom字段为false时，可以缺省为仅包含name的字符串
     * @description "chart-config-y" 等价于 {name:"chart-config-y",custom:false}
     * @description 当为面板时，name值固定为panel
     * @example ["chart-config-tooltip","chart-config-legend"]
     * @example [{name:"chart-config-x",custom:true},"chart-confi g-tooltip"]
     */
    default?: Array<
      | string
      | {
          /** 组件name字段 */
          name: string;
          /** 是否是自定义组件,缺省时默认为false */
          custom?: boolean;
        }
    >;
    order?: Partial<Record<PackageChildrenAcceptKey, string>>[];
    /**
     * 是否显示children分类标签页，默认为true
     */
    tabVisible?: boolean;
    /** 子组件最大数量（所有子组件） */
    maxCount?: number;
    /**
     * 组件模板，以组件name为key,默认配置项为value
     */
    template?: {
      panel?: {
        info: {
          name: string;
        };
        basic?:Record<string,any>
      };
    } & Record<string, Record<string, any>>;
  };
  export type CanvasFieldType = BasicCanvasFieldType & {
    children?: PackageChildrenType;
  };
  export type PackageJSONType = BasicPackageType & {
    canvas: CanvasFieldType;
  };
  export type ComponentPackageType = BasicPackageType & CanvasFieldType;
}

export type ConfigTypeSupportedInControllerRuntime =
  | Legacy.ConfigTypeSupportedInControllerRuntime
  | Responsive.ConfigTypeSupportedInControllerRuntime;

export type EndpointsType = Record<
  string,
  Omit<EndpointType, 'id'> & {
    description?: string;
  }
>;

/**
 * different from the `targetPlatform` in `process.env`, as the env variable
 * indicates the webpage is about to running in which platform
 *
 * The `TargetPlatform` here defines what kind of features the native platform
 * side is suppoorted and can be configured over the webpage editor
 *
 * Although the app is built by `flutter`, the flutter side may not always implement
 * a feature on both `IOS` and `Android` platform.
 */
export type TargetPlatform = 'IOS' | 'H5' | 'Android' | 'miniProgram' | 'SPA';

export type TargetPlatformSpecifier = {
  [Key in TargetPlatform]?: boolean;
};

/**
 * ****************************************************
 * basic packagejson file type defined
 * ****************************************************
 */
export type BasicCanvasFieldType = {
  cn_name: string;
  icon: string;
  category: string;
  type: Legacy.ConfigType['type'] | Responsive.ConfigType['type'];
  view?: {
    width: number;
    height: number;
  };
  unique?: boolean; // controls the logical nodes can only be created once
  supportTheme?: boolean;
  platform?: TargetPlatformSpecifier;
  /**
   * 基础属性配置
   */
  basic?: {
    /**
     * 默认基础配置
     * @description 格式为一维的普通对象，key为css属性名，value为css属性值
     * @description 优先级： default > extend > flatten
     * @example
     * {
        "height": "150px",
        "width": "100%",
        "backgroundColor": "#f40"
      }
     */
    default?: Record<string, any>;
    /**
     * 扩展的基础配置
     * @description 格式同config，支持对新增和修改平台默认属性配置
     * @description 优先级： default > extend > flatten
     * @example
     * {
        "test": {
          "type": "CText",
          "name": "basic测试",
          "default": "xxx"
        }
      }
     */
    extend?: Record<string, any>;
    /**
     * 基于扁平化对象快捷配置
     * @description 格式为扁平化对象，key为对象扁平化路径值，例如："basic.layout.position",value为config配置对象，可部分缺省
     * @example 
     *  {
     *    "basic.size.minSize.height": {
     *      "default":"180px",
     *       "required": true
     *    }
     *  }
     * @description 优先级： default > extend > flatten
     */
    flatten?:Record<string,any>;
  };
  config?: Record<string, any>;
  apis?: DataConfigPackage;
  api_data?: {
    [key: string]: Array<Record<string, any>>;
  };
  form?: {
    showConfig: boolean;
  };
  events?: EndpointsType;
  handlers?: EndpointsType;
  verticalLayout?: boolean; // controlls the logical nodes' display layout style in interaction editor
};

type BasicPackageType = {
  name: string;
  version: string;
  user: string | null;
};
