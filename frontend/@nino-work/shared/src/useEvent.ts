import { useCallback, useLayoutEffect, useRef, useState } from 'react';

export const useEvent = <T extends Array<any>, S>(handler: (...args: T) => S) => {
  const handlerRef = useRef<any>(handler);

  useLayoutEffect(() => {
    handlerRef.current = handler;
  });

  return useCallback((...args: T) => {
    const fn = handlerRef.current;
    return fn(...args) as S;
  }, []);
};

export function useRefState<T>(initValue: T) {
  const stateRef = useRef(initValue);
  const [state, setState] = useState(initValue);
  stateRef.current = state;
  return [state, setState, stateRef] as const;
}
