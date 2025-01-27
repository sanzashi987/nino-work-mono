import React from 'react';
import ReactDOM from 'react-dom/client';

import Dialog from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';

interface ModalProps {
  title?: React.ReactNode;
  content: React.ReactNode;
  action?: React.ReactNode;
  onClose: () => void;
}

const Modal: React.FC<ModalProps> = ({ title, content, onClose, action }) => {
  const defaultActions = action ?? <Button onClick={onClose}>Close</Button>;
  return (
    <Dialog open onClose={onClose}>
      {title && <DialogTitle>{title}</DialogTitle>}
      <DialogContent dividers>
        {content}
      </DialogContent>
      <DialogActions>
        {defaultActions}
      </DialogActions>
    </Dialog>
  );
};

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
