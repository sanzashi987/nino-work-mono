import { PanelMetaRuntime, RootMetaType } from './meta';

export type MergeParams = {
  delta: PanelMetaRuntime['delta'];
  core: Record<string, any>;
  id: string;
  theme: string;
  breakpoint: RootMetaType['breakpoint'];
  breakpoints: ProjectMetaConfig['breakpoints'];
};
