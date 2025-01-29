/* eslint-disable no-restricted-syntax */
export type StandardResponse<T> = {
  msg: string
  data: T
  code: number
};

export type DefineApiOptions = {
  method?: 'GET' | 'POST',
  url: string
  onError?(input?: any): Promise<any>
};

const defaultHeaders = {
  // Accept: 'application/json, text/html, */*'
  'Content-Type': 'application/json'
};

type PathMeta = {
  dynamic: boolean
  name: string
  optional:boolean
};

export const defineApi = <Req, Res>(options: DefineApiOptions) => {
  const { method = 'GET', url, onError = Promise.reject } = options;
  const pathMetas = url.split('/').map((param) => {
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

  type Requester = Req extends undefined
    ? (input?: null, opts?: RequestInit) => Promise<Res>
    : (input: Req, opts?: RequestInit) => Promise<Res>;

  // @ts-ignore
  const requester: Requester = async (input: any = {}, opts: RequestInit = {}) => {
    const { headers = defaultHeaders, ...others } = opts;
    const inputNext = { ...input };

    const dynamicPaths:string[] = [];
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

    const isGet = method === 'GET';

    if (isGet) {
      const search = new URLSearchParams(input);
      fullurl += url.includes('?') ? search.toString() : `?${search.toString()}`;
    }
    const res = await fetch(fullurl, {
      headers,
      method,
      ...others,
      body: isGet ? undefined : JSON.stringify(input)
    });

    if (res.redirected && res.headers.get('Content-Type')?.includes('text/html')) {
      window.location.href = res.url;
      return onError();
    }

    if (!res.ok) {
      // eslint-disable-next-line prefer-promise-reject-errors
      return onError(`Response Status Error:${res.status}`);
    }

    const data = await res.json() as StandardResponse<Res>;

    if (data.code !== 0) {
      return onError(data?.msg);
    }

    return data.data;
  };
  return requester;
};
