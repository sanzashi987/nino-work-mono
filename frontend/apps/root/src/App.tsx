import React from 'react';
import { BrowserRouter } from 'react-router-dom';
import { createSubApp } from '@nino-work/ui-components';
import Routes from './pages';

export default createSubApp((props) => (
  <BrowserRouter basename={props.basename}>
    <Routes />
  </BrowserRouter>
));
