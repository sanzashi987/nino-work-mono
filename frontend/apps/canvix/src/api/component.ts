import { ApiResponseType, FileType } from '@canvix/shared';

type UpdateComponentPayload = {
  componentCode: string;
  data: string;
  panelId: string;
  screenId: string;
  source: string;
  sourceName: string;
  sourceType: string;
};

/** 更新组件静态数据源 */
export const updateStaticData = (params: { code: string; json: string }) => authRequest.POSTFORM('files/upload', { ...params, suffix: 'json' });

/** 更新组件历史数据源配置 */
export const updateComponentRemote = (params: UpdateComponentPayload) => authRequest.POST('component-operation/update', params);

/** 请求组件历史数据源配置 */
export const fetchDataSource = (data: {
  screenId: string;
  id: string;
  sourceName: string;
  type: string;
  panelId: string;
}) => {
  const { screenId, id, sourceName, type, panelId } = data;
  return authRequest.GET('component-operation/find-source-new', {
    panelId,
    name: sourceName,
    type,
    screenId,
    code: id
  });
};

/** 获取平台已上架组件列表 */
export const getSystemComponents = (params: {}): ApiResponseType<
Array<{
  category: string;
  cn_name: string;
  createTime: string;
  icon: string;
  name: string;
  type: Exclude<FileType, 'group' | 'subpanel'>;
  updateTime: string;
  version: string;
  id: string;
}>
> => authRequest.POST('/component-operation/listAvailable', params);

/** 组件升级 */
export const upgradeComponent = (
  params: {
    id: string;
    version: string;
    name: string;
    config: any;
    panelId: string;
    user?: string | null;
  }[]
) => authRequest.POST('component-operation/offUpgrade', params);

/** 上架测试组件 */
export const onShelve = (file: Blob) => authRequest.POSTFORM('component-cus/uploadByStream', { file });
