/* eslint-disable @typescript-eslint/no-namespace */
import type { MutableRefObject, ReactNode } from 'react';
// import type { BasicAssetParams, ComInfo, FileType } from '@canvas/utilities';
import type { ConfigTypeSupportedInControllerRuntime, SandboxRunnerType } from '@canvix/shared';

export type PanelOption = {
  panelId: string;
  data?: any[];
};

export type ScreenOption = {
  screenId: number;
  publishToken?: string;
  width: number;
  height: number;
  comId: string;
};
export type RenderPanelOption = PanelOption | ScreenOption;

/**
 * Definition for responsive utils
 */
export type UnifiedRenderOption = Record<string, any> & {
  id: string;
  data?: any[];
  userProps?: Record<string, any>;
  /**
   * The chain is a uri to locate any component. It is come from a path array
   * that joined by the symbol `.`. Every single digit is the index in the
   * current depth.
   * i.e. For the current component uri the path array from root to itelf
   * is `[0, 1, 0]`, which means the location in component tree structure is:
   * --
   *   --
   *   --
   *     -- (represents here)
   * And written in a string is `0.1.0`.
   *
   * Generate a the new chain for the child/children components is required.
   * Normally a component will only increase the depth of the component tree
   * by one, so that the `nextChain` value passed to render function shall be in
   * the form of `{{CURRENT_CHAIN}}.{{CHILD_INDEX}}`. The `CURRENT_CHAIN` can
   * be read from the component top-level props `props.chain`.
   */
  nextChain: string;
}; /*  & {
  [exntendableKeys: string]: string;
} */
export type UnifiedRenderUtil = (opt: UnifiedRenderOption) => ReactNode;
export type ResponsivePanelUtils = {
  render: UnifiedRenderUtil;
};

export type UnifiedRenderOptionInsideWrapper = Pick<
UnifiedRenderOption,
'id' | 'data' | 'userProps'
> & {
  key: number;
};
export type UnifiedRenderUtilInsideWrapper = (opt: UnifiedRenderOptionInsideWrapper) => ReactNode;
export type ResponsivePanelUtilsInsideWrapper = {
  render: (opt: UnifiedRenderOptionInsideWrapper) => ReactNode;
};

export type PrimitiveUtils<PanelUtils> = {
  general: PanelUtils
};

export type ComUtils = {
  /** @deprecated */
  getAssetsUrl: (fileName: string) => string;
  $emit: (name: string, payload: any) => void;
  runInSandbox: SandboxRunnerType;
};

export type FullUtils<PanelUtils> = ComUtils &
PrimitiveUtils<PanelUtils>['general'] & {
  getRuntimeConfig: (config: any) => any;
};

export type ControllerBasicProps<
  Config,
  PanelUtils,
  Children,
  OptionProps extends object = object,
> = {
  workspaceId: string | null;
  /** `projectId` will be used as the uid to make property requests */
  projectId: string;
  config: Config;
  userProps?: Record<string, any>;
  primitiveUtils: PrimitiveUtils<PanelUtils>;
  children?: Children;
} & OptionProps;

export type ComponentRuntimeStaticProps<PanelUtils> = {
  utils: ComUtils &
  PrimitiveUtils<PanelUtils>['general'] & {
    containerRef: MutableRefObject<HTMLDivElement | null>;
  };
  userProps?: Record<string, any>;
};

export type LoaderRuntimeBasicProps<PanelUtils> = {
  ref: MutableRefObject<any>;
  mounted: () => void;
} & ComponentRuntimeStaticProps<PanelUtils>;

type ControllerCombinedBasicProps<ControllerProps extends { config: any }, ComRuntimeProps> = {
  ready: boolean;
  outer: ControllerProps;
  runtime: BasicStates<ControllerProps['config']> & ComRuntimeProps;
  transitionRef: MutableRefObject<any>;
  children: ReactNode;
};

export type BasicStates<Config> = {
  config: Config;
  data?: any;
};

export namespace ResponsiveController {
  export type Config = ConfigTypeSupportedInControllerRuntime;
  export type OptionProps = {
    chain: string;
    panelId: string;
    forceUpdate: () => void;
    childrenAllowed: Record<string, true>;
  };

  export type Props<Children = ReactNode> = ControllerBasicProps<
  Config,
  ResponsivePanelUtils,
  Children,
  OptionProps
  >;
  export type States = BasicStates<Config>;
  export type LoaderProps = LoaderRuntimeBasicProps<ResponsivePanelUtilsInsideWrapper> &
  States & { chain: string; };
  export type ContainerProps<Children = ReactNode> = ControllerCombinedBasicProps<
  Props<Children>,
  LoaderRuntimeBasicProps<ResponsivePanelUtilsInsideWrapper> & States
  >;
  export type KeyToNameEntries = [keyof Config, string];
  export type ComponentRuntimeProps = Omit<LoaderProps, 'mounted' | 'ref' | 'form'>;
  export type LoaderBasicProps = Record<string, any> & Pick<LoaderProps, 'ref' | 'mounted'>;
}
