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
  // 'Content-Type': 'application/json'
};

export const defineApi = <Req, Res>(options: DefineApiOptions) => {
  const { method = 'GET', url, onError = Promise.reject } = options;

  type Requester = Req extends undefined
    ? (input?: null, opts?: RequestInit) => Promise<Res>
    : (input: Req, opts?: RequestInit) => Promise<Res>;

  // @ts-ignore
  const requester: Requester = async (input: any = {}, opts: RequestInit = {}) => {
    const { headers = defaultHeaders, ...others } = opts;
    let fullurl = url;
    const isGet = method === 'GET';
    if (isGet) {
      const search = new URLSearchParams(input);
      if (search.size > 0) {
        fullurl += url.includes('?') ? search.toString() : `?${search.toString()}`;
      }
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
