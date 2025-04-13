import { Button, ButtonProps, IconButton } from '@mui/material';
import React, { createContext, useCallback, useContext, useMemo, useState } from 'react';

// type ButtonProps = React.ComponentProps<typeof Button>['onClick'];

export type RequestButtonProps = ButtonProps & {
  mode?: 'button' | 'icon'
  onClick?: (e: React.MouseEvent) => Promise<any>
};
const LoadingContext = createContext<{ loading:boolean, setLoading:(loading:boolean)=>void } | null>(null);

export const LoadingGroup: React.FC<React.PropsWithChildren> = ({ children }) => {
  const [loading, setLoading] = useState(false);

  const ctx = useMemo(() => ({ loading, setLoading }), [loading]);

  return (
    <LoadingContext.Provider value={ctx}>
      {children}
    </LoadingContext.Provider>
  );
};

const RequestButton: React.FC<RequestButtonProps> = ({ onClick, mode, children, ...rest }) => {
  const [loading, setLoading] = useState(false);
  const inCtx = useContext(LoadingContext);
  const handleClick = useCallback((e:React.MouseEvent) => {
    (inCtx?.setLoading ?? setLoading)(true);

    Promise.resolve().then(() => onClick?.(e)).finally(() => {
      (inCtx?.setLoading ?? setLoading)(false);
    });
  }, [inCtx?.setLoading, onClick]);

  const Com = mode === 'icon' ? IconButton : Button;

  return React.createElement(
    Com,
    {
      ...rest,
      loading: inCtx?.loading ?? loading,
      onClick: handleClick
    },
    children
  );
};

export default RequestButton;
