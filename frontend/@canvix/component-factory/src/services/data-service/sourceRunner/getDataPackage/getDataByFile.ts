import type { IdentifierSource, ApiReturnType, GetValueEntryType } from '@canvix/shared';
import { getErrorInfo } from './utils';
import { canvasApiService } from '../../constants';
import { post } from '../../requestService';

const queryDataBySourceId = (
  sourceId: string,
  identifier: IdentifierSource,
  config: object = {}
): Promise<any> => post(
  `${canvasApiService}/canvas-pro-mobile/V1/source-connect/getData?sourceId=${sourceId}&screenId=${identifier.projectId}`,
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
    const { resultCode, data, resultMessage } = res;
    if (resultCode !== 0 || !data) throw new Error(resultMessage);
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
