import { createContext } from 'react';

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
  fullpage?:boolean
};

type EnumType = {
  name: string
  value: string
};

export type UserInfoResponse = {
  user_id: number
  username: string
  menus: MenuMeta[]
  permissions: EnumType[]
  roles: EnumType[]
};

export type MicroFrontendContextType = {
  info: UserInfoResponse | null
  menus: MenuMeta[]
  matched: MenuMeta | null
  updateTitle(title: string): void
};

export const MicroFrontendContext = createContext<MicroFrontendContextType>({
  /** user info */
  info: null,
  menus: [],
  matched: null,
  updateTitle() {}
});

export const getUserInfo = () => fetch('/backend/user/v1/info').then((res) => {
  if (res.ok) {
    return res.json() as Promise<UserInfoResponse>;
  }
  return Promise.reject(new Error('Fail to fetch user info'));
});

export const getImportMap = () => fetch('/backend/user/v1/misc/importmap').then((res) => {
  if (res.ok) {
    return res.json() as Promise<MenuMeta[]>;
  }
  return Promise.reject(new Error('Fail to fetch import map'));
});
