import React, { useCallback, useContext, useMemo, useState } from 'react';
import ReactDOM from 'react-dom/client';

import Dialog, { DialogProps } from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent, { DialogContentProps } from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import Button from '@mui/material/Button';
import { Box, Stack } from '@mui/material';
import { useForm, UseFormReturn } from 'react-hook-form';
import RequestButton from '../RequestButton';
import FormBuilder, { type FormBuilderProps } from '../FormBuilder';

type ModalProps = Omit<DialogProps, 'open' | 'content' | 'onClose'> & {
  title?: React.ReactNode;
  content: React.ReactNode;
  contentProps?:DialogContentProps
  action?: React.ReactNode;
  onClose: () => void;
};

export const OpenModalContext = React.createContext<
// @ts-ignore
// eslint-disable-next-line @typescript-eslint/no-throw-literal
{ close: VoidFunction, form: UseFormReturn }>({ close: () => { throw 'not inside context'; }, form: {} });

const Modal: React.FC<ModalProps> = ({ title, content, onClose, action, contentProps = {}, ...dialogProps }) => {
  const defaultActions = action ?? <Button onClick={onClose}>Close</Button>;
  const form = useForm();
  const ctx = useMemo(() => ({ close: onClose, form }), []);
  return (
    <OpenModalContext.Provider value={ctx}>
      <Dialog maxWidth="sm" fullWidth {...dialogProps} open>
        {title && <DialogTitle>{title}</DialogTitle>}
        <DialogContent sx={{ p: 2 }} {...contentProps}>
          {content}
        </DialogContent>
        <DialogActions>
          {defaultActions}
        </DialogActions>
      </Dialog>
    </OpenModalContext.Provider>
  );
};

type SimpleFormSubmit<FormData> = {
  onOk: (form:UseFormReturn<FormData, any, undefined>) => Promise<any>
};

const SimpleFormAction = <FormData, >({ onOk }: SimpleFormSubmit<FormData>) => {
  const [loading, setLoading] = useState(false);
  const { close, form } = useContext(OpenModalContext);

  const onSubmit = useCallback(() => {
    setLoading(true);
    onOk(form as any).then(close).finally(() => {
      setLoading(false);
    });
  }, []);

  return (
    <Stack flexDirection="row-reverse">
      <RequestButton loading={loading} variant="contained" size="medium" type="submit" onClick={onSubmit}>
        Submit
      </RequestButton>
      <Box mr={1}>
        <RequestButton loading={loading} variant="outlined" onClick={close}>Cancel</RequestButton>
      </Box>
    </Stack>
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

type OpenSimpleFormProps<FormData> = {
  modalProps: Omit<ModalProps, 'onClose' | 'content'>,
  formProps: FormBuilderProps<FormData> & SimpleFormSubmit<FormData>;
};
export const openSimpleForm = <FormData,>({ modalProps, formProps }: OpenSimpleFormProps<FormData>) => {
  const defaultAction = <SimpleFormAction onOk={formProps.onOk} />;

  return openModal({
    ...modalProps,
    content: <FormBuilder {...formProps} />,
    action: modalProps.action ?? defaultAction
  });
};

export default openModal;
