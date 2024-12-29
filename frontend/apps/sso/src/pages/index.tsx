import React from 'react';
import { createBrowserRouter } from 'react-router-dom';

const BrowserRouter = createBrowserRouter([
  {
    path: 'login',
    Component: React.lazy(() => import('./Login'))
  }
]);

export default BrowserRouter;
