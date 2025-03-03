import type { AxiosRequestConfig } from 'axios';
import fetch from './axiosInstance';

export interface ParamsType {
  [propName: string]: any;
}

type RequestType = (url: string, param?: ParamsType, config?: AxiosRequestConfig) => Promise<any>;

export const get: RequestType = (url, params, config = {}): Promise<any> => {
  return fetch.get(url, { params, ...config });
};

export const post: RequestType = (url, data, config = {}): Promise<any> => {
  return fetch.post(url, data, { ...config });
};

export const put: RequestType = (url, data, config = {}): Promise<any> => {
  return fetch.put(url, data, { ...config });
};

export { genCancelToken } from './axiosInstance';
const val = { get, post, put };
export default val;
