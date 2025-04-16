export type PageSize = {
  page: number
  size: number
};

export type PaginationResponse<T> = {
  data: T[]
  index: number
  size: number
  total: number
};

export type ModelMeta = {
  name: string
  code: string
  description: string
};

export type SubAppInjectProps = {
  basename?: string
};

export type Enum<T = string> = {
  value: T
  name: string
};
