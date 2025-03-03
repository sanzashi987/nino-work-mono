export { default as API } from './getDataByApi';
export { getDataByCommonSQL as MySQL } from './getDataByDatabase';
export { getDataByCommonSQL as Oracle } from './getDataByDatabase';
export { getDataByCommonSQL as SQLServer } from './getDataByDatabase';
export { getDataByPGSQL as PostgreSQL } from './getDataByDatabase';
export { default as Static } from './getDataByStatic';
export { default as File } from './getDataByFile';
const dummy = async () => {
  return null;
};

export const WebSocket = dummy;
export const Passive = dummy;
