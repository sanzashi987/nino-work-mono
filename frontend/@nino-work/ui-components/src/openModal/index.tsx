import React, { useMemo } from 'react';
import ReactDOM from 'react-dom/client';

import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent, { DialogContentProps } from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';

interface ModalProps {
  title?: React.ReactNode;
  content: React.ReactNode;
  contentProps?:DialogContentProps
  action?: React.ReactNode;
  onClose: () => void;
}

export const OpenModalContext = React.createContext<
// eslint-disable-next-line @typescript-eslint/no-throw-literal
{ close: VoidFunction }>({ close: () => { throw 'not inside context'; } });

const Modal: React.FC<ModalProps> = ({ title, content, onClose, action, contentProps = {} }) => {
  const defaultActions = action ?? <Button onClick={onClose}>Close</Button>;
  const ctx = useMemo(() => ({ close: onClose }), []);

  return (
    <OpenModalContext.Provider value={ctx}>
      <Dialog open onClose={onClose}>
        {title && <DialogTitle>{title}</DialogTitle>}
        <DialogContent sx={{ p: 0 }} {...contentProps}>
          {content}
        </DialogContent>
        <DialogActions>
          {defaultActions}
        </DialogActions>
      </Dialog>
    </OpenModalContext.Provider>
  );
};

const SimpleForm = () => { };

const openModal = (props:Omit<ModalProps, 'onClose'>) => {
  const modalRoot = document.createElement('div');
  document.body.appendChild(modalRoot);

  const root = ReactDOM.createRoot(modalRoot);

  const handleClose = () => {
    root.unmount();
    document.body.removeChild(modalRoot);
  };

  root.render(<Modal {...props} onClose={handleClose} />);

  return { close: handleClose };
};

export default openModal;
