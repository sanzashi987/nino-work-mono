import React, { lazy } from 'react';
import { createBrowserRouter, Navigate } from 'react-router-dom';
import AuthGuard from './RouterGuard';

const BrowserRouter = createBrowserRouter([
  {
    path: 'login',
    index: true,
    Component: lazy(() => import('./Login'))
  },
  {
    Component: AuthGuard,
    children: [
      {
        path: 'home',
        Component: lazy(() => import('./Dashboard')),
        children: [
          {
            path: 'root/app',
            Component: lazy(() => import('./Applications')),
            children: [
              {
                Component: lazy(() => import('./Applications/AppsManagement')),
                index: true
              },
              {
                path: 'permission/:appId',
                Component: lazy(() => import('./Applications/Permission'))
              }
            ]
          }

        ]
      }
    ]
  },
  {
    path: '*',
    Component: () => <Navigate to="/login" />
  }
]);

export default BrowserRouter;
