import React, { PropsWithChildren } from 'react';
import { Helmet, HelmetProvider } from 'react-helmet-async';

type PageContainerProps = PropsWithChildren<{
  title?: string
  description?: string
}>;

const PageContainer: React.FC<PageContainerProps> = ({ title, description, children }) => (
  <HelmetProvider>
    <div>
      <Helmet>
        <title>{title}</title>
        <meta name="description" content={description} />
      </Helmet>
      {children}
    </div>
  </HelmetProvider>
);

export default PageContainer;
