import type { InteractionConfigType } from '@canvas/event-core';

export const getItemSwitchable = (
  activeNodes: string[],
  activeEdges: string[],
  nodes: InteractionConfigType['nodes'],
  edges: InteractionConfigType['edges'],
): null | boolean => {
  const edgeStatus = Array.from(new Set(activeEdges.map((e) => !!edges[e]?.disable)));
  const nodeStatus = Array.from(new Set(activeNodes.map((e) => !!nodes[e]?.disable)));
  // 选中节点是否已全部导出到交互编辑器
  const nodeExport = activeNodes.every((id) => nodes[id]);
  if (!nodeExport) return null;
  let res = null;
  if (edgeStatus.length === 1 && nodeStatus.length === 0) {
    res = edgeStatus[0];
  } else if (edgeStatus.length === 0 && nodeStatus.length === 1) {
    res = nodeStatus[0];
  } else if (edgeStatus.length === 1 && edgeStatus.length === nodeStatus.length) {
    if (edgeStatus[0] === nodeStatus[0]) {
      res = edgeStatus[0];
    }
  }
  return res;
};

export const getCopyEnable = (activeNodes: string[], activeEdges: string[]): boolean => {
  if (!activeNodes.length || activeEdges.length) return false;
  const newList = activeNodes.filter((e) => /^@nodes\/.*_[0-9a-zA-Z]{6}/.test(e));
  return newList.length === activeNodes.length;
};

export type MenuColorPaletteType = Record<
  string,
  {
    name: string;
    foregroundColor: string;
    backgroundColor: string;
  }
>;

export type MenuLogicalNodesType = {
  name: string;
  cn_name: string;
  version: string;
  category: string;
  type: string;
}[];

export const menuConfig: MenuColorPaletteType = {
  block: {
    name: '模块接口',
    foregroundColor: '#6d4eff',
    // backgroundColor: '#4823a4',
    backgroundColor: '#3b217b',
  },

  'global-node': {
    name: '面板功能',
    foregroundColor: '#2d2e2f', //'var(--canvas-widget-darker-bgcolor)',
    backgroundColor: 'var(--canvas-ui-lvl1-bgcolor)',
  },
  // 'local-node': {
  //   name: '面板节点',
  //   foregroundColor: '#2d2e2f', //'var(--canvas-widget-darker-bgcolor)',
  //   backgroundColor: 'var(--canvas-panel-main-bgcolor)',
  // },
  'process-control': {
    name: '流程控制',
    foregroundColor: '#437C8E', //'#0d5dff',
    backgroundColor: '#1C292C', //'#0941b3',
  },
  'data-process': {
    name: '数据处理',
    foregroundColor: '#6C9135', //'#139b00',
    backgroundColor: '#252D1B', //'#0d6e00',
  },
  // 'input-device': {
  //   name: '输入设备',
  //   foregroundColor: '#4C4D83', //'#6d4eff',
  //   backgroundColor: '#1E1F2A', //'#4823a4',
  // },
  'result-process': {
    name: '结果处理',
    foregroundColor: '#2e589a',
    backgroundColor: '#2c3d58',
  },
};
