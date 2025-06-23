import React from 'react';
import ReactDOM from 'react-dom/client';
import { registerApplication, start } from 'single-spa';
import { getImportMap } from '@nino-work/mf';
import App from './App';
import mfStore from './pages/mf-store';

const importMapPromise = getImportMap();

// const menuConfig: MenuMeta[] = [
//   {
//     name: 'oss.nino.work',
//     code: 'oss.nino.work',
//     path: '/home/oss',
//     icon: '',
//     type: 1,
//     order: 0
//   }
// ];

importMapPromise.then(menuConfig => {
  menuConfig.forEach(menu => {
    const basename = menu.path.startsWith('/') ? menu.path.slice(1) : menu.path;
    registerApplication(
      menu.code,
      () => System.import(/* webpackIgnore: true */ menu.code),
      location => location.pathname.startsWith(menu.path),
      () => {
        const ninoAppCtx = mfStore[menu.code];
        return {
          domElement: () => document.getElementById('nino-sub-app'),
          basename,
          ninoAppCtx,
        };
      }
    );
  });

  start();
});

const root = ReactDOM.createRoot(document.getElementById('root') as HTMLElement);
root.render(<App importMapPromise={importMapPromise} />);
