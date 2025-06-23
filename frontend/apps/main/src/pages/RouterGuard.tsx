import React, { useEffect, useMemo, useState } from 'react';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import { loading } from '@nino-work/ui-components';
import { usePromise } from '@nino-work/shared';
import { getUserInfo, MenuType, NinoAppProvider } from '@nino-work/mf';
import PageContainer from '@/components/PageContainer';
import mfStore from './mf-store';

const AuthGuard: React.FC = () => {
  const { data: userInfo } = usePromise(getUserInfo);
  const [title, setTitle] = useState('Dashboard');

  const menus = useMemo(() => {
    if (!userInfo) {
      return [];
    }
    const { menus: m } = userInfo;
    return m.filter(menu => menu.type === MenuType.Menu).sort((a, b) => a.order - b.order);
  }, [userInfo]);

  const location = useLocation();

  const matched = useMemo(() => {
    if (!userInfo) {
      return null;
    }
    if (location.pathname.startsWith('/home')) {
      return null;
    }

    return (
      menus.find(e => {
        const { path } = e;
        return location.pathname.startsWith(path) || path.startsWith(location.pathname);
      }) ?? null
    );
    // return menus.map((e) => e.path).some((e) => location.pathname.startsWith(e) || e.startsWith(location.pathname));
  }, [location.pathname, menus, userInfo]);

  const ctx = useMemo(() => {
    if (!userInfo) {
      return null;
    }
    return { info: userInfo, menus, matched, updateTitle: setTitle };
  }, [userInfo, menus, matched]);

  useMemo(() => {
    if (matched && ctx) {
      mfStore[matched.code] = ctx;
    }
  }, [matched, ctx]);

  useEffect(() => {
    if (matched?.name) {
      setTitle(matched.name);
    }
  }, [matched?.name]);

  if (userInfo === null || ctx === null) {
    return loading;
  }

  return (
    <NinoAppProvider value={ctx}>
      <PageContainer title={title}>{matched ? <Outlet /> : <Navigate to="/home" />}</PageContainer>
    </NinoAppProvider>
  );
};

export default AuthGuard;
