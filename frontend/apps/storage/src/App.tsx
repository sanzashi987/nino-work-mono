import React, { useMemo } from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as defaultTheme } from '@nino-work/shared';
import routes from './pages';

const theme = createTheme(defaultTheme);

type AppProps = {
  basename?:string
};

const App: React.FC<AppProps> = ({ basename }) => {
  console.log('basename', basename);
  const router = useMemo(() => createBrowserRouter(routes, { basename }), []);
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <RouterProvider router={router} />
    </ThemeProvider>
  );
};

export default App;
