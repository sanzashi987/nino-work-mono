import type { FileType } from '../data-source';

export type IDConfig = {
  config: { id: string };
};
export type IDConfigType = {
  config: {
    id: string
    type:FileType
  } ;
};

export type IDComConfig = {
  config: {
    id: string
    type: FileType;
    com?: ComInfo;
  };
};

export type ComDefault = {
  id: string;
  name: string;
};

export type DefaultAttr = Record<string, any>;

export type ComInfoBasic = {
  /**
   * 组件名称
   * @description 英文名，唯一标识
   * @example `chart-pie`
   */
  name: string;
  /**
   * 组件版本号
   * @example `1.1.0`
   */
  version: string;
  /**
   * 用户，不为null时为自定义组件
   * @example `admin`
   */
  user: string | null;
};

export type ComInfo = {
  /**
   * 组件分类
   * @description 用于图层中图标的显示
   * @example `chart-pie`
   */
  category: string;
  /**
   * 组件图标
   * @example `assets/chart-pie.png`
   */
  icon: string;
  // controller?: string;
  /**
   * 组件中文名，用于全局搜索
   * @example `饼图`
   */
  cn_name: string;
  /**
   * 是否开启调试模式
   */
  isDebugger?: boolean;
} & ComInfoBasic;

export type Identifier = {
  projectId: string;
  comId: string;
  panelId: string;
} & Pick<ComInfo, 'name' | 'version' | 'user' | 'isDebugger'>;
