import mitt from 'mitt';
/**
 * The global event chanel singleton that ***only***
 * for the editor
 */
export const eventHub = mitt();

export const UndoRedoEvents = {
  redo: 'undoRedo.redo',
  undo: 'undoRedo.undo',
  timeTravel: 'undoRedo.timeTravel',
  getHistoryStatus: 'undoRedo.getHistoryStatus',
  broadcast: 'undoRedo.broadcast'
} as const;

export const SidebarEvents = {
  globalFilter: 'sidebar.globalFilter',
  appConfigs: 'sidebar.appConfigs',
  layerPanel: 'sidebar.layerPanel',
  componentPanel: 'sidebar.componentPanel',
  historyPanel: 'sidebar.historyPanel',
  variablPanel: 'sidebar.variablPanel',
  upgradePanel: 'sidebar.upgradePanel',
  palettePanel: 'sidebar.palettePanel',
  debuggerPannel: 'sidebar.debuggerPannel'
};

export const ReservedPanelEvents = {
  configPanel: 'reserved.configPanel',
  previewPanel: 'reserved.previewPanel'
};

const ModalEvents = { globalSearch: 'modal.globalSearch' };

export const EditorEvents = {
  // refetchLayerState: 'editor.refetchLayerState',
  palettePreview: 'editor.palettePreview',
  rename: 'editor.rename',
  renameLogical: 'editor.renameLogical',
  ...UndoRedoEvents,
  ...ModalEvents,
  refreshBlock: 'editor.refreshBlock'
};

export const PreviewPanelEvents = { updatePackagePaths: 'previewPanel.updatePackagePaths' };

export const DataEvents = {
  /**
   * 主面板（编辑器面板、数据源配置面板）获取数据
   * @description 该事件由主面板发送至预览面板，n-->1
   * @description 使用时无需拼上面板id与组件id
   */
  requestDataInHostFrame: 'dataEvents.requestDataInHostFrame',
  /**
   * 预览面板中请求数据(预览面板iframe --> 预览面板内面板)
   * @description 该事件由预览面板中的iframe 发送至预览面板内的动态面板 ，1---> n
   * @description 使用时，需要拼接上面板id
   */
  fetchDataInSubFrame: 'dataEvents.fetchDataInSubFrame',
  /**
   * 主面板中设置数据
   * @description 需要精确到组件
   * @description 使用时需要拼接上面板id与组件id
   */
  setDataInHostFrame: 'dataEvents.setDataInHostFrame',
  /**
   * 预览面板中发送数据
   * @description 由iframe中的面板发送数据到iframe，n-->1
   * @description 使用时无需拼上面板id与组件id
   */
  sendDataInSubFrame: 'dataEvents.sendDataInSubFrame',
  /**
   * 在预览面板中获取数据
   * @description 需要精确到组件
   * @description 使用时需要拼接上面板id与组件id
   */
  getDataInSubFrame: 'dataEvents.getDataInSubFrame'
};
