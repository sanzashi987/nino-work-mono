import { lazy } from 'react';
import { createBrowserRouter } from 'react-router-dom';

const BrowserRouter = createBrowserRouter([
  {
    path: 'bucket',
    children: [
      {
        path: 'list',
        index: true,
        Component: lazy(() => import('./list'))
      },
      {
        path: 'detail/:id',
        Component: lazy(() => import('./detail'))
      }
    ]
  }
]);

export default BrowserRouter;
