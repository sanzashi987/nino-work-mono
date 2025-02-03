import React from 'react';
import { RouterProvider } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as defaultTheme } from '@nino-work/shared';
import router from './pages';

const theme = createTheme(defaultTheme);

const App = () => (
  <ThemeProvider theme={theme}>
    <CssBaseline />
    <RouterProvider router={router} />
  </ThemeProvider>
);

export default App;
