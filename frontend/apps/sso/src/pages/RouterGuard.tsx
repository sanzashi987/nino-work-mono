import React, { useEffect, useMemo } from 'react';
import { Navigate, Outlet, useLocation } from 'react-router-dom';
import Cookie from 'js-cookie';
import { testToken } from '@/api';

type AuthGuardProps = {};

const AuthGuard: React.FC<AuthGuardProps> = (props) => {
  const location = useLocation();

  const authed = useMemo(() => {
    const hasCookie = Cookie.get('login_token');
    return !!hasCookie;
  }, [location]);

  // useEffect(() => {
  //   const timer = setTimeout(testToken, 60 * 1000);
  //   return () => {
  //     clearTimeout(timer);
  //   };
  // }, []);

  return authed ? <Outlet /> : <Navigate to="/login" />;
};

export default AuthGuard;
