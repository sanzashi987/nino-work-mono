export enum SourceType {
  Static = 'Static',
  API = 'API',
  // MySQL = 'MySQL',
  // Oracle = 'Oracle',
  // SQLServer = 'SQLServer',
  // PostgreSQL = 'PostgreSQL',
  File = 'File',
  Passive = 'Passive',
}

export type SourceKey = keyof typeof SourceType;

export enum SourceName {
  Static = '静态数据',
  // API = 'API接口',
  // MySQL = 'MySQL',
  // Oracle = 'Oracle',
  // SQLServer = 'SQLServer',
  // PostgreSQL = 'PostgreSQL',
  File = '文件上传',
  Passive = '被动数据',
}
