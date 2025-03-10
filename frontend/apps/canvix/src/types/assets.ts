export type BasicAssetParams = {
  isDebugger?: boolean;
  /**
   * @description 资产唯一表标识名 .e.g 资产id或者文件名
   */
  path: string;
  /**
   * @description 组件唯一标识名称
   */
  name?: string;
  /**
   * @description 组件当前版本
   */
  version?: string;
  /**
   * @description 自定义组件标识 1-自定义组件
   */
  cType?: 0 | 1;
  /**
   * @description 大屏id
   */
  screenId?: string;
  /**
   * @description 用于解决管理端资产替换后的缓存问题
   */
  updateTime?: string;
  /**
   * 用户名,saas版为projectCode
   */
  user?: string | null;
  /**
   * 页面类型
   * @deprecated 优先通过user字段获取
   */
  source?: 'model' | 'screen' | 'template';
};
/**
 * @description 上传资产方法参数
 * @deprecated replaced by `Request.AssetsUploadPyload` in `@canvas/types`
 */
export type UploadAssetParams = {
  /**
   * @description 资产分组，未分组传"-1"
   */
  groupCode?: string;
  /**
   * @description 资产类型，design:媒体(设计)资产，data:数据资产,font:字体资产,cover:缩略图
   */
  type: 'design' | 'font' | 'data' | 'cover';
};
/**
 * @description 根据文件名获取资产路径方法参数
 */
export type TemplateAssetParams = {
  /**
   * @description 组件唯一标识名称
   */
  name: string;
  /**
   * @description 组件当前版本
   */
  version: string;
  /**
   * @description 文件名
   */
  fileName: string;
  /**
   * @description 自定义组件标识
   */
  user?: string | null;
  /**
   * @description 本地开发测试组件
   */
  isDebugger?: boolean;
};
/**
 * @description 搜索资产方法参数
 */
export type AssetsSearchParams = {
  /**
   * @description fileType 资产分类中的二级类型 .e.g 全部，视频，图片....image/video/
   */
  fileType?: string;
  /**
   * @description filter 资产格式 .e.g 对应upload accept字段 eg:.png,image/png
   */
  filter?: string;
  /**
   * @description groupCode 资产分组
   */
  groupCode?: string | null;
  /**
   * @description fileName 要搜索的资产名称
   */
  fileName?: string;
  /**
   * @description sort 排序方式: 0-按时间正排序，1-按时间负排序
   */
  sort?: number;
  /**
   * @description page 页数
   */
  page?: number;
  /**
   * @description size 每页条目数量
   */
  size?: number;
  /**
   * @description type 资产分类
   */
  type?: 'design' | 'data' | 'font';
};
