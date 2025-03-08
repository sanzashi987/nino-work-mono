import type { InteractionConfigType } from '@canvix/event-core';

export const getItemSwitchable = (
  activeNodes: string[],
  activeEdges: string[],
  nodes: InteractionConfigType['nodes'],
  edges: InteractionConfigType['edges']
): null | boolean => {
  const edgeStatus = Array.from(new Set(activeEdges.map((e) => !!edges[e]?.disable)));
  const nodeStatus = Array.from(new Set(activeNodes.map((e) => !!nodes[e]?.disable)));
  // 选中节点是否已全部导出到交互编辑器
  const nodeExport = activeNodes.every((id) => nodes[id]);
  if (!nodeExport) return null;
  let res = null;
  if (edgeStatus.length === 1 && nodeStatus.length === 0) {
    res = edgeStatus.at(0)!;
  } else if (edgeStatus.length === 0 && nodeStatus.length === 1) {
    res = nodeStatus.at(0)!;
  } else if (edgeStatus.length === 1 && edgeStatus.length === nodeStatus.length) {
    if (edgeStatus[0] === nodeStatus[0]) {
      res = edgeStatus.at(0)!;
    }
  }
  return res;
};

export const getCopyEnable = (activeNodes: string[], activeEdges: string[]): boolean => {
  if (!activeNodes.length || activeEdges.length) return false;
  const newList = activeNodes.filter((e) => /^@nodes\/.*_[0-9a-zA-Z]{6}/.test(e));
  return newList.length === activeNodes.length;
};

export type MenuColorPaletteType = Record<string, {
  name: string;
  foregroundColor: string;
  backgroundColor: string;
}>;

export type MenuLogicalNodesType = {
  name: string;
  cn_name: string;
  version: string;
  category: string;
  type: string;
}[];
