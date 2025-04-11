import { useCallback, useEffect, useState } from 'react';

type Options = {
  deps?: any[]
};
const empty: any[] = [];

const usePromise = <T>(promise:(() => Promise<T>) | Promise<T>, opts?: Options) => {
  const deps = opts?.deps ?? empty;
  const [val, setVal] = useState<T | null>(null);
  const [key, setKey] = useState(0);

  const refetch = useCallback(() => {
    setKey((k) => k + 1);
  }, []);

  useEffect(() => {
    const p = promise instanceof Promise ? promise : promise();
    p.then(setVal);
  }, [key, ...deps]);

  return { data: val, refetch };
};

export default usePromise;
