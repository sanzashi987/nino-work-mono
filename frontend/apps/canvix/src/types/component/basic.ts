import type { ShowInPanel } from './logic';

export type ConfigMode = 'simple' | 'detail';

/**
 * 基础的Config type，用户package.json中配置内容
 *
 * 所有控件的config type均继承于此类型
 */
export type PackageConfigType<Value = any> = {
  /**
   * 显示名称
   */
  name: string;
  /**
   * 控件类型
   */
  type: string;
  /**
   * 控件所占栅格数，取值范围1-24
   */
  col?: number;
  /**
   * tips显示内容
   */
  description?: string;
  /**
   * 动态切换显示、隐藏配置
   */
  showInPanel?: ShowInPanel;
  /**
   * 是否可以显示显示隐藏开关，默认为false，仅CGroup、CSuite等组合控件生效
   */
  enableHide?: boolean;
  /**
   * 是否显示标题，默认为true
   */
  nameVisible?: boolean;
  /**
   * 是否响应主题配置，默认不响应
   */
  themeEnable?: boolean;
  /**
   * 是否响应断点配置，默认不响应
   */
  breakpointsEnable?: boolean;
  /**
   * 配置项在哪些模式下隐藏，默认所有模式下均显示，不支持在详细模式下隐藏
   */
  hideInModes?: Exclude<ConfigMode, 'detail'>[];
  /**
   * 默认值，仅基础控件需要
   */
  default: Value;
  /**
   * 是否必填,必填时标签上会显示*
   * @description  处理逻辑由控件自行实现,目前仅以下组件实现该逻辑（CText）
   * @default false
   */
  required?: boolean;
};
