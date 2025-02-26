import defineApi from './impls';

type LoginRequest = {
  username: string
  password: string
  expiry: number
};

type LoginResponse = {
  jwt_token: string
};

const prefix = '/backend/root/v1';
export const login = defineApi<LoginRequest, LoginResponse>({
  url: `${prefix}/users/login`,
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
