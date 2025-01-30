const theme = {
  spacing: 4,
  shape: { borderRadius: 2 },
  palette: {
    // mode: 'dark',
    // primary: {
    //   main: '#fc7b21',
    // },
    contrastThreshold: 3,
    tonalOffset: 0.2
    // text: { primary: '#bcc9d4' }
  },
  typography: {
    fontSize: 10,
    htmlFontSize: 14,
    button: { textTransform: 'none' }
  }
} as const;

export default theme;
