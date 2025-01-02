import { lazy } from 'react';
import { createBrowserRouter } from 'react-router-dom';
import AuthGuard from './RouterGuard';

const BrowserRouter = createBrowserRouter([
  {
    path: 'login',
    Component: lazy(() => import('./Login'))
  },
  {
    Component: AuthGuard,
    children: [
      {
        path: 'dashboard',
        Component: lazy(() => import('./Dashboard')),
        children: [
          {
            path: 'manage/user'

          },
          {
            path: 'manage/role'
          },
          {
            path: 'manage/app',
            Component: lazy(() => import('./Applications'))
          },
          {
            path: 'manage/permission'
          }
        ]
      }
    ]
  }
]);

export default BrowserRouter;
