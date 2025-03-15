import { CFunctionRequiredInterface } from './cFunction';
import { ComInfo } from '../../com-config';
import { CanvixResponsePromise, Pagination, Request, Response } from '../../requests';
import { AssetsSearchParams, BasicAssetParams, TemplateAssetParams } from '../../panel/assets';
import { ChangeParams } from './basic';
import { ConfigMode } from '../../component/basic';
import { MenuContextValueType } from '../contextMenu';

export type ComponentContextType = {
  comInfo: Omit<ComInfo, 'icon' | 'category'>;
  comId: string;
  workspaceCode: string;
  utils: {
    getComponentTemplateAssetsUrl: (params: TemplateAssetParams) => string;
    uploadAssets: (file: File, params?: Request.AssetsUploadPayload) => CanvixResponsePromise<any>;
    getAssetsUrl: (p: BasicAssetParams) => string;
    getAssetsList: (p: AssetsSearchParams) => CanvixResponsePromise<any>;
    getAssetsGroupList: () => CanvixResponsePromise<any>;
    getBaseRoute: () => string;
  } & CFunctionRequiredInterface;
};

export type FontOption = {
  label: string;
  value: string; // 字体资产命名以font_开头
  updateTime?: string;
  projectCode: string;
};

export type LoadFontsType = (params: {
  family: string;
  updateTime?: string;
  user?: string;
}) => void;

export type FontAssetContextType = {
  fontList: FontOption[];
  updateFontList(): void;
  loadFonts: LoadFontsType;
  getFontList: (param: Request.FontListPayload) => CanvixResponsePromise<
  {
    data: Response.FontResponseType[];
  } & Pagination
  >;
};

export type PropertyContextType = {
  isValueUpdated: (params: Omit<ChangeParams, 'value' | 'end'>) => boolean;
  openContextMenu: MenuContextValueType;
  resetConfig: (params: Omit<ChangeParams, 'value' | 'end'>) => void;
  deleteConfig?: (params: { keyChains: ChangeParams['keyChain'][] }) => void;
};

export type StatusContextType = {
  breakpoint: 'default' | string;
  configMode: ConfigMode;
  setConfigMode: React.Dispatch<React.SetStateAction<ConfigMode>>;
};
