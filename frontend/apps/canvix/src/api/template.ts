import { Request, Responsive } from '@canvas/types';
import { cancelTimeout } from '@app/consts';

/** 查询模板分组列表 */
export const getGroupList = () => authRequest.POST('screenTemplate/selectGroup', {});

/** 新增模板分组 */
export const addGroup = (groupName: string) => authRequest.POSTFORM('screenTemplate/addGroup', { groupName });

/** 删除模板分组 */
export const deleteGroup = (code: string) => authRequest.GET('screenTemplate/deleteGroup', { groupCode: code });

/** 重命名模板分组 */
export const updateGroup = (params: Request.TemplateUpdateGroupPayload) => authRequest.POSTFORM('screenTemplate/updateGroupsName', params);

/** 管理员查询模板列表 */
export const getTemplateList = (params: Request.TemplateListPayload) => authRequest.POST('screenTemplate/pageList', params);

/** 更新模板 */
export const updateTemplate = (params: Responsive.API.TemplateUpdatePayload) => authRequest.POST('screenTemplate/update', params);

/** 批量更新大屏分组 */
export const updateTemplateGroup = (params: Responsive.API.TemplateBatchMovePayload) => authRequest.POST('screenTemplate/move', params);

// 复制大屏模板
export const copyTemplate = (id: string) => authRequest.GET(`screenTemplate/copy/${id}`);

/** 删除大屏模板 */
export const deleteTemplate = (ids: string[]) => authRequest.DELETE('screenTemplate/delete', null, { data: ids });

/** 创建大屏模板页面 */
export const createTemplate = (params: Request.TemplateCreatePayload) => authRequest.POST('screenTemplate/add', { ...params });

/** 导出模板 */
export const downloadTemplate = (ids: string[]) => authRequest.POST('screenTemplate/downloadScreen', ids, {
  responseType: 'blob',
  ...cancelTimeout
});

/** 导入模板 */
export const importTemplate = (params: Request.TemplateImportPayload) => authRequest.POSTFORM('screenTemplate/importScreen', params, cancelTimeout);

/** 获取模板详情 */
export const getTemplateDetail = (params: { id: string }) => authRequest.GET('screenTemplate/detail', params);

/** 用户获取模板列表 */
export const getUserTemplateList = (params: Request.TemplateListPayload) => authRequest.POST('screenTemplate/pageList_cus', params);

/** 用户获取模板分页 */
export const getUserTemplateGroups = () => authRequest.POST('screenTemplate/selectGroup_cus', {});

/** 获取用户类型列表 */
export const getUserTypes = () => authRequest.GET('screenTemplate/userTypes');

/** 模板上下架 */
export const updateTemplateState = (params: Request.TemplateStateUpdatePayload) => authRequest.POST('screenTemplate/state', params);
