import { ConfigMode, PackageConfigType } from './basic';
import { ShowInPanel } from './logic';

export type ValueType = number | string;

/**
 * CObjectTabs 运行时、配置时value类型
 */
export type CObjectValueType = Record<string, { value: any; order: number }>;
/**
 * CTabs 运行时value类型
 * 配置时value类型为Array<Record<string,any>>
 */
export type CArrayValueType = Array<{ id: string; value: any }>;

export type CreateObject = (target: Record<string, any>, key: string | number, value: any) => Record<string, any>;

export type PathsHandler = (params: {
  target: any;
  result?: Record<string, any>;
  lastKey?: string;
}) => PathsConfig;

export type GenerateType<T extends PackageConfigType = PackageConfigType> = {
  /**
   * 获取控件值
   * @description 无该方法时，默认取default字段
   * @param params
   * @returns
   */
  getValue?: (params: {
    propertyFn: PropertyFn;
    createObject: CreateObject;
    result: Record<string, any>;
    key: string;
    config: T;
  }) => any;

  /**
   * 获取控件的路径信息（资产、调色板、数据转换等）
   * @description 无该方法时，不收集该控件的路径信息
   * @description parseValue,mergeValue需指定该方法
   * @description 支持一个路径对应多个不同类型的配置（调色板、数据转换等）
   * @description type: palette,parse,assets,merge
   * @param params
   * @example
   *  createObject(result, key, {
      configs: [{
        type: 'palette',
        value,
      }, {
        type: "parse",
      }],
      component: config.type,
    });
   * @returns
   */
  getPaths?: (params: {
    pathsHandler: PathsHandler;
    config: T;
    result: Record<string, any>;
    createObject: CreateObject;
    key: string;
  }) => any;

  /**
   * 配置项值转为运行时值
   * @description attr config ==> attr runtime
   * @description 无该方法时，运行时值=配置项值
   */
  parseValue?: (params: {
    value: any;
    config: {
      getAssetsUrl?: (fileName: string) => string;
    };
  }) => any;

  /**
   * 属性合并方法，组件升级时调用
   * @description 无该方法时,采用lodash默认合并方法
   * @param params
   * @returns
   */
  mergeValue?: (params:{
    /** 模板属性，基于新版本生成的默认值属性 */
    targetValue:any;
    /** 旧版本属性，用户修改过后的属性值，优先级更高 */
    sourceValue:any;
    mergeObject: (target:any, source:any) => any;
  }) => any;
};

export type PropertyFn = (target: any, obj?: Record<string, any> | any[]) => Record<string, any> | any[];

export type PathsConfig = Record<string, {
  component: string;
  configs: Array<{
    type: string;
    value: string
  }>
}>;

export type ChangeParams = {
  keyChain: Array<number | string>;
  value: any;
  end?: boolean;
  themeEnable?: boolean;
  breakpointsEnable?: boolean;
};

/**
 * C系列基础组件Props,运行时type
 *
 * 如：CText、CColor
 */
export type CProps<Config extends PackageConfigType = any> = {
  className?: string;
  config: Config;
  keyChain: Array<number | string>;
  /**
   * name路径，用于搜索时展示，不需要透传给CComponents
   */
  nameKeyChain: string[];
  /**
   * 缩进层级、CMenu默认为0，其余控件默认为1
   */
  level?: number;
  value: Config['default'];
  /**
   * 控件值是否更新过，用于控制高亮样式及重置操作
   */
  isModified?: boolean;
  onChange: (params: ChangeParams) => void;
};

export type VisibleParams = {
  /** 显隐配置 */
  showInPanel: ShowInPanel | undefined;
  /** 路径 */
  keyChain: (string | number)[];
  /** 配置项在哪些模式下隐藏 */
  hideInModes?: ConfigMode[];
  /** 当前选择的模式 */
  currentMode?: ConfigMode;
};

/**
 * C系列高阶组件Props，运行时type
 *
 * GuiComponentHoc以及CGroup、CSuite、CTab等组合控件
 */
export interface CHocProps<Config extends PackageConfigType = PackageConfigType> extends CProps<Config> {
  /**
   * 是否位于CSuite组件中
   */
  inSuite?: boolean;

  // value: any;
  utils?: {
    getVisibleByConfig: (params: VisibleParams) => boolean;
  };
}
