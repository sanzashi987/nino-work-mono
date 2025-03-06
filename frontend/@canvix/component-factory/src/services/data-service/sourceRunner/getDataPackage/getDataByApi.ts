import type { ApiReturnType, GetValueEntryType, IdentifierSource, RequestApi } from '@canvix/shared';
import { getApiUrl, getErrorInfo } from './utils';
import requestService, { post, RequestConfig } from '../../requestService';

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
  config: RequestConfig = {}
): Promise<any> => post(
  getApiUrl(identifier.projectId),
  params,
  config
);

async function getData(
  { parser, ...source }: RequestApi,
  identifier: IdentifierSource,
  config: RequestConfig
): Promise<ApiReturnType> {
  try {
    const params = {
      ...defaultApiParams,
      ...source
    };
    const res = await requestApi(params, identifier, { ...config, parser });
    const { code, data, msg } = res;
    if (code !== 0 || !data) throw new Error(msg);
    const { data: resData, proxy } = data;
    if (proxy) {
      // 服务器代理请求
      return { needUpdate: true, output: resData };
    }
    try {
      // 前端发送请求
      const conf = JSON.parse(resData);
      const { method, url } = conf || {};
      const requestType: RequestType = method.toLowerCase();
      const payload = requestType === 'get' ? {} : params.body;
      const apiRes = await requestService[requestType]?.(url, payload, {
        ...conf,
        headers: { ...conf.headers, ...params.headers },
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
