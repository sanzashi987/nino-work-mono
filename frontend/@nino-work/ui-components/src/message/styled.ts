import { emphasize, styled, Theme } from '@mui/material/styles';
import { SnackbarProvider } from 'notistack';

const StyledSnackbarProvider = styled(SnackbarProvider)(({ theme }: { theme: Theme }) => {
  const { mode } = theme.palette;
  const backgroundColor = emphasize(
    theme.palette.background.default,
    mode === 'light' ? 0.02 : 0.2
  );
  return {
    '&.variant-default': {
      backgroundColor,
      color: theme.palette.getContrastText(backgroundColor),
    },
  };
});

export default StyledSnackbarProvider;
