/* eslint-disable @typescript-eslint/ban-types */
import { createContext } from 'react';
import type {
  PanelMetaRuntime,
  RootMetaType,
  EditorFeaturesType,
  EditorFeaturesRegisterType,
  ThemeContextType
} from '@app/types';
import type { SourceKey } from '@/types';
import { noop } from '@/utils';

function returnVoidObject() {
  return {};
}

export const RootMetaContext = createContext<RootMetaType | null>(null);

export const PanelMetaContext = createContext<PanelMetaRuntime | null>(null);

export const EditorFeatures = createContext<EditorFeaturesType>({});
export const EditorFeaturesRegister = createContext<EditorFeaturesRegisterType>({
  registerFeatures: () => [],
  unregisterFeatures: noop,
  getFeaturesAsync: returnVoidObject
});

export type DubLoaderType = {
  dub: (name: string, deps: string[] | Function, callback: (module: any) => void) => void;
  cachedComponents: Map<string, any>;
};

/**
 * static method & cache object provider
 */
export const DubLoaderContext = createContext<DubLoaderType>({
  dub: noop,
  cachedComponents: new Map()
});

/**
 * The `updateBreakpoint` method is provided in `Screen` component,
 * with the `RootMetaContext`.
 * The `Panel` is the component takes and only takes this method,
 * but it is a nth-great-grandchildren component of `Screen`,
 *  to avoid passing method through stacks of component
 */
export const UpdateBreakPointEditor = createContext<(width: number) => void>(noop);

export type ScreenType = 'screen' | 'model' | 'template';
export type ScreenConfigContextType = {
  screenCode: string;
  userIdentify: string;
  projectCode: string | null;
};
export const ScreenConfigContext = createContext<ScreenConfigContextType>({
  projectCode: '',
  userIdentify: '',
  screenCode: ''
});

/**
 * 编辑器context
 */
export type EditorConfigContextType = {
  basic: {
    /**
     * 页面类型，与后端接口保持一致
     * @description 用于更新数据源时，接口入参
     */
    type: ScreenType;
    guidePath: string;
    /**
     * @example "页面","模板","模块"
     */
    title: string;
  };
  /** 数据源面板相关配置 */
  dataPanel?: {
    sourceTypes: SourceKey[];
  };
  /** 组件面板相关配置 */
  componentPanel?: {
    custom?: boolean;
    debug?: boolean;
  };
};

export const EditorConfigContext = createContext<EditorConfigContextType | null>(null);

export type DebugInfo = {
  mode: boolean;
  url: string;
  secret: string;
};
export type DebugInfoContextType = {
  debugInfo: DebugInfo;
  setDebugger: (k: DebugInfo) => void;
};
export const DebugInfoContext = createContext<DebugInfoContextType | null>(null);

export const ThemeContext = createContext<ThemeContextType | null>(null);
