/// <reference types="node" />
export { default as theme } from './theme';
export { nanoid, uuid, createComponentId } from './uuid';
export const noop = () => {};
export const unImplemented = () => {
  throw new Error('Function not implemented.');
};

export function returnVoidObject() {
  return {};
}
export type * from './types';
export const DATE_TIME_FORMAT = 'YYYY-MM-DD HH:mm';
export * from './env';
export { stop, blockKeyEvent, prevent, preventFormEvent } from './event';
export { default as usePromise } from './usePromise';
export { queryStringify, parseQuery } from './url';
export { usePagination } from './usePagination';
export { createMemo, useCreateMemo } from './createMemo';
export { useEvent, useRefState } from './useEvent';
export { shallowEqual, shallowClone, strictEquality } from './shallow';
