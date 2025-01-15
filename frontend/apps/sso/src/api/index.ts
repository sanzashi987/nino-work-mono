import { defineApi } from './lib';

type LoginRequest = {
  username: string
  password: string
  expiry: number
};

type LoginResponse = {
  jwt_token: string
};

const prefix = '/backend/v1';
export const login = defineApi<LoginRequest, LoginResponse>({
  url: `${prefix}/login`,
  method: 'POST'
});

type EnumType = {
  name: string
  code: string
};

export enum MenuType {
  Menu = 1,
  Catelog = 2,
  Button = 3,
}

export type MenuMeta = {
  name: string
  code: string
  icon: string
  hyperlink: boolean
  path: string
  type: MenuType
  order: number
};

export type UserInfoResponse = {
  user_id: number
  username: string
  menus: MenuMeta[]
  permissions: EnumType[]
  roles: EnumType[]
};

export const getUserInfo = defineApi<undefined, UserInfoResponse>({ url: `${prefix}/info` });

export const testToken = defineApi<{}, void>({ url: `${prefix}/token` });

export type PagninationRequest = {
  page: number
  size: number
};

type PaginationResponse<T> = {
  data: T[]
  pageIndex: number
  pageSize: number
  pageTotal: number
  recordTotal: number
};

export enum AppStatus {
  ENABLE = 0,
  DISABLE = 1,
}

type AppModel = {
  id: number
  name: string
  code: string
  description: string
  status: AppStatus
};

export type AppListResponse = PaginationResponse<AppModel>;

export const getAppList = defineApi<PagninationRequest, AppListResponse>({ url: `${prefix}/apps/list` });

export type CreateAppRequest = Pick<AppModel, 'code' | 'name' | 'description'>;

export type CreateAppResponse = AppModel;

export const createApp = defineApi<CreateAppRequest, CreateAppResponse>({
  url: `${prefix}/apps/create`,
  method: 'POST'
});
