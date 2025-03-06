import type { InteractionNodeType } from '@canvas/event-core';
import type { EndpointType, Legacy, Responsive } from '@canvas/component-factory';
import { MenuColorPaletteType } from './components/helpers';

export type EndpointResType = {
  source: EndpointType[];
  target: EndpointType[];
  childList: string[];
};

export type EndpointsStatusType = {
  deprecated: boolean;
  isDelete?: boolean;
  isVertical?: boolean;
};
type ToEndpointsReturnType = {
  endpoints: EndpointResType;
} & EndpointsStatusType;

type ToEndpointsType = (node: InteractionNodeType) => Promise<ToEndpointsReturnType>;
type GetRefNodeEndpointsType = (id: string, refresh?: boolean) => Promise<ToEndpointsReturnType>;
type FindComponentByIdType = (id: string) => Legacy.ConfigTypeRuntime | Responsive.ConfigType;

export type TailEditorInterface = {
  toEndpoints: ToEndpointsType;
  findComponentById: FindComponentByIdType;
  switchPanel(panelId: string): void;
  getRefNodeEndpoints: GetRefNodeEndpointsType;
  menuPalette: MenuColorPaletteType;
};
