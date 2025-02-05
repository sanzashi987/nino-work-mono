import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as defaultTheme } from '@nino-work/shared';
import Routes from './pages';

const theme = createTheme(defaultTheme);

type AppProps = {
  basename?:string
};

// const router = createBrowserRouter(routes);

const App: React.FC<AppProps> = ({ basename }) => {
  console.log('basename', basename);

  // return basename;
  // const router = useMemo(() => createBrowserRouter(routes, { basename }), []);
  return (
    <ThemeProvider theme={theme}>
      <CssBaseline />
      <BrowserRouter basename={basename}>
        <Routes />
      </BrowserRouter>
    </ThemeProvider>
  );
};

export default App;
