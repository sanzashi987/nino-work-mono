import React, { lazy } from 'react';
import { Navigate, RouteObject, useRoutes } from 'react-router-dom';

const routes :RouteObject[] = [
  {
    path: '/root',
    children: [
      {
        path: 'app',
        children: [
          {
            index: true,
            Component: lazy(() => import('./Applications/index'))

          },
          {
            path: 'permission/:appId',
            Component: lazy(() => import('./Applications/Permission'))
          }
        ]
      },
      {
        path: 'role',
        children: [
          {
            index: true,
            Component: lazy(() => import('./Roles/index'))
          }
        ]

      },
      {
        path: 'role',
        children: [
          {
            index: true,
            Component: lazy(() => import('./Users/index'))
          }
        ]

      }
    ]
  },
  {
    path: '*',
    Component: () => <Navigate to="/root" />
  }
];

export default () => useRoutes(routes);
