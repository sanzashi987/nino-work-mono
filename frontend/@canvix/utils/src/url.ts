export function queryStringify(params: Record<string, string>) {
  return new URLSearchParams(params).toString();
}
