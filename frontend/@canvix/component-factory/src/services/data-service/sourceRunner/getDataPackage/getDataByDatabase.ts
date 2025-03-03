import type { AxiosRequestConfig } from 'axios';
import { getErrorInfo } from './utils';
import { canvasApiService } from '../../constants';
import { post } from '../../requestService';
import type {
  GetValueEntryType,
  ApiReturnType,
  PostgreSqlConfig,
  MySqlOracleConfig,
  IdentifierSource,
} from '../../types';

const queryDataByDatabase = (
  params: Record<string, any>,
  identifier: IdentifierSource,
  config: AxiosRequestConfig = {},
): Promise<any> => {
  return post(
    `${canvasApiService}/canvas-pro-mobile/V1/source-connect/find-table-content?screenId=${identifier.dashboardId}`,
    params,
    config,
  );
};

async function getDataByDatabase(
  params: Record<string, any>,
  identifier: IdentifierSource,
  config: AxiosRequestConfig,
): Promise<ApiReturnType> {
  try {
    const res = await queryDataByDatabase(params, identifier, config);
    const { resultCode, data, resultMessage } = res;
    if (resultCode !== 0 || !data) throw new Error(resultMessage);
    return { needUpdate: true, output: data };
  } catch (e) {
    console.log(e);
    return { needUpdate: false, output: [], error: getErrorInfo(e) };
  }
}

export const getDataByPGSQL: GetValueEntryType<PostgreSqlConfig> = async (
  source,
  identifier,
  config,
) => {
  const { dbName, sourceId, sqlContent, patternName } = source;
  if (!dbName || !sourceId || !sqlContent || !patternName) return { needUpdate: false, output: [] };
  const params = { dbName, sourceId, sqlContent, patternName };
  return getDataByDatabase(params, identifier, config);
};

export const getDataByCommonSQL: GetValueEntryType<MySqlOracleConfig> = async (
  source,
  identifier,
  config,
) => {
  const { dbName, sourceId, sqlContent } = source;
  if (!dbName || !sourceId || !sqlContent) return { needUpdate: false, output: [] };
  const params = { dbName, sourceId, sqlContent };
  return getDataByDatabase(params, identifier, config);
};
