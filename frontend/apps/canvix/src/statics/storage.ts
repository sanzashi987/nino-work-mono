import { ValueRW } from '@/component/value';
import {
  DynamicPanelMeta, PackagePaths, FileType, ComInfo, PackagePath,
  ComponentItemType
} from '@/types';
import { eventHub, PreviewPanelEvents } from './event-hub';
import { composePackageKey } from './keys';

const packagePathKey = 'packagePaths';
const blockListKey = 'blockList';

class StorageHub extends ValueRW {
  /**
    * 获取项目中实际用到的路径配置
    * @description 保存项目前调用；基于运行时的路径配置，项目中的组件清理脏数据
    * @param projectConfig 项目根配置
    * @returns
    */
  getUsedPaths = (projectConfig: { panels: Record<string, DynamicPanelMeta> }) => {
    const { panels } = projectConfig;
    // 必定包含主面板
    const packageKeys: Array<string | null> = [composePackageKey({ type: 'panel' })];

    Object.values(panels).forEach(({ components, interaction: { nodes } }) => {
    // 组件
      Object.values(components).forEach((item) => {
        const { type, com } = item;
        packageKeys.push(composePackageKey({ com, type }));
      });
      // 逻辑节点
      Object.values(nodes).forEach((item) => {
        const { type, com } = item;
        packageKeys.push(composePackageKey({ com, type }));
      });
    });

    // 清理脏数据
    const res: PackagePaths = {};
    const paths = this.getPackagePaths();
    new Set(packageKeys).forEach((key) => {
      if (!key || !paths[key]) return;
      res[key] = paths[key];
    });
    // initPackagePaths(res);
    return res;
  };

  /**
   * 获取页面中的组件路径配置
   * @returns
   */
  getPackagePaths() {
    return this.getValue()[packagePathKey] as PackagePaths;
  }

  /**
   * 更新单个组件路径配置
   * @description 组件升级、新增时调用
   * @param params
   */
  updatePackagePaths(params: {
    type: FileType;
    com?: ComInfo;
    config?: PackagePath;
  }) {
    const { type, com, config } = params;
    const key = composePackageKey({ com, type });
    if (!key || !config) return;
    const oldValue = this.getPackagePaths();
    const newConfig = {
      ...oldValue,
      [key]: config
    };
    this.initPackagePaths(newConfig);
  }

  mergePackagePaths(paths: PackagePaths) {
    if (Object.keys(paths).length < 1) return;
    const oldValue = this.getPackagePaths();
    const newConfig = {
      ...oldValue,
      ...paths
    };
    this.initPackagePaths(newConfig);
  }

  initPackagePaths(config: PackagePaths) {
    eventHub.emit(PreviewPanelEvents.updatePackagePaths, config);
    this.setValue({ [packagePathKey]: config });
  }

  /** 获取缓存的模块列表，用于全局搜索 */
  getCachedBlockList() {
    return this.getValue()[blockListKey] as ComponentItemType[] | null;
  }

  /** 设置缓存的模块列表，用于全局搜索 */
  setCachedBlockList(list: ComponentItemType[]) {
    this.setValue({ [blockListKey]: list });
  }
}

export default new StorageHub();
