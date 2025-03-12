import { ComponentClass, ComponentProps, ComponentType } from 'react';

import { LogicalNodeProps, BasicAssetParams, LayerList } from '../panel';
import { ComInfo, IDComConfig, Identifier } from './common';
import {
  BasicStates, ControllerBasicProps, FullUtils, LoaderRuntimeBasicProps, ResponsiveController, ResponsivePanelUtils, ResponsivePanelUtilsInsideWrapper
} from './controller';
import { ConfigTypeSupportedInControllerRuntime } from './responsive';
import { ConnectorProps } from '../event-core';

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

export interface Controller<
  Config extends IDComConfig['config'],
  PanelUtils,
  Children,
  OptionProps extends object = object,
  PanelUtilsRuntime extends object = object,
> extends ComponentClass<
  ConnectorProps<ControllerBasicProps<Config, PanelUtils, Children, OptionProps>>,
  BasicStates<Config>
  > {
  memoProps():LoaderRuntimeBasicProps<PanelUtilsRuntime> & BasicStates<Config>;
  createUtils(): FullUtils<PanelUtilsRuntime>;
  mounted(): void;
  getIdentifier(): Identifier
  getRuntimeProps(): LoaderRuntimeBasicProps<PanelUtilsRuntime> & BasicStates<Config>
}

export type ComWrapperInstance = Controller<
ConfigTypeSupportedInControllerRuntime,
ResponsivePanelUtils,
LayerList,
ResponsiveController.OptionProps,
ResponsivePanelUtilsInsideWrapper
>;

export type ComWrapperProps = ComponentProps<ComWrapperInstance>;

type ConnectInputProps = ComWrapperProps;

type KeyToRemove = keyof Omit<ResponsiveController.OptionProps, 'chain'>;
export type ConnectOuptutProps = { id: string } & Omit<
ConnectInputProps,
'projectId' | 'workspaceId' | 'config' | 'children' | KeyToRemove
>;

export type ConnectInput = ComponentType<ConnectInputProps>;
export type ConnectOuptut = ComponentType<ConnectOuptutProps>;

export type NodeComType = ComponentClass<
ConnectorProps<LogicalNodeProps>,
BasicStates<LogicalNodeProps['config']>
>;
