import { ApiResponseType, RES_NULL } from '@canvas/utilities';
import type { Request, Responsive, Response } from '@canvas/types';
import { cancelTimeout } from '@app/consts';
import { authRequest, basicRequest } from './axios';

/** ------ 分组查询 ------- */
// 查询大屏分组列表
export const getGroupList = (params: Request.PageGroupPayload) => authRequest.POST('group/list', params);

// 新增大屏分组
export const addGroup = (params: Request.PageGroupCreatePayload) => authRequest.POST('group/create', params);

// 删除大屏分组
export const deleteGroup = (data: string[]) => authRequest.DELETE('group/delete', null, { data });

// 更新大屏分组信息（名称）
export const updateGroup = (params: Request.PageGroupUpdatePayload) => authRequest.POST('group/update', params);

/** ------ 大屏查询 ------- */

/** 查询页面列表 */
export const getScreenList = (
  params: Request.PageListPayload
): ApiResponseType<Response.ScreenResponseType[]> => authRequest.POST('screen-operation/list', params);

// 批量更新大屏分组
export const updateScreenGroup = (params: Responsive.API.TemplateBatchMovePayload) => authRequest.POST('screen-operation/move', params);

/** 复制页面 */
export const copyScreen = (id: string) => authRequest.GET(`screen-operation/copy/${id}`);

/** 删除页面 */
export const deleteScreen = (ids: string[]) => authRequest.DELETE('screen-operation/delete', null, { data: ids });

// 发布页面模板
export const publishScreen = (params: Responsive.API.PagePublishPayload) => authRequest.POST('screen-operation/publish', params);

/** 创建空白大屏页面 */
export const createScreen = (params: Responsive.API.PageCreatePayload) => authRequest.POST('screen-operation/create', params);

// 导出大屏
export const downloadScreen = (ids: string[]) => authRequest.POST('screen-operation/downloadScreen', ids, {
  responseType: 'blob',
  ...cancelTimeout
});

// 导入大屏
export const importScreen = (params: Request.TemplateImportPayload) => authRequest.POSTFORM('screen-operation/importScreen', params, cancelTimeout);

// 编译大屏成app
export const compileScreen = (code: string) => authRequest.POST(
  `screen-operation/downloadApp?code=${code}&type=h5`,
  {},
  {
    responseType: 'blob',
    ...cancelTimeout
  }
);

/** 获取大屏详细信息 （发布页，不鉴权） */
export const getPublishProjectInfo = (
  id: string,
  params?: { secret: string; publishToken: string; parentId: string }
) => basicRequest.GET(`facade/get-screen/${id}`, params);

/** 修改（保存）大屏 */
export const updateScreen = (params: Responsive.API.PageUpdatePayload) => authRequest.POST('/screen-operation/update', params);

/** 获取大屏详细信息 */
export const getScreenInfo = (id: string): ApiResponseType<Responsive.API.ScreenResponseType> => authRequest.GET(`/screen-operation/info/${id}`);

// 获取引用面板节点交互配置
export const getRefNodeInteraction = (data: { refId: string; screenId: string }) => Promise.resolve(RES_NULL);

export const createPageByTemplate = (params: Request.PageCreateByTemplatePayload) => authRequest.POST('/screen-operation/addByTemplate', params);
