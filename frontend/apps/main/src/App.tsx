import React, { useCallback } from 'react';
import { ThemeOptions, ThemeProvider, createTheme } from '@mui/material/styles';
import { RouterProvider } from 'react-router-dom';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as themeConfig, usePromise } from '@nino-work/shared';
import './index.scss';

import { loading } from '@nino-work/ui-components';
import { MenuMeta } from '@nino-work/mf';
import router from './pages';

const theme = createTheme(themeConfig as ThemeOptions);

type AppProps = {
  importMapPromise : Promise<MenuMeta[]>
};
const App: React.FC<AppProps> = ({ importMapPromise }) => {
  const { data } = usePromise(useCallback(() => importMapPromise, []));

  const isLoading = data === null;

  if (isLoading) {
    return loading;
  }

  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <RouterProvider router={router} />
    </ThemeProvider>
  );
};

export default App;
