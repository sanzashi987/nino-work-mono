import { ReactNode } from 'react';
import type { PackageComInfo } from '../com-config';

export type ComponentItemType = PackageComInfo & {
  groupCode?: string | null;
  projectCode?: string;
};

export type ModuleCategory = {
  /**
   * 分类名称
   */
  label: string;
  /**
   * 分类图标
   */
  icon?: ReactNode;
  /**
   * 子分类
   */
  children?: ModuleCategoryType;
  /**
   * 通过正则的方式，对组件category字段进行组件的过滤
   */
  categoryRegExp?: RegExp;
  /**
   * 过滤方法
   */
  filter?: (list: ComponentItemType[]) => ComponentItemType[];
  /**
   * 默认为true，为false时只能通过点击生成组件
   */
  draggable?: boolean;
  /**
   * 是否支持原生拖拽，默认为false
   * */
  native?: boolean;
  /**
   * 额外的配置，
   */
  config?: {
    /**
     * 逻辑节点前景色
     */
    foregroundColor?: string;
    /**
     * 逻辑节点背景色
     */
    backgroundColor?: string;
  };
};

export type ModuleCategoryType = Record<string, ModuleCategory>;
