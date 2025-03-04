import type { Identifier } from '../../com-config';
import type { SourceType } from './source-type';

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
  auxiliaries: AuxiliariesType; // filters
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

// api configs
type ApiConfg = Record<string, any>;

export type DatabaseCollection = MySqlOracleConfig | PostgreSqlConfig | SqlServerConfig;

export type IdentifierSource = {
  sourceName: string;
} & Identifier;

export type SourceRunnerProps = {
  setData(sourceName: string, data: Record<string, any>[]): void;
  $emit(sourceName: string, payload: Record<string, any>): void;
  sourceName: string;
  sourceConfig: SourceConfigRuntime;
};

export type DataResponseType = Record<
string,
{
  origin: Record<string, any>[];
  value: Record<string, any>[];
}
>;

export type ApiReturnType = {
  needUpdate: boolean;
  output: ResponseType | Record<string, any>[];
  error?: any;
};

export type GetValueEntryType<T = any> = (
  source: T,
  identifier: IdentifierSource,
  config: any,
) => Promise<ApiReturnType>;

export type RequestApi = {
  sourceId: string;
  body?: Record<string, any>;
  headers?: Record<string, any>;
  querys?: Record<string, any>;
  path?: Record<string, any>;
};
