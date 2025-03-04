import type { ConfigType } from '../responsive/config';
import type { TargetPlatformSpecifier } from './platform';

type MappingTargetType = 'number' | 'string' | 'boolean';

export type MappingStructure = Record<
string,
{
  description: string;
  type: MappingTargetType;
  optional?: boolean;
}
>;

export type SingleSourcePackage = {
  description: string;
  fields: MappingStructure;
  name: string;
  /** 是否默认开启受控模式 */
  controlledMode?: boolean;
};

export type DataConfigPackage = {
  [sourceName: string]: SingleSourcePackage;
};

export type BasicCanvasFieldType = {
  cn_name: string;
  icon: string;
  category: string;
  type: ConfigType['type'];
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
    flatten?:Record<string, any>;
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

export type BasicPackageType = {
  name: string;
  version: string;
  user: string | null;
  isDebugger?: boolean;
};

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
      basic?:Record<string, any>
    };
  } & Record<string, Record<string, any>>;
};
export type CanvixFieldType = BasicCanvasFieldType & {
  children?: PackageChildrenType;
};
export type PackageJSONType = BasicPackageType & {
  canvas: CanvixFieldType;
};
export type ComponentPackageType = BasicPackageType & CanvixFieldType;
