/* eslint-disable no-plusplus */
/* eslint-disable no-bitwise */
const urlAlphabet = 'ModuleSymbhasOwnPr0123456789ABCDEFGHNRVfgctiUvzKqYTJkLxpZXIjQW';

export const nanoid = (size = 6) => {
  let id = '';
  let i = size;
  while (i--) {
    id += urlAlphabet[(Math.random() * 62) | 0];
  }
  return id;
};

export const createComponentId = (type) => `${type}_${nanoid()}`;
