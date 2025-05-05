import React from 'react';
import ReactDOMClient from 'react-dom/client';
import singleSpaReact from 'single-spa-react';

const lifecycles = singleSpaReact({
  React,
  ReactDOMClient,
  errorBoundary() {
    return <div>Error</div>;
  },
  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-expect-error
  loadRootComponent: () => import('./App').then(mod => mod.default),
});

export const { bootstrap, mount, unmount } = lifecycles;
