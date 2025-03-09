import React from 'react';
import ReactDOM from 'react-dom/client';

declare global{
  interface Window {
    // eslint-disable-next-line @typescript-eslint/ban-types
    dub(name: string, deps: string[] | Function, callback: Function): void;
  }
}
