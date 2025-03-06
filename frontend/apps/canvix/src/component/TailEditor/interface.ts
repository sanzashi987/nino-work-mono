import { createContext } from 'react';
import type { TailEditorInterface } from './types';

export type { TailEditorInterface } from './types';
// Only avoid optional type check, the value is expected to be overriden

async function defaultEndpoints() {
  return {
    deprecated: true,
    endpoints: {
      source: [],
      target: [],
      childList: [],
    },
  };
}
function noop() {
  return;
}

export const defaultMethods: TailEditorInterface = {
  // baseRoute: 'dashboard',
  toEndpoints: defaultEndpoints,
  // getDashboardId() {
  //   return '';
  // },
  switchPanel: noop,
  findComponentById() {
    return {} as any;
  },
  getRefNodeEndpoints: defaultEndpoints,
  menuPalette: {},
};

export const TailEditorContext = createContext<TailEditorInterface>(defaultMethods);
TailEditorContext.displayName = 'TailEditorContext';
