import type { AxiosRequestConfig } from 'axios';
import { getErrorInfo } from './utils';
import { canvasApiService } from '../../constants';
import { post } from '../../requestService';
import type { GetValueEntryType, ApiReturnType, IdentifierSource } from '../../types';

const queryDataBySourceId = (
  sourceId: string,
  identifier: IdentifierSource,
  config: AxiosRequestConfig = {},
): Promise<any> => {
  return post(
    `${canvasApiService}/canvas-pro-mobile/V1/source-connect/getData?sourceId=${sourceId}&screenId=${identifier.dashboardId}&rak-token=${identifier.rakToken}`,
    {},
    config,
  );
};

async function getData(
  params: Record<string, any>,
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  identifier: IdentifierSource,
  config: AxiosRequestConfig,
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
  config,
) => {
  const { sourceId } = source;
  if (!sourceId) return { needUpdate: false, output: [] };
  const params = { sourceId };
  return getData(params, identifier, config);
};

export default getDataByFile;
