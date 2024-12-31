import React, { createContext, useEffect, useMemo } from 'react';
import {
  Navigate, Outlet, useInRouterContext, useLocation, useNavigation
} from 'react-router-dom';
import Cookie from 'js-cookie';
import { getUserInfo, testToken, UserInfoResponse } from '@/api';
import { usePromise } from '@/utils';
import loading from '@/components/Loading';

type AuthGuardProps = {};

export const UserContext = createContext<UserInfoResponse>({
  user_id: 0,
  username: '',
  menus: [],
  permissions: [],
  roles: []
});

const AuthGuard: React.FC<AuthGuardProps> = (props) => {
  const userInfo = usePromise(() => getUserInfo());
  // const location = useLocation();
  const { state, location } = useNavigation();

  console.log(state, location);
  const authed = true;
  // useEffect(() => {
  //   const timer = setTimeout(testToken, 60 * 1000);
  //   return () => {
  //     clearTimeout(timer);
  //   };
  // }, []);
  if (userInfo === null) {
    return loading;
  }

  return (
    <UserContext.Provider value={userInfo}>
      {authed ? <Outlet /> : <Navigate to="/dashboard" />}
    </UserContext.Provider>
  );
};

export default AuthGuard;
