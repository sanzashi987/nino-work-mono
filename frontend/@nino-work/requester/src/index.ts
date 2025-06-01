import NinoFetchAbortController, { AbortConfig } from './abort';

/* eslint-disable no-restricted-syntax */
export type StandardResponse<T = any> = {
  msg: string;
  data: T;
  code: number;
};

export type ParserOption = {
  [K in keyof Body]: Body[K] extends (...args: any) => Promise<any> ? K : never;
}[keyof Body];

export type ParserConfig = {
  parser?: ParserOption;
};

export interface DefineApiOptions<Res, Out> extends AbortConfig, ParserConfig {
  method?: 'GET' | 'POST' | 'POSTFORM' | 'DELETE';
  url: string;
  onError?(payload?: any): Promise<any>;
  onResponse?(input: StandardResponse<Res>): Promise<Out>;
  headers?: Record<string, string>;
}

const defaultHeaders = {
  // Accept: 'application/json, text/html, */*'
  'Content-Type': 'application/json',
};

type PathMeta = {
  dynamic: boolean;
  name: string;
  optional: boolean;
};

function defaultOnResponse<Input, Output>(input: Input): Promise<Output> {
  if (typeof input === 'object') {
    if ('code' in input) {
      if (input.code === 0 && 'data' in input) {
        return Promise.resolve(input.data as Output);
      }
    }
  }
  return Promise.reject(input);
}

export interface NinoRequestInit extends RequestInit, AbortConfig, ParserConfig {}

export async function autoParseResponse(response: Response, parser?: ParserOption) {
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
  }
  if (contentType.includes('text/')) {
    return response.text(); // 解析文本数据
  }
  if (contentType.includes('image/') || contentType.includes('application/pdf')) {
    return response.blob(); // 解析二进制数据（如图片、PDF）
  }
  if (
    contentType.includes('multipart/form-data') ||
    contentType.includes('application/x-www-form-urlencoded')
  ) {
    return response.formData(); // 解析表单数据
  }
  if (contentType.includes('application/octet-stream')) {
    return response.arrayBuffer(); // 解析原始二进制数据
  }
  // 如果 Content-Type 未知，默认返回json
  return response.json();
}

export const defineApi = <Req, Res = void, Out = Res>(
  options: DefineApiOptions<Res, Out>,
  mock?: Res
) => {
  const {
    method = 'GET',
    url,
    onError = Promise.reject,
    headers = defaultHeaders,
    onResponse = defaultOnResponse,
    timeout,
  } = options;
  const pathMetas = url.split('/').map(param => {
    const meta: PathMeta = { dynamic: false, optional: false, name: param };
    let name = param;
    if (param.startsWith(':')) {
      name = name.slice(1);
      meta.dynamic = true;
      if (param.endsWith('?')) {
        name = name.slice(0, -1);
        meta.optional = true;
      }
    }
    meta.name = name;
    return meta;
  });

  const hasQuery = pathMetas.at(-1).name.includes('?');

  type Requester = Req extends undefined
    ? (input?: null, opts?: NinoRequestInit) => Promise<Res>
    : (input: Req, opts?: NinoRequestInit) => Promise<Res>;

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-expect-error
  const requester: Requester = async (
    input: Record<string, any> = {},
    opts: NinoRequestInit = {}
  ) => {
    const { headers: overrideHeaders, signal, timeout: runtimeTimeout, ...others } = opts;
    const inputNext = { ...input };

    const dynamicPaths: string[] = [];
    for (const meta of pathMetas) {
      if (!meta.dynamic) {
        dynamicPaths.push(meta.name);
      } else {
        const val = inputNext[meta.name];
        if (val) {
          delete inputNext[meta.name];
          dynamicPaths.push(val);
        } else if (!meta.optional) {
          throw new Error('required params not given');
        }
      }
    }

    let fullurl = dynamicPaths.join('/');
    let body: BodyInit;
    const isGet = method === 'GET';

    if (isGet) {
      const search = new URLSearchParams(input);
      fullurl += hasQuery ? search.toString() : `?${search.toString()}`;
    } else if (method === 'POSTFORM') {
      delete headers['Content-Type'];
      const formData = new FormData();
      for (const [key, value] of Object.entries(inputNext)) {
        if (Array.isArray(value) && value[0] instanceof File) {
          delete inputNext[key];
          for (const file of value) {
            formData.append(`${key}[]`, file, file.name);
          }
          // eslint-disable-next-line no-continue
          continue;
        }
        if (value !== undefined) {
          formData.append(key, value);
        }
      }
      body = formData;
    } else {
      body = JSON.stringify(inputNext);
    }

    const signalWithDefault =
      signal ?? new NinoFetchAbortController({ timeout: runtimeTimeout ?? timeout }).signal;

    const res = await fetch(fullurl, {
      headers: { ...headers, ...overrideHeaders },
      method: method === 'POSTFORM' ? 'POST' : method,
      ...others,
      signal: signalWithDefault,
      body,
    });

    if (res.redirected && res.headers.get('Content-Type')?.includes('text/html')) {
      window.location.href = res.url;
      return onError(`Redirected to ${res.url}`);
    }

    if (!res.ok) {
      // eslint-disable-next-line prefer-promise-reject-errors
      return onError(`Response Status Error:${res.status}`);
    }

    // const data = await res.json() as StandardResponse<Res>;
    const data = (await autoParseResponse(
      res,
      options.parser ?? opts.parser
    )) as StandardResponse<Res>;

    return onResponse(data).catch(onError);
  };

  if (mock) {
    return (() => Promise.resolve(mock)) as unknown as Requester;
  }

  return requester;
};
