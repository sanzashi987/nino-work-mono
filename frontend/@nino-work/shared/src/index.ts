export { default as theme } from './theme';
export { nanoid, uuid, createComponentId } from './uuid';
export const noop = () => { };
export function returnVoidObject() {
  return {};
}
export type * from './types';
export const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm';
