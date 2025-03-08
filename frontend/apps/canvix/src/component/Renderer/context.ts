import type { BasicAssetParams, ComInfo } from '@canvas/utilities';

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
