import { useEffect, useState } from 'react';

const empty:any[] = [];

const usePromise = <T, Deps extends any[]>(cb: () => Promise<T>, deps?: Deps) => {
  const [val, setVal] = useState<T | null>(null);

  useEffect(() => {
    cb().then(setVal);
  }, deps ?? empty);

  return val;
};

export default usePromise;
