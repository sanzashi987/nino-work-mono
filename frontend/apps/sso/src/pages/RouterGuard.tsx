import React, { createContext, useMemo } from 'react';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import { loading } from '@nino-work/ui-components';
import { getUserInfo, MenuMeta, MenuType, UserInfoResponse } from '@/api';
import { usePromise } from '@/utils';

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
  const { data: userInfo } = usePromise(() => getUserInfo());
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
    if (location.pathname === '/home') {
      return true;
    }

    return menus.map((e) => e.path).some((e) => location.pathname.startsWith(e) || e.startsWith(location.pathname));
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
      {authed ? <Outlet /> : <Navigate to="/home" />}
    </UserContext.Provider>
  );
};

export default AuthGuard;
