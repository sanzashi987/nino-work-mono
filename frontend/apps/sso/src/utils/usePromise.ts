import { useCallback, useEffect, useRef, useState } from 'react';

const usePromise = <T>(cb: () => Promise<T>) => {
  const [val, setVal] = useState<T | null>(null);
  const [key, setKey] = useState(0);
  const cbRef = useRef<() => Promise<T>>(null);
  cbRef.current = cb;

  const refetch = useCallback(() => {
    setKey((k) => k + 1);
  }, []);

  useEffect(() => {
    cbRef.current?.().then(setVal);
  }, [key]);

  return { data: val, refetch };
};

export default usePromise;
