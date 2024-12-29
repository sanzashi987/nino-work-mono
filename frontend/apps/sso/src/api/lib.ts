export type StandartResponse<T> = {
  msg: string
  data: T
  code: number
};

type DefineApiOptions = {
  method?: 'GET' | 'POST',
  url: string
};

const defaultHeaders = {
  Accept: 'application/json',
  'Content-Type': 'application/json'
};

export const defineApi = <Req extends Record<string, any>, Res>(options: DefineApiOptions) => {
  const { method = 'GET', url } = options;

  return async (input?: Req, opts: RequestInit = {}) => {
    const { headers = defaultHeaders, ...others } = opts;
    let fullurl = url;
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

    if (!res.ok) {
      return Promise.reject(new Error(`Response Status Error:${res.status}`));
    }

    const data = await res.json() as StandartResponse<Res>;

    if (data.code !== 0) {
      return Promise.reject(data);
    }

    return data.data;
  };
};
