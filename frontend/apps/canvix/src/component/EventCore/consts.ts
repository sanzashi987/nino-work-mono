import { TargetDescription } from '@/types';

export const BLOCK_EVENT_PREFIX = 'block-event';
export const BLOCK_ACTION_PREFIX = 'block-action';

export const BLOCK_EVENT_ID = `@nodes/${BLOCK_EVENT_PREFIX}`;
export const BLOCK_ACTION_ID = `@nodes/${BLOCK_ACTION_PREFIX}`;
export const PANEL_NODE_ID = '@nodes/panel-node';

export const BLOCK_ACTION_INVOKE_FORWARD = 'instance.invokeForward';

export const PANEL_LOCAL_ID = 'canvix-panel';
export const PANEL_ACTION_SET_PANEL_DATA = 'setPanelData';
export const PANEL_NODE_ENTRY = 'entry';

const PANEL_EVENT_REFLECTER = 'reflect.event';
const PANEL_ACTION_REFLECTER = 'reflect.action';

function reflectEvent(nodeId: string) {
  return [nodeId, [{ target: PANEL_EVENT_REFLECTER, targetNode: PANEL_LOCAL_ID }]] as [
    string,
    TargetDescription[],
  ];
}
function reflectAction(nodeId: string) {
  return [nodeId, [{ target: PANEL_ACTION_REFLECTER, targetNode: PANEL_LOCAL_ID }]] as [
    string,
    TargetDescription[],
  ];
}

export const ENABLE_REFLECT_PANEL = [
  reflectAction(`${PANEL_NODE_ID}.instance.reflectPanelAction`),
  ...[
    `${PANEL_NODE_ID}.instance.reflectPanelEvent`,
    `${BLOCK_EVENT_ID}.instance.reflectPanelEvent`
  ].map(reflectEvent),
  [
    `${PANEL_LOCAL_ID}.${BLOCK_ACTION_INVOKE_FORWARD}`,
    [{ targetNode: BLOCK_ACTION_ID, target: BLOCK_ACTION_INVOKE_FORWARD }]
  ]
] as [string, TargetDescription[]][];

export function createReflectPanelHandlerPayload<T>(handlerName: string, handlerData: T) {
  return {
    target: PANEL_NODE_ENTRY,
    targetNode: PANEL_LOCAL_ID,
    value: {
      handlerName,
      handlerData
    }
  };
}
