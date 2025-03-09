import { ParserOption, RequestApi } from '@/types';

/* eslint-disable prefer-promise-reject-errors */
export type ParamsType = Omit<RequestApi, 'sourceId' | 'parser'>;

type RequestType = (url: string, param?: ParamsType, config?: RequestConfig) => Promise<any>;

export type RequestConfig = RequestInit & {
  parser?: ParserOption;
};

async function autoParseResponse(response: Response, parser?: ParserOption) {
  if (parser) {
    return response[parser]();
  }
  // 获取 Content-Type 头
  const contentType = response.headers.get('Content-Type');

  if (!contentType) {
    return response.json();
  }

  // 根据 Content-Type 选择解析方法
  if (contentType.includes('application/json')) {
    return response.json(); // 解析 JSON 数据
  } if (contentType.includes('text/')) {
    return response.text(); // 解析文本数据
  } if (contentType.includes('image/') || contentType.includes('application/pdf')) {
    return response.blob(); // 解析二进制数据（如图片、PDF）
  } if (contentType.includes('multipart/form-data') || contentType.includes('application/x-www-form-urlencoded')) {
    return response.formData(); // 解析表单数据
  } if (contentType.includes('application/octet-stream')) {
    return response.arrayBuffer(); // 解析原始二进制数据
  }
  // 如果 Content-Type 未知，默认返回json
  return response.json();
}

function commonHandler(res: Response, parser?: ParserOption) {
  if (res.ok) {
    return autoParseResponse(res, parser);
  }
  return Promise.reject({
    name: 'Response status error',
    status: res.status,
    message: 'Status code error'
  });
}

function commonHandleUrl(url: string, params: ParamsType, config: RequestConfig) {
  const pathMetas = url.split('/').map((param) => {
    let name = param;
    if (param.startsWith('{') && param.endsWith('}')) {
      const key = param.slice(1, -1);
      name = params?.path?.[key] || 'undefined';
    }
    return name;
  });

  const hasQuery = pathMetas.at(-1)?.includes('?');
  const search = new URLSearchParams(params?.querys);

  const headers = { ...config.headers, ...params?.headers };

  let fullUrl = pathMetas.join('/');
  fullUrl += hasQuery ? search.toString() : `?${search.toString()}`;
  return { url: fullUrl, headers };
}

export const get: RequestType = (url, params = {}, config = {}): Promise<any> => {
  const { url: fullUrl, headers } = commonHandleUrl(url, params, config);

  return fetch(fullUrl, { ...config, headers, method: 'GET' }).then((res) => commonHandler(res, config.parser));
};

export const post: RequestType = (url, params = {}, config = {}): Promise<any> => {
  const { url: fullUrl, headers } = commonHandleUrl(url, params, config);

  const body = JSON.stringify(params.body);
  return fetch(fullUrl, {
    ...config,
    headers,
    body,
    method: 'POST'
  }).then((res) => commonHandler(res, config.parser));
};

export const put: RequestType = (url, params = {}, config = {}): Promise<any> => {
  const { url: fullUrl, headers } = commonHandleUrl(url, params, config);

  const body = JSON.stringify(params.body);
  return fetch(fullUrl, {
    ...config,
    headers,
    body,
    method: 'PUT'
  }).then((res) => commonHandler(res, config.parser));
};

const val = { get, post, put };
export default val;
export { default as RefreshTimer } from './refreshTimer';
