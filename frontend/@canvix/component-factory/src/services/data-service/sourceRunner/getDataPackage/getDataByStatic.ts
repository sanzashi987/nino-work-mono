import type { IdentifierSource, GetValueEntryType } from '@canvix/shared';
import { getDebugUrl, getDefaultStaticSource, getErrorInfo, getFileUrl } from './utils';
import { RES_ERR_NOT_FOUND } from '../../constants';
import { get } from '../../requestService';

const staticDataCache: Record<string, any> = {};

const getDefaultUrl = (params: IdentifierSource) => {
  if (params.isDebugger) return getDebugUrl(params);
  return getDefaultStaticSource(params);
};
const getModifiedStaticUrl = (params: IdentifierSource, source: Record<string, any>) => {
  const { projectId, isDebugger } = params;
  if (isDebugger) return getDebugUrl(params);
  return getFileUrl(source.fileId, projectId);
};

const getStaticData: GetValueEntryType = async (source, identifier, config) => {
  // 有fileId取修改之后的数据，否则取默认静态数据
  const url = source?.fileId ? getModifiedStaticUrl(identifier, source) : getDefaultUrl(identifier);
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
