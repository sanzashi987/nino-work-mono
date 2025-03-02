import { Response, Request } from '@canvas/types';
import { ApiResponseType } from '@app/types';

/** 获取数据源列表 */
export const getConnectList = (
  params: Request.DataSourceSourceListPayload
): ApiResponseType<Response.ConnectResponseType[]> => authRequest.POST('jdbc-connect-template/list-all', params);

/** 分页获取数据源列表 */
export const getConnectListByPage = (
  params: Request.DataSourceListPayload
): ApiResponseType<Response.ConnectResponseType[]> => authRequest.POST('jdbc-connect-template/list-page', params);

/** 删除数据源链接 */
export const deleteConnect = (sourceIds: string[]) => authRequest.DELETE('jdbc-connect-template/delete', null, { data: sourceIds });

/** 更新数据源链接 */
export const updateConnect = (
  params: Response.ConnectUpdateType
): ApiResponseType<Response.ConnectResponseType> => authRequest.POSTFORM('jdbc-connect-template/update', params);

/** 新增数据源链接 */
export const createConnect = (
  params: Response.ConnectUpdateType
): ApiResponseType<Response.ConnectResponseType> => authRequest.POSTFORM('jdbc-connect-template/create', params);

/** 获取链接详情 */
export const getConnectDetail = (sourceId: string): ApiResponseType<Response.ConnectResponseType> => authRequest.GET(`jdbc-connect-template/info/${sourceId}`);

/** 数据源ip查询 */
export const getDataListByIp = (
  params: Request.DataSourceSearchPayload
): ApiResponseType<Response.ConnectResponseType[]> => authRequest.POST('jdbc-connect-template/searchByIp', params);

/** 数据源ip替换 */
export const replaceIp = (params: Request.DataSourceReplacePayload) => authRequest.POST('jdbc-connect-template/replaceIp', params);

/** 获取文件上传的数据 */
export const getFileDataBySourceId = (params: Request.DataSourceFileDataPayload) => authRequest.POST('source-connect/getData', null, { params });
