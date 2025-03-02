import { createContext } from 'react';
import type { SourceKey } from '@canvix/shared';
import { noop } from '@canvix/utils';
import type {
  PanelMetaRuntime,
  RootMetaType,
  EditorFeaturesType,
  EditorFeaturesRegisterType,
  DubLoaderType,
  ThemeContextType
} from '@app/types';

function returnVoidObject() {
  return {};
}

export const RootMetaContext = createContext<RootMetaType | null>(null);
RootMetaContext.displayName = 'RootMeta';

export const PanelMetaContext = createContext<PanelMetaRuntime | null>(null);
PanelMetaContext.displayName = 'PanelMeta';

/**
 * Used by redux store provider, all initValue is set to `undefined`
 * For the context `any`, refered to https://react-redux.js.org/api/provider
 * shows "Initial value doesn't matter, as it is overwritten with the internal
 * state of Provider." so that given the type `any`. The context only plays as
 * a required middleware, the type of the store will be annotated from the
 * `createStoreFromSlice` factory by its generics
 * */
export const RootStoreContext = createContext<any>(undefined);
RootStoreContext.displayName = 'RootStore';
export const ItemStateContext = createContext<any>(undefined);
ItemStateContext.displayName = 'ItemState';
export const UIStateContext = createContext<any>(undefined);
UIStateContext.displayName = 'UIState';
export const EditorStateContext = createContext<any>(undefined);
ItemStateContext.displayName = 'EditorState';

export const EditorFeatures = createContext<EditorFeaturesType>({});
EditorFeatures.displayName = 'EditorFeatures';
export const EditorFeaturesRegister = createContext<EditorFeaturesRegisterType>({
  registerFeatures: () => [],
  unregisterFeatures: noop,
  getFeaturesAsync: returnVoidObject
});
EditorFeaturesRegister.displayName = 'EditorFeaturesRegister';

/**
 * static method & cache object provider
 */
export const DubLoaderContext = createContext<DubLoaderType>({
  dub: noop,
  cachedComponents: new Map()
});
DubLoaderContext.displayName = 'DubLoader';

/**
 * The `updateBreakpoint` method is provided in `Screen` component,
 * with the `RootMetaContext`.
 * The `Panel` is the component takes and only takes this method,
 * but it is a nth-great-grandchildren component of `Screen`,
 *  to avoid passing method through stacks of component
 */
export const UpdateBreakPointEditor = createContext<(width: number) => void>(noop);
UpdateBreakPointEditor.displayName = 'UpdateBreakPointEditor';

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
ScreenConfigContext.displayName = 'ScreenConfig';

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
EditorConfigContext.displayName = 'EditorConfig';

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
DebugInfoContext.displayName = 'DebugInfoContext';

export const ThemeContext = createContext<ThemeContextType | null>(null);
ThemeContext.displayName = 'ThemeContext';
