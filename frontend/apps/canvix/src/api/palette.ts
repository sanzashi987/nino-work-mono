import type {
  PaletteListFromResponse,
  ApiResponseType,
  RequestCreatePalettePayload,
  UnsafeRequestCreatePalettePayload,
  RequestUpdatePalettePayload
} from '@app/types';

/** 查询所有调色板列表 */
export const getPaletteList = (): ApiResponseType<PaletteListFromResponse> => authRequest.POST('/system-theme/list', {});

/** 修改单个调色板 */
export const updatePalette = (params: RequestUpdatePalettePayload) => authRequest.POST('/system-theme/update', params);

/** 删除调色板 */
export const deletePalette = (data: number[]) => authRequest.DELETE('/system-theme/delete', '', { data });

/** 新增单个调色板 */
export const createPalette = (params: RequestCreatePalettePayload) => unsafe_createPalette({ ...params, flag: 1 });

export const unsafe_createPalette = (params: UnsafeRequestCreatePalettePayload) => authRequest.POST('/system-theme/create', params);
