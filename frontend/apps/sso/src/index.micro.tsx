import React from 'react';
import ReactDOM from 'react-dom/client';
import { registerApplication, start } from 'single-spa';
import type { MenuMeta } from './api';
import App from './App';

const menuConfig: MenuMeta[] = [
  {
    name: 'oss.nino.work',
    code: 'oss.nino.work',
    path: '/home/oss',
    icon: '',
    type: 1,
    order: 0
  }
];

menuConfig.forEach((menu) => {
  registerApplication(
    menu.code,
    () => import(/* webpackIgnore: true */ menu.code),
    (location) => location.pathname.startsWith(menu.path),
    { domElementGetter: () => document.getElementById('nino-sub-app'), basename: 'home' }
  );
});

start();

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(
  <App />
);
