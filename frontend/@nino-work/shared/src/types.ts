export type PageSize = {
  page: number;
  size: number;
};

export type PaginationResponse<T> = {
  data: T[];
  page: number;
  total: number;
};

export type ModelMeta = {
  name: string;
  code: string;
  description: string;
};

export type SubAppInjectProps = {
  basename?: string;
};

export type Enum<T = string> = {
  value: T;
  name: string;
};

export type FilterNever<T> = {
  [K in keyof T as T[K] extends never ? never : K]: T[K];
};

export type NonNullableField<T, K extends keyof T> = {
  [P in keyof T]: P extends K ? NonNullable<T[P]> : T[P];
};
