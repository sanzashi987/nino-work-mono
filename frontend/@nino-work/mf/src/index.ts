import { createElement, createContext, FC, ReactNode } from 'react';

export enum MenuType {
  Menu = 1,
  Catelog = 2,
  Button = 3,
}

export type MenuMeta = {
  name: string;
  code: string;
  icon: string;
  // basename for the app
  path: string;
  type: MenuType;
  order: number;
  fullpage?: boolean;
};

type EnumType = {
  name: string;
  value: string;
};

export type UserInfoResponse = {
  user_id: number;
  username: string;
  menus: MenuMeta[];
  permissions: EnumType[];
  roles: EnumType[];
};

export type NinoAppContextType = {
  info: UserInfoResponse | null;
  menus: MenuMeta[];
  matched: MenuMeta | null;
  updateTitle(title: string): void;
};

export const NinoAppContext = createContext<NinoAppContextType>({
  /** user info */
  info: null,
  menus: [],
  matched: null,
  updateTitle() {},
});

// eslint-disable-next-line import/no-mutable-exports
export let updateTitle = (_: string) => null;

export const NinoAppProvider: FC<{
  value: NinoAppContextType;
  children: ReactNode;
}> = ({ value, children }) => {
  updateTitle = value.updateTitle;

  return createElement(NinoAppContext.Provider, { value }, children);
};

export const getUserInfo = () =>
  fetch('/backend/user/v1/info').then(res => {
    if (res.ok) {
      return res.json() as Promise<UserInfoResponse>;
    }
    return Promise.reject(new Error('Fail to fetch user info'));
  });

export const getImportMap = () =>
  fetch('/backend/user/v1/misc/importmap').then(res => {
    if (res.ok) {
      return res.json() as Promise<MenuMeta[]>;
    }
    return Promise.reject(new Error('Fail to fetch import map'));
  });
