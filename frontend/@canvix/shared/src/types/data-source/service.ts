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
