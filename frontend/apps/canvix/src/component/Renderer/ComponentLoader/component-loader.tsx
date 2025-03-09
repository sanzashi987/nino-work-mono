import React, { useCallback, useEffect, useRef, useState } from 'react';
import { ResponsiveController } from '@canvix/component-factory';
import { composeCacheId } from '@/component/dub';
import WidthDebugWrapper from './debugComWrapper';
import { RuntimeInterface } from '../context';

function craeteComponentLoader({ loadModule, cachedComponents }: RuntimeInterface) {
  const ComponentLoader: React.ForwardRefRenderFunction<any, ResponsiveController.LoaderBasicProps> = ({ mounted, ...comProps }, ref) => {
    const { com, id } = comProps.config;
    const { name, version, user, isDebugger } = com;
    const [component, setComponent] = useState<Record<string, React.ComponentType<any> | null>>({});

    const firstMounted = useRef(false);

    const loadComponent = useCallback(() => {
      const cacheId = composeCacheId(name, version);
      const cachedComponent = cachedComponents.get(cacheId);
      if (cachedComponent) setComponent({ [version]: cachedComponent });
      else {
        const params = {
          id, name, version, user, isDebugger
        };
        loadModule(params).then((module: any) => {
          setComponent({ [version]: module });
          cachedComponents.set(cacheId, module);
        });
      }
    }, [name, id, version, isDebugger, user]);

    useEffect(() => {
      loadComponent();
    }, [loadComponent]);
    useEffect(() => {
      if (component[version] && !firstMounted.current) {
        mounted();
        firstMounted.current = true;
      }
    }, [component, mounted, version]);

    const target = component[version];
    const Com = isDebugger ? WidthDebugWrapper(target!) : target;
    return Com ? <Com {...comProps} ref={ref} /> : null;
  };

  return React.forwardRef(ComponentLoader);
}

export default craeteComponentLoader;
