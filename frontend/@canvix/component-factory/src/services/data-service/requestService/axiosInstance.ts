import axios from 'axios';

const fetch = axios.create({
  timeout: 60000,
  responseType: 'json',
  headers: {
    'Content-Type': 'application/json',
  },
});

fetch.interceptors.request.use(
  (config: any) => {
    // 数据源发布页面可用，不携带auth
    // config.headers.auth = JSON.parse(localStorage.getItem('userInfo') || '{}').token;
    return config;
  },
  (error: any) => {
    return Promise.reject(error);
  },
);

fetch.interceptors.response.use(
  (response: any) => {
    return response.data;
  },
  (error: any) => {
    const errorString = error.toString();
    if (errorString.includes('timeout')) {
      return Promise.reject({ errCode: 408, message: '请求超时' });
    }
    return Promise.reject(error);
  },
);
export default fetch;

export const genCancelToken = () => {
  return axios.CancelToken.source();
};
