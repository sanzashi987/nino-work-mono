import { Request, Responsive, Response } from '@canvas/types';
import { ApiResponseType } from '@canvas/utilities';
import { cancelTimeout } from '@app/consts';

/** 新增模块资产 */
export const addBlock = (data: Request.TemplateCreatePayload) => authRequest.POST('assetModel/add', data);

/** 删除模块资产 */
export const deleteBlock = (ids: string[]) => authRequest.DELETE('assetModel/delete', null, { data: ids });

/** 获取模块资产详情 */
export const getBlockInfo = (code: string) => authRequest.GET('assetModel/detail', { code });

/** 更新模块资产 */
export const updateBlock = (data: Responsive.API.TemplateUpdatePayload) => authRequest.POST('assetModel/update', data);

/** 分页查询模块资产列表 */
export const getBlockPageList = (data: Request.TemplateListPayload) => authRequest.POST('assetModel/pageList', data);

/** 查询模块资产列表全量  */
export const getBlockList = (): ApiResponseType<Response.BlockResponseType[]> => authRequest.POST('assetModel/list', {});

/** 新增分组 */
export const addBlockGroup = (groupName: string) => authRequest.POSTFORM('assetModel/addGroup', { groupName });

/** 获取分组列表 */
export const getBlockGroupList = (): ApiResponseType<Response.GroupResponseType> => authRequest.POST('assetModel/selectGroup', {});

/** 分组重命名 */
export const renameBlockGroup = (data: Request.TemplateUpdateGroupPayload) => authRequest.POSTFORM('assetModel/updateGroupsName', data);

/** 删除分组 */
export const deleteBlockGroup = (groupCode: string) => authRequest.GET('assetModel/deleteGroup', { groupCode });

/** 移动分组 */
export const moveBlockGroup = (data: Responsive.API.TemplateBatchMovePayload) => authRequest.POST('assetModel/move', data);

/**  复制模块 */
export const copyBlock = (id: string) => authRequest.GET(`assetModel/copy/${id}`);

/** 导出模块 */
export const downloadBlock = (ids: string[]) => authRequest.POST('assetModel/downloadModel', ids, {
  responseType: 'blob',
  ...cancelTimeout
});

/** 导入模块 */
export const importBlock = (params: Request.TemplateImportPayload) => authRequest.POSTFORM('assetModel/importModel', params, cancelTimeout);

/** 基于模板创建模块 */
export const createBlockByTemplate = (params: Request.PageCreateByTemplatePayload) => authRequest.POST('assetModel/addByTemplate', params);

/** 用户获取已上架模块列表 */
export const getUserBlockList = (params: Request.TemplateListPayload) => authRequest.POST('assetModel/pageList_cus', params);

/** 用户获取模块分组 */
export const getUserBlockGroup = () => authRequest.POST('assetModel/selectGroup_cus', {});
