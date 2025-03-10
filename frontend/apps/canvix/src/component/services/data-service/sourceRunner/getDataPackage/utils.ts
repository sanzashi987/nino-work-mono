import { IdentifierSource } from '@/types';

export const getErrorInfo = (error: DOMException & { msg?: string }) => ({
  name: error?.name,
  status: error?.code,
  message: error?.msg || error?.message
  // data: error?.response?.data
});

const origin = '/backend/canvix/v1';

export function getFileUrl(sourceId:string, projectId:string) {
  return `${origin}/data-source/file?sourceId$=${sourceId}&projectId=${projectId}`;
}

export function getApiUrl(projectId:string) {
  return `${origin}/data-source/api?projectId=${projectId}`;
}

export function getDefaultStaticSource(params: IdentifierSource) {
  const { sourceName, name, version, user, projectId } = params;
  const search = new URLSearchParams({
    sourceName, name, version, user: user ?? undefined, projectId
  } as Record<string, string>).toString();

  return `${origin}/data-source/static?${search}`;
}

export const getDebugUrl = (params: IdentifierSource) => {
  const { sourceName, name, version } = params;
  const detailStr = localStorage.getItem('debugInfo');
  const detail = detailStr ? JSON.parse(detailStr) : {};
  const { appUrl } = detail;
  return `${appUrl}/${name}/${version}/static/${sourceName}.json`;
};
