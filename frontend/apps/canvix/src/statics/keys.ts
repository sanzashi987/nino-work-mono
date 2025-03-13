import { DeltaKey, FileType, ComInfo, ActiveDescriber } from '@/types';
import { uuid } from '@/utils';

/**
 * used as the link to concat `panelId` and `comId`
 * to one string
 */
export const SPLIT_KEY = '_$$_';

export function parseCrossPanelId(input: string) {
  const [panelId, componentId] = input.split(SPLIT_KEY);
  return {
    panelId,
    componentId
  };
}
export function isPanel(type: string) {
  return type === 'panel';
}

export function composeCrossPanelId(panelId: string, componentId: string) {
  return `${panelId}${SPLIT_KEY}${componentId}`;
}

export function composeDraggableId(panelId: string, componentId: string, type: string) {
  return isPanel(type) ? componentId : composeCrossPanelId(panelId, componentId);
}

export function isDefaultPanel(panelId: string) {
  return panelId === 'default';
}

export function isLogical(type: string) {
  return type === 'logical';
}

export function isEdge(type: string) {
  return type === 'edge';
}

export const LAYER_ROUTE_KEY = '.';

export function parseLayerString(routeString: string) {
  if (routeString === '') return [];
  return routeString.split(LAYER_ROUTE_KEY).map(Number);
}

export function composeLayerString(path: number[]) {
  return path.join(LAYER_ROUTE_KEY);
}

export const COM_NAME_ID_SPLITER = '_';

export function parseNameId(nameId: string) {
  const splitArr = nameId.split(COM_NAME_ID_SPLITER);
  const random = splitArr.pop();
  const name = splitArr.join(COM_NAME_ID_SPLITER);
  return {
    name,
    id: random
  };
}

export function composeNameId(name: string, uid: string) {
  return `${name}${COM_NAME_ID_SPLITER}${uid}`;
}

export function generateNameId(name: string) {
  return composeNameId(name, uuid());
}

export function getNearestPanelId(describer: ActiveDescriber) {
  return isPanel(describer.type) ? describer.id : describer.panelId;
}

export const deltaKeyJoiner = '/';
const NULL_DELTAS = {
  globalBreakpoint: null,
  localBreakpoint: null,
  theme: null,
  comId: null
};
export const parseDeltaKey = (deltaKey: string) => {
  const deltas = deltaKey.split(deltaKeyJoiner);
  if (deltas.length !== 4) return { ...NULL_DELTAS };
  return {
    globalBreakpoint: deltas.at(0)!,
    localBreakpoint: deltas.at(1)!,
    theme: deltas.at(2)!,
    comId: deltas.at(3)!
  };
};

type DeltaKeyComposer = {
  globalBreakpoint: string;
  localBreakpoint: string;
  theme: string;
  comId: string;
};

export function composeDeltaKey({
  globalBreakpoint,
  localBreakpoint,
  theme,
  comId
}: DeltaKeyComposer): DeltaKey {
  return [globalBreakpoint, localBreakpoint, theme, comId].join(deltaKeyJoiner) as DeltaKey;
}

/** ---------- palette --------- */

export const packageKeyJoiner = '_$$_';

type PackageKeyComposer = {
  type: FileType;
  com?: ComInfo | null;
};

/**
 * 基于com、type生成package配置的key值
 * @param params
 * @returns
 */
export const composePackageKey = (params: PackageKeyComposer) => {
  const { type, com } = params;
  if (type === 'panel') return 'panel';
  if (!com) return null;
  const { name, version, user } = com;
  if (!user) return `${name}${packageKeyJoiner}${version}`;
  return `${name}${packageKeyJoiner}${version}${packageKeyJoiner}${user}`;
};

export function composeModuleKey(name: string, user?: string | null) {
  return (user ? `${name}(${user})` : name);
}

/** 基于面板id,组件id生成新的事件名称 */
export const composeDataEventName = (
  basename: string,
  config: { panelId: string; comId?: string }
) => {
  const { panelId, comId } = config;
  return `${basename}.${panelId}.${comId}`;
};
