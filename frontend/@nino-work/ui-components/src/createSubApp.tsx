import React, { useMemo } from 'react';
import { ThemeProvider, createTheme } from '@mui/material/styles';
import CssBaseline from '@mui/material/CssBaseline';
import { theme as defaultTheme, SubAppInjectProps } from '@nino-work/shared';
import { NinoAppProvider } from '@nino-work/mf';

const theme = createTheme(defaultTheme);

const createSubApp = (
  renderRouter: (props: SubAppInjectProps) => React.ReactNode,
  overrideTheme = theme
) => {
  const App: React.FC<SubAppInjectProps> = props => {
    const children = useMemo(() => {
      const { basename } = props;
      const withDefaults = { ...props, basename };
      return React.createElement(renderRouter, withDefaults);
    }, [props]);
    return (
      <NinoAppProvider value={props.ninoAppCtx}>
        <ThemeProvider theme={overrideTheme}>
          <CssBaseline />
          {children}
        </ThemeProvider>
      </NinoAppProvider>
    );
  };
  return App;
};

export default createSubApp;
