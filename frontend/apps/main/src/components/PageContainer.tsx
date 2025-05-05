import { Paper } from '@mui/material';
import React, { PropsWithChildren } from 'react';
import { Helmet, HelmetProvider } from 'react-helmet-async';

type PageContainerProps = PropsWithChildren<{
  title?: string;
  description?: string;
}>;

const PageContainer: React.FC<PageContainerProps> = ({ title, description, children }) => (
  <HelmetProvider>
    <Helmet>
      <title>{title}</title>
      <meta name="description" content={description} />
    </Helmet>
    <Paper sx={{ height: '100%' }} elevation={80}>
      {children}
    </Paper>
  </HelmetProvider>
);

export default PageContainer;
