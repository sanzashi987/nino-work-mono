import { type Request } from '@canvas/types';
import { cancelTimeout } from '@app/consts';

/** 获取平台组件列表，包含已上架和未上架 */
export const getComList = (params: Omit<Request.TemplateListPayload, 'groupCode'>) => authRequest.POST('view-component-static/listOnOrOff', params);

/**
 * 批量上架组件
 * @param params
 * @returns
 */
export const onShelf = (params: Array<{ name: string; version: string }>) => authRequest.POST('component-operation/onShelf', params);

/**
 * 批量下架组件
 * @param params
 * @returns
 */
export const offShelf = (params: Array<{ name: string; version: string }>) => authRequest.POST('component-operation/offShelf', params);

/**
 * 导入组件
 * @param params
 * @returns
 */
export const importComponent = (params: Request.ComponentImportPayload) => authRequest.POSTFORM('component-operation/importComponent', params, cancelTimeout);

/**
 * 删除组件
 * @param params
 * @returns
 */
export const deleteComponent = (ids: string[]) => authRequest.DELETE('component-operation/delete', null, { data: ids });
