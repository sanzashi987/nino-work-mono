import React, { createContext, useMemo } from 'react';
import {
  Navigate, Outlet, useLocation
} from 'react-router-dom';
import {
  getUserInfo, MenuMeta, MenuType, UserInfoResponse
} from '@/api';
import { usePromise } from '@/utils';
import loading from '@/components/Loading';

type AuthGuardProps = {};

type UserContextType = {
  info: UserInfoResponse | null
  menus: MenuMeta[]
};

export const UserContext = createContext<UserContextType>({
  info: null,
  menus: []
});

const AuthGuard: React.FC<AuthGuardProps> = (props) => {
  const userInfo = usePromise(() => getUserInfo());
  const menus = useMemo(
    () => {
      if (!userInfo) {
        return [];
      }
      const { menus: m } = userInfo;
      return m.filter((menu) => menu.type === MenuType.Menu).sort((a, b) => a.order - b.order);
    },
    [userInfo]
  );

  const location = useLocation();

  const authed = useMemo(() => {
    if (!userInfo) {
      return true;
    }
    return menus.map((e) => e.path).some((e) => e.startsWith(location.pathname));
  }, [location.pathname, menus, userInfo]);

  const ctx = useMemo(() => {
    if (!userInfo) {
      return null;
    }
    return ({ info: userInfo, menus });
  }, [userInfo, menus]);

  if (userInfo === null || ctx === null) {
    return loading;
  }

  return (
    <UserContext.Provider value={ctx}>
      {authed ? <Outlet /> : <Navigate to="/dashboard" />}
    </UserContext.Provider>
  );
};

export default AuthGuard;
