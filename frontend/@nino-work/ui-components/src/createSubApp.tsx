import React, { useMemo } from 'react';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as defaultTheme, SubAppInjectProps } from '@nino-work/shared';

const theme = createTheme(defaultTheme);

const createSubApp = (renderRouter: (props:SubAppInjectProps) => React.ReactNode) => {
  const App: React.FC<SubAppInjectProps> = (props) => {
    const children = useMemo(() => {
      const { basename = '/' } = props;
      const withDefaults = { ...props, basename };
      return React.createElement(renderRouter, withDefaults);
    }, [props]);
    return (
      <ThemeProvider theme={theme}>
        <CssBaseline />
        {children}
      </ThemeProvider>
    );
  };
  return App;
};

export default createSubApp;
