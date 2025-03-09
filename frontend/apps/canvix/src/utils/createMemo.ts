import { useMemo } from 'react';
import { strictEquality } from './shallow';

export function createMemo<T extends Array<any>, S>(
  func: (...args: T) => S,
  equalityFunc = strictEquality,
  initInput: T = [] as unknown as T
) {
  let lastInput: T = initInput;
  let lastRes: any = null;
  return function wrapper(...args: T) {
    // if (args.reduce((last, val, index) => val !== lastInput[index] || last, false)) {
    // if (args.some((val, index) => val !== lastInput[index])) {
    if (args.some((val, index) => !equalityFunc(val, lastInput[index]))) {
      lastRes = func(...args);
      lastInput = args;
    }
    return lastRes as S;
  };
}

export function useCreateMemo<T extends Array<any>, S>(
  func: (...args: T) => S,
  equalityFunc = strictEquality,
  initInput: T = [] as unknown as T
) {
  return useMemo(() => createMemo(func, equalityFunc, initInput), []);
}
