import React, { useMemo } from 'react';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import { loading } from '@nino-work/ui-components';
import { usePromise } from '@nino-work/shared';
import { getUserInfo, MenuType, UserContext } from '@nino-work/mf';

const AuthGuard: React.FC = () => {
  const { data: userInfo } = usePromise(getUserInfo);
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

  const matched = useMemo(() => {
    if (!userInfo) {
      return null;
    }
    if (location.pathname === '/home') {
      return null;
    }

    return menus.find((e) => {
      const { path } = e;
      return location.pathname.startsWith(path) || path.startsWith(location.pathname);
    }) ?? null;
    // return menus.map((e) => e.path).some((e) => location.pathname.startsWith(e) || e.startsWith(location.pathname));
  }, [location.pathname, menus, userInfo]);

  const ctx = useMemo(() => {
    if (!userInfo) {
      return null;
    }
    return ({ info: userInfo, menus, matched });
  }, [userInfo, menus, matched]);

  if (userInfo === null || ctx === null) {
    return loading;
  }

  return (
    <UserContext.Provider value={ctx}>
      {matched ? <Outlet /> : <Navigate to="/home" />}
      {/* <Outlet /> */}
    </UserContext.Provider>
  );
};

export default AuthGuard;
