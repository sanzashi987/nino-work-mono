import type { IdentifierSource, ApiReturnType, GetValueEntryType } from '@/types';
import { getFileUrl, getErrorInfo } from './utils';
import { post } from '../../requester';

const queryDataBySourceId = (
  sourceId: string,
  identifier: IdentifierSource,
  config: object = {}
): Promise<any> => post(
  getFileUrl(sourceId, identifier.projectId),
  {},
  config
);

async function getData(
  params: Record<string, any>,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  identifier: IdentifierSource,
  config: object
): Promise<ApiReturnType> {
  try {
    const res = await queryDataBySourceId(params.sourceId, identifier, config);
    const { code, data, msg } = res;
    if (code !== 0 || !data) throw new Error(msg);
    return { needUpdate: true, output: data };
  } catch (e) {
    console.log(e);
    return { needUpdate: false, output: [], error: getErrorInfo(e) };
  }
}

const getDataByFile: GetValueEntryType<{ sourceId: string }> = async (
  source,
  identifier,
  config
) => {
  const { sourceId } = source;
  if (!sourceId) return { needUpdate: false, output: [] };
  const params = { sourceId };
  return getData(params, identifier, config);
};

export default getDataByFile;
