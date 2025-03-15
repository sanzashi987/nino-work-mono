export type CanvixResponse<T = any> = {
  resultCode: number;
  data: T;
  resultMessage: string;
};

export type CanvixResponsePromise<T = any> = Promise<CanvixResponse<T>>;

export type Pagination = {
  pageIndex: number;
  pageSize: number;
  pageTotal: number;
  recordTotal: number;
};
