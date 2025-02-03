import React, { lazy } from 'react';
import { createBrowserRouter, Navigate } from 'react-router-dom';

const BrowserRouter = createBrowserRouter([
  {
    path: 'oss',
    children: [
      {
        index: true,
        Component: lazy(() => import('./list'))
      },
      {
        path: 'detail/:id',
        Component: lazy(() => import('./detail'))
      }
    ]
  },
  {
    path: '*',
    Component: () => <Navigate to="/oss" />
  }
]);

export default BrowserRouter;
