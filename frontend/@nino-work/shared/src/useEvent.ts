import { useCallback, useLayoutEffect, useRef } from 'react';

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
