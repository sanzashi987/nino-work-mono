import fetch from './fetch';

export interface ParamsType {
  [propName: string]: any;
}

type RequestType = (url: string, param?: ParamsType, config?: {}) => Promise<any>;

export const get: RequestType = (url, params, config = {}): Promise<any> => fetch.get(url, { params, ...config });

export const post: RequestType = (url, data, config = {}): Promise<any> => fetch.post(url, data, { ...config });

export const put: RequestType = (url, data, config = {}): Promise<any> => fetch.put(url, data, { ...config });

export { genCancelToken } from './fetch';
const val = { get, post, put };
export default val;
