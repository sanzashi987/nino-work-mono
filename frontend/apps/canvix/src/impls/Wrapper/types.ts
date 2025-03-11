import type Controller from '@/component/Controller';
import type { ResponsiveController, ResponsivePanelUtils, ResponsivePanelUtilsInsideWrapper } from '@/component/Controller';
import type { ConfigTypeSupportedInControllerRuntime, LayerList, BasicAssetParams, ComInfo } from '@/types';

export type ComWrapperInstance = Controller<
ConfigTypeSupportedInControllerRuntime,
ResponsivePanelUtils,
LayerList,
ResponsiveController.OptionProps,
ResponsivePanelUtilsInsideWrapper
>;

export type ComWrapperProps = ComWrapperInstance['props'];

export type RuntimeInterface = {
  getAssetsUrl: (params: BasicAssetParams) => string;
  getRealBasic: (
    basic: Record<string, any>,
    screenId?: string,
  ) => {
    /** 适用于spring */
    springBasic: Record<string, any>;
    /** 直接作用于style */
    normalBasic: Record<string, any>;
  };
  loadModule: (
    params: Pick<ComInfo, 'name' | 'version'> & { id: string; user?: string },
  ) => Promise<any>;
  cachedComponents: Map<string, any>;
};
