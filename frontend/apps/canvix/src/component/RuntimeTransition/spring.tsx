import { useImperativeHandle, useMemo } from 'react';
import { SpringValue, useSpring } from '@react-spring/web';

const useSpringBasic = (basicConfig: Record<string, any>, hide: boolean, transitionRef: any) => {
  const [springs, api] = useSpring<Record<string, any>>(() => {
    const display = hide ? 'none' : basicConfig.display ?? 'block';
    return { ...basicConfig, display };
  }, [basicConfig, hide]);

  useImperativeHandle(transitionRef, () => api, [api]);

  const res = useMemo(() => {
    const nextSprings = { ...springs };
    Object.keys(springs)
      .filter((key) => basicConfig[key] === undefined && key !== 'display')
      .forEach((key) => {
        nextSprings[key] = new SpringValue('unset');
      });

    return nextSprings;
  }, [basicConfig]);

  // console.log(res);
  return res;
};

export default useSpringBasic;
