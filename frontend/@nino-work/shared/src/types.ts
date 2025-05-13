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

export type FilterObjectByType<T extends object, TargetType> = {
  [B in {
    [K in keyof T]-?: T[K] extends TargetType ? K : never;
  }[keyof T]]: T[B];
};

export type KeyFromFilterObjectByType<T extends object, TargetType> = {
  [K in keyof T]-?: T[K] extends TargetType ? K : never;
}[keyof T];

export type Writeable<T> = { -readonly [Key in keyof T]: T[Key] };

export type DeepWriteable<T> = { -readonly [P in keyof T]: DeepWriteable<T[P]> };

export type NonOptional<T> = {
  [P in keyof T]: T[P];
};
