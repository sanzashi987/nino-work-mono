import React, { lazy } from 'react';
import { createBrowserRouter, Navigate } from 'react-router-dom';
import AuthGuard from './RouterGuard';

const dashboard = lazy(() => import('./Dashboard'));

const BrowserRouter = createBrowserRouter([
  {
    path: 'login',
    Component: lazy(() => import('./Login')),
  },
  {
    path: 'home',
    Component: AuthGuard,
    children: [
      { index: true, Component: dashboard },
      { path: '*', Component: dashboard },
    ],
  },
  {
    path: '*',
    element: <Navigate to="/login" />,
  },
]);

export default BrowserRouter;
