export type PagninationRequest = {
  page: number
  size: number
};

export type PaginationResponse<T> = {
  data: T[]
  page_index: number
  page_size: number
  page_total: number
  record_total: number
};
