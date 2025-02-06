import React, { lazy } from 'react';
import { createBrowserRouter, Navigate } from 'react-router-dom';
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
        path: 'home/*',
        Component: lazy(() => import('./Dashboard'))
      }
    ]
  },
  {
    path: '*',
    Component: () => <Navigate to="/login" />
  }
]);

export default BrowserRouter;
