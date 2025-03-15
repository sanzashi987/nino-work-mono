import { createContext } from 'react';
import { noop } from '@/utils';
import { ComponentContextType, FontAssetContextType, PropertyContextType, StatusContextType } from '@/types';

async function defaultAsyncResponse() {
  return { data: '', resultCode: 0, resultMessage: '' };
}

const defaultComInfo = { name: '', version: '', user: null, cn_name: '' };
const defaultComponentInfo: ComponentContextType = {
  comInfo: defaultComInfo,
  comId: '',
  workspaceCode: '',
  utils: {
    getComponentTemplateAssetsUrl: () => '',
    uploadAssets: defaultAsyncResponse,
    getAssetsUrl: () => '',
    getAssetsGroupList: defaultAsyncResponse,
    getAssetsList: defaultAsyncResponse,
    getDataList: defaultAsyncResponse as () => any,
    getAssetsDetailById: defaultAsyncResponse as () => any,
    getVariableList: () => [],
    getDataDetail: defaultAsyncResponse as () => any,
    getBaseRoute: () => ''
  }
};

const context = createContext<ComponentContextType>(defaultComponentInfo);
export const ComponentContext = context;

export const FontAssetContext = createContext<FontAssetContextType>({
  fontList: [],
  updateFontList: noop,
  loadFonts: noop,
  getFontList: defaultAsyncResponse as any
});
FontAssetContext.displayName = 'FontAsset';

export const PropertyContext = createContext<PropertyContextType>({
  isValueUpdated: () => false,
  openContextMenu: noop,
  resetConfig: noop
});
PropertyContext.displayName = 'PropertyContext';

export const StatusContext = createContext<StatusContextType>({
  breakpoint: 'default',
  configMode: 'simple',
  setConfigMode: noop
});
StatusContext.displayName = 'StatusContext';
