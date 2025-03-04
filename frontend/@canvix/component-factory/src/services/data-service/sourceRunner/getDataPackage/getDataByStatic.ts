import type { IdentifierSource, GetValueEntryType } from '@canvix/shared';
import { getErrorInfo } from './utils';
import { canvasApiService, RES_ERR_NOT_FOUND } from '../../constants';
import { get } from '../../requestService';

const staticDataCache: Record<string, any> = {};
const getDebugUrl = (params: IdentifierSource) => {
  const { sourceName, name, version } = params;
  const detailStr = localStorage.getItem('debugInfo');
  const detail = detailStr ? JSON.parse(detailStr) : {};
  const { appUrl } = detail;
  return `${appUrl}/${name}/${version}/static/${sourceName}.json`;
};

const getDefaultUrl = (params: IdentifierSource) => {
  const { sourceName, name, version, user, isDebugger } = params;
  if (isDebugger) return getDebugUrl(params);
  const baseUrl = `${canvasApiService}/canvas-pro-mobile/V1/files/file`;
  const userPath = user ? `/${user}` : '';
  return `${baseUrl}${userPath}/mobile/${name}/${version}/static/${sourceName}.json`;
};
const getFileUrl = (params: IdentifierSource, source: Record<string, any>) => {
  const { name, version, user, projectId, isDebugger } = params;
  if (isDebugger) return getDebugUrl(params);
  return `${canvasApiService}/canvas-pro-mobile/V1/files/file/${
    source.fileId
  }?code=${name}&comVersion=${version}${
    user ? '&cType=1' : ''
  }&screenId=${projectId}`;
};

const getStaticData: GetValueEntryType = async (source, identifier, config) => {
  // 有fileId取修改之后的数据，否则取默认静态数据
  const url = source?.fileId ? getFileUrl(identifier, source) : getDefaultUrl(identifier);
  try {
    if (!staticDataCache[url]) {
      staticDataCache[url] = get(url, {}, config);
    }
    const res = await staticDataCache[url];
    return { needUpdate: true, output: res };
  } catch (e) {
    // console.log(e);
    delete staticDataCache[url];
    return { needUpdate: false, output: [], error: e ? getErrorInfo(e) : RES_ERR_NOT_FOUND };
  }
};

export default getStaticData;
