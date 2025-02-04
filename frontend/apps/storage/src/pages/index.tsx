import React, { lazy } from 'react';
import { Navigate, RouteObject } from 'react-router-dom';

const routes :RouteObject[] = [
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
    Component: () => <Navigate to="oss" />
  }
];

export default routes;
