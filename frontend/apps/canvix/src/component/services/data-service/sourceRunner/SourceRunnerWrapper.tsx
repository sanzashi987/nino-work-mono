/* eslint-disable react-hooks/rules-of-hooks */
import React, {
  useContext, useEffect, useImperativeHandle, useMemo, useRef
} from 'react';
import { SourceRunnerProps, SourceType } from '@/types';
import SourceRunner from './sourceRunner';
import type { GetIdentifierType } from '../../proto-service/types';

type HandlerType = 'fetchData' | 'setDataRaw' | 'invokeFilter' /* | 'dataResponse' */;

type PanelDataContextType = any;
export const PassiveDataContext = React.createContext<PanelDataContextType>(null);
PassiveDataContext.displayName = 'PassiveData';

type ImperativeHandleType = {
  [K in HandlerType]: SourceRunner[K];
} & {
  getDataResponse: () => SourceRunner['dataResponse'];
};

type SourceRunnerWrapperProps = {
  getIdentifier: GetIdentifierType;
} & SourceRunnerProps;

const SourceRunnerWrapper = React.forwardRef<ImperativeHandleType, SourceRunnerWrapperProps>(
  (props, ref) => {
    const injectedDataLast = useRef([]);
    const sourceRunner = useMemo(() => new SourceRunner({ ...props }, props.getIdentifier), []);

    useImperativeHandle(
      ref,
      () => ({
        invokeFilter: sourceRunner.invokeFilter,
        fetchData: sourceRunner.fetchData,
        setDataRaw: sourceRunner.setDataRaw,
        getDataResponse: () => sourceRunner.dataResponse
      }),
      []
    );

    useEffect(() => () => {
      sourceRunner.destroy();
    }, []);

    useEffect(() => {
      sourceRunner.updateConfig(props.sourceConfig);
    }, [props.sourceConfig]);

    if (props.sourceConfig.type === SourceType.Passive) {
      const injectedData = useContext(PassiveDataContext);
      if (injectedData !== injectedDataLast.current) {
        injectedDataLast.current = injectedData;
      }
      sourceRunner.setDataRaw(injectedDataLast.current);
    }

    return null;
  }
);

SourceRunnerWrapper.displayName = 'SourceRunnerWrapper';
export default SourceRunnerWrapper;
