import { AssetsSearchParams, GroupCode, queryStringify } from '@canvas/utilities';
import { Request, Response } from '@canvas/types';
import type { ApiResponseType } from '@app/types';
import { cancelTimeout } from '@app/consts';

import defineApi from './impls';

/**
 * 获取资产分组列表
 * */
export const getAssetsGroupList = () => authRequest.POST('assets/selectGroup', {});

/**
 * 新增资产分组
 * */
export const addAssetsGroup = (params: Pick<Request.TemplateUpdateGroupPayload, 'groupName'>) => authRequest.POSTFORM('assets/addGroup', params);

/**
 * 删除资产分组
 * */
export const deleteAssetsGroup = (params: Pick<Request.TemplateUpdateGroupPayload, 'groupCode'>) => authRequest.GET('assets/deleteGroup', params);

/**
 * 重命名分组
 * */
export const updateGroupName = (params: Request.TemplateUpdateGroupPayload) => authRequest.POSTFORM('assets/updateGroupsName', params);

/**
 * 查询资产列表
 * */
export const getAssetsList = (
  params: AssetsSearchParams
): ApiResponseType<Response.AssetsMetaType[]> => authRequest.POST('assets/selectMyAssets', params);

/**
 * 重命名资产
 * */
export const renameAssets = (params: Pick<Response.AssetsMetaType, 'fileId' | 'fileName'>) => authRequest.POSTFORM('assets/updateMyAssetsName', params);

/**
 * 删除资产
 * */
export const deleteAssets = (ids: string[]) => authRequest.DELETE('assets/deleteAssets', null, { data: ids });

/**
 * 改变资产的分组
 * */
export const changeAssetsGroup = (params: Request.AssetsUpdateGroupPayload) => {
  const { fileIds, ...restParams } = params;
  return authRequest.POST(`assets/updateAssetsGroup?${queryStringify(restParams)}`, fileIds);
};

/** 上传资产(设计资产、数据源、字体、项目缩略图) */
export const uploadAssets = (
  file: File,
  params?: Request.AssetsUploadPayload
): ApiResponseType<Response.AssetBasicType> => authRequest.POSTFORM(
  'assets/upload',
  { groupCode: GroupCode.NOGROUP, ...params, file },
  cancelTimeout
);

/**
 * 根据文件id获取资产详情
 * */
export const getAssetsDetailById = (fileId: string): ApiResponseType<Response.AssetsMetaType> => authRequest.GET(`assets/detail?fileId=${fileId}`);

/**
 * 资产替换
 * */
export const replaceAssets = (params: Request.AssetsReplacePayload) => authRequest.POSTFORM('assets/replace', params, cancelTimeout);

/**
 * 资产导出
 * */
export const downloadAssets = (ids: string[]) => authRequest.POST('assets/loadAsset', ids, {
  responseType: 'blob',
  ...cancelTimeout
});

/**
 * 资产导入
 * */
export const importAssets = (params: Request.TemplateImportPayload) => authRequest.POSTFORM('assets/importAsset', params, cancelTimeout);
