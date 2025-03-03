import type { AxiosRequestConfig } from 'axios';
import type { Identifier } from '@canvas/utilities';
import { SourceType } from '@canvas/types';

// Data config types
export type MappingListType = Record<string, string>; // component required field as key, rawData field as value

export type FilterConfig = {
  useFilters: boolean;
  filters: { id: string; enable: boolean }[];
};

export type AuxiliariesType = Record<string, any> & FilterConfig;

export type SourceConfigType = {
  mappingList?: MappingListType;
  autoUpdate?: number;
  controlledMode: boolean;
  auxiliaries: AuxiliariesType; //filters
  type: SourceType;
  source: SourceConfigCollection;
};

export type SourceConfigRuntime = SourceConfigType & { filters: string[] };

export type DataConfigTypeRuntime = {
  [sourceName: string]: SourceConfigRuntime;
};

export type DataConfigType<T = SourceConfigType> = {
  [sourceName: string]: T;
};

type MappingTargetType = 'number' | 'string' | 'boolean';

export type MappingStructure = Record<
  string,
  {
    description: string;
    type: MappingTargetType;
    optional?: boolean;
  }
>;

export type SingleSourcePackage = {
  description: string;
  fields: MappingStructure;
  name: string;
  /**是否默认开启受控模式 */
  controlledMode?: boolean;
};

export type DataConfigPackage = {
  [sourceName: string]: SingleSourcePackage;
};

type DatabaseConfig<T = any> = {
  sourceId: string;
  sqlContent: string;
} & T;

export type StaticDataConfig = Record<string, any>;

export type SourceConfigCollection =
  | StaticDataConfig
  | ApiConfg
  | PostgreSqlConfig
  | MySqlOracleConfig
  | SqlServerConfig;

export type MySqlOracleConfig = DatabaseConfig<{ dbName: string }>;
export type PostgreSqlConfig = DatabaseConfig<{ dbName: string; patternName: string }>;
export type SqlServerConfig = DatabaseConfig<{ dbName: string }>;

export type DatabaseCollection = MySqlOracleConfig | PostgreSqlConfig | SqlServerConfig;

// api configs
type ApiConfg = Record<string, any>;

export type IdentifierSource = {
  sourceName: string;
} & Identifier;

export type DispatchType = (params: Record<string, any>) => any;

export type SqlParams = {
  connectId: number;
  dbName: string;
  sqlContent: string;
  tableName: string;
};

export type ErrorType = {
  errCode: number;
  errMessage: string;
};

type ResponseType = {
  resultCode: 200 | 201 | 401 | 403 | 404;
  data: Record<string, any>;
  resultMessage: string;
};

export type ApiReturnType = {
  needUpdate: boolean;
  output: ResponseType | Record<string, any>[];
  error?: any;
};

export type SourceRunnerProps = {
  setData(sourceName: string, data: Record<string, any>[]): void;
  $emit(sourceName: string, payload: Record<string, any>): void;
  sourceName: string;
  sourceConfig: SourceConfigRuntime;
  /*  identifier: Identifier; */
};

export type DataResponseType = Record<
  string,
  {
    origin: Record<string, any>[];
    value: Record<string, any>[];
  }
>;

export type GetValueEntryType<T = any> = (
  source: T,
  identifier: IdentifierSource,
  config: AxiosRequestConfig,
) => Promise<ApiReturnType>;

export type {RequestApi} from "./sourceRunner/getDataPackage/getDataByApi";