import { useCallback, useEffect, useRef, useState } from 'react';

type Options = {
  deps?: any[]
};
const empty: any[] = [];

const usePromise = <T>(cb: () => Promise<T>, opts?: Options) => {
  const deps = opts?.deps ?? empty;
  const [val, setVal] = useState<T | null>(null);
  const [key, setKey] = useState(0);
  const cbRef = useRef<() => Promise<T>>(null);
  cbRef.current = cb;

  const refetch = useCallback(() => {
    setKey((k) => k + 1);
  }, []);

  useEffect(() => {
    cbRef.current?.().then(setVal);
  }, [key, ...deps]);

  return { data: val, refetch };
};

export default usePromise;
