import { Box } from '@mui/material';
import React, { PropsWithChildren } from 'react';
import { Helmet, HelmetProvider } from 'react-helmet-async';

type PageContainerProps = PropsWithChildren<{
  title?: string
  description?: string
}>;

const PageContainer: React.FC<PageContainerProps> = ({ title, description, children }) => (
  <HelmetProvider>
    <Helmet>
      <title>{title}</title>
      <meta name="description" content={description} />
    </Helmet>
    <Box
      height="100%"
      sx={{
        position: 'relative',
        '&:before': {
          content: '""',
          background: 'radial-gradient(#d2f1df, #d3d7fa, #bad8f4)',
          backgroundSize: '400% 400%',
          animation: 'gradient 15s ease infinite',
          position: 'absolute',
          height: '100%',
          width: '100%',
          opacity: '0.3'
        }
      }}
    >
      {children}
    </Box>
  </HelmetProvider>
);

export default PageContainer;
