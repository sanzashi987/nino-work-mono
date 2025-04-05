export function queryStringify(params: Record<string, string>) {
  return new URLSearchParams(params).toString();
}

export function parseQuery(search:string) :Record<string, string> {
  const s = new URLSearchParams(search);
  return Object.fromEntries(s.entries());
}
