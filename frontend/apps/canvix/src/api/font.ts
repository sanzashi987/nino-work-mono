import { type AssetsSearchParams } from '@canvas/utilities';
import type { Request, Response } from '@canvas/types';
import type { ApiResponseType } from '@app/types';
import { cancelTimeout } from '@app/consts';

/** 新增字体资产 */
export const addFont = (params: Request.FontCreatePayload) => authRequest.POSTFORM('font/fontAdd', params, cancelTimeout);

/** 更新字体资产 */
export const updateFont = (params: Request.FontUpdatePayload) => authRequest.POSTFORM('font/fontUpdate', params, cancelTimeout);

/** 获取字体列表(分页) */
export const getFontList = (params: AssetsSearchParams) => authRequest.POST('font/fontList', params);

/** 获取字体列表（全部） */
export const getAllFont = (
  params: Request.FontAllListPayload
): ApiResponseType<Response.FontResponseType[]> => authRequest.POSTFORM('font/fonts', params);

/** 删除字体资产 */
export const deleteFonts = (ids: string[]) => authRequest.DELETE('font/fontDelete', null, { data: ids });
