import type { ApiResponseType } from '@canvas/utilities';
import type { Response, Request, Responsive } from '@canvas/types';
import { cancelTimeout } from '@app/consts';

/** 新增自定义组件分组 */
export const addCustomComGroup = (params: Pick<Request.PageGroupUpdatePayload, 'name'>) => authRequest.POST('component-cus-group/create', params);

/** 删除自定义组件分组 */
export const deleteCustomComGroup = (data: string[]) => authRequest.DELETE('component-cus-group/delete', null, { data });

/** 修改自定义组件分组  */
export const updataCustomCom = (params: Request.PageGroupUpdatePayload) => authRequest.POST('component-cus-group/update', params);

/**
 * 查询自定义组件分组列表
 * @number state 0 - 已上架 1 - 未上架
 */
export const getCustomGroupList = (
  params?: Request.ComponentGroupListPayload
): ApiResponseType<Response.GroupResponseType> => authRequest.POST('component-cus-group/list', params || {});

/** 上传自定义组件  */
export const addNewCustomCom = (params: Request.TemplateImportPayload) => authRequest.POSTFORM('component-cus/create', params, cancelTimeout);

/** 组件上下架 */
export const shelfCom = (params: Request.TemplateStateUpdatePayload) => authRequest.POST('component-cus/state', params);

/** 导出自定义组件 */
export const exportCom = (ids: string[]) => authRequest.POST('component-cus/load', ids, { responseType: 'blob', ...cancelTimeout });

/** 更新自定义组件 */
export const updateCom = (params: Request.ComponentUpdatePayload) => authRequest.POSTFORM('component-cus/update', params, cancelTimeout);

/** 删除自定义组件 */
export const deleteCom = (data: string[]) => authRequest.DELETE('component-cus/delete', null, { data });

/** 查询自定义组件列表 */
export const getComList = (params: Request.ComponentListPayload) => authRequest.POST('component-cus/pageList', params);

/** 查询已上架自定义组件列表 */
export const getOnShelfCustomComponentList = (): ApiResponseType<
Response.CustomComponentResponseType[]
> => authRequest.POST('component-cus/list', {});

/** 更新自定义组件分组 */
export const batchUpdateCom = (params: Responsive.API.TemplateBatchMovePayload) => authRequest.POST('component-cus/move', params);
