export type PagninationRequest = {
  page: number
  size: number
};

export type PaginationResponse<T> = {
  data: T[]
  page_index: number
  page_size: number
  record_total: number
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
