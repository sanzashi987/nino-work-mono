import React from 'react';
import { ThemeOptions, ThemeProvider, createTheme } from '@mui/material/styles';
import { RouterProvider } from 'react-router-dom';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as themeConfig } from '@nino-work/shared';
import './index.scss';

import router from './pages';

const theme = createTheme(themeConfig as ThemeOptions);

const App = () => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <RouterProvider router={router} />
  </ThemeProvider>
);

export default App;
