import type { ApiReturnType, GetValueEntryType, IdentifierSource, RequestApi } from '@canvix/shared';
import { getErrorInfo } from './utils';
import { canvasApiService } from '../../constants';
import requestService, { post } from '../../requestService';

type RequestType = 'get' | 'put' | 'post';
const defaultApiParams: Omit<RequestApi, 'sourceId'> = {
  body: {},
  headers: {},
  querys: {},
  path: {}
};

const requestApi = (
  params: RequestApi,
  identifier: IdentifierSource,
  config: any = {}
): Promise<any> => post(
  `${canvasApiService}/canvas-pro-mobile/V1/facade/request-api?screenId=${identifier.projectId}`,
  params,
  config
);

async function getData(
  source: RequestApi,
  identifier: IdentifierSource,
  config: any
): Promise<ApiReturnType> {
  try {
    const params = {
      ...defaultApiParams,
      ...source
    };
    const res = await requestApi(params, identifier, config);
    const { resultCode, data, resultMessage } = res;
    if (resultCode !== 0 || !data) throw new Error(resultMessage);
    const { data: resData, proxy } = data;
    if (proxy) {
      // 服务器代理请求
      return { needUpdate: true, output: resData };
    }
    try {
      // 前端发送请求
      const config = JSON.parse(resData);
      const { method, url } = config || {};
      const requestType: RequestType = method.toLowerCase();
      const data = requestType === 'get' ? {} : params.body;
      const apiRes = await requestService[requestType]?.(url, data, {
        ...config,
        headers: { ...config.headers, ...params.headers },
        // data: params.body ?? {},
        params: params.querys ?? {}
      });
      return { needUpdate: true, output: apiRes };
    } catch (e) {
      console.log(e);
      return { needUpdate: false, output: [], error: getErrorInfo(e) };
    }
  } catch (e) {
    console.log(e);
    return { needUpdate: false, output: [], error: getErrorInfo(e) };
  }
}

const getDataByApi: GetValueEntryType<{ sourceId: string }> = async (
  source,
  identifier,
  config
) => {
  const { sourceId } = source;
  if (!sourceId) return { needUpdate: false, output: [] };
  return getData(source, identifier, config);
};

export default getDataByApi;

export { defaultApiParams };
