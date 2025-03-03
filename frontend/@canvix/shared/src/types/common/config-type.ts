export type IDConfig = {
  config: { id: string };
};

export type Default = {
  id: string;
  name: string;
};

export type DefaultAttr = Record<string, any>;

export type SubpanelInfoDefault = {
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
  dashboardId: string;
  comId: string;
  rakToken?: string;
  panelId: string;
} & Pick<ComInfo, 'name' | 'version' | 'user' | 'isDebugger'>;
