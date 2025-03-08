import type { ReactNode } from 'react';

export type FeatureCategory = 'file' | 'edit' | 'view' | 'help' | 'hidden';

export type FeatureItemBase<T> = T & {
  id: string;
  category: FeatureCategory;
  group?: string;
  name: string | (() => string);
  icon?: ReactNode | (() => ReactNode);
  /**
   * @return {boolean} boolean type controls the button is clickable but always visible
   * @return {null} null will make the button disabled and invisible
   */
  disabled?: (() => boolean | null) | boolean | null;
  // currently do not support sub-features
  // children?: Record<string, FeatureItemBase<T>>;
  callback?: (...e: any[]) => any;
  scoped?: string;
  contextMenu?: boolean;
  /**
   * if `static` is `true`, the shortcut of feature will be binded to the `window`
   * in event capture phase and the `scoped` field will be ignored
   */
  static?: boolean;
};

export type FeatureItemProps = FeatureItemBase<{
  shortcutWin?: string | string[];
  shortcutMac?: string | string[];
}>;

export type FeatureItemRuntime = FeatureItemBase<{
  shortcutWin?: string | string[];
  shortcutMac?: string | string[];
  shortcutNode: ReactNode;
}>;

export type FeatureRuntimeMap = Record<string, FeatureItemRuntime>;

export type EditorFeaturesType = FeatureRuntimeMap;

export type EditorFeaturesRegisterType = {
  registerFeatures: (items: FeatureItemProps[]) => string[];
  unregisterFeatures: (ids: string[]) => void;
  getFeaturesAsync: () => FeatureRuntimeMap;
};
