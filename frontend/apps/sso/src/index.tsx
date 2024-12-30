import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import { RouterProvider } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import router from './pages';

const darkTheme = createTheme({
  palette: { mode: 'dark' }
});

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
  <ThemeProvider theme={darkTheme}>
    <CssBaseline />
    <RouterProvider router={router} />
  </ThemeProvider>
);
