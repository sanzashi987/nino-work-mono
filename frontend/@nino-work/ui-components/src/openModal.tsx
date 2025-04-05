import React, { useCallback, useContext, useMemo } from 'react';
import ReactDOM from 'react-dom/client';
import Dialog, { DialogProps } from '@mui/material/Dialog';
import DialogTitle from '@mui/material/DialogTitle';
import DialogContent, { DialogContentProps } from '@mui/material/DialogContent';
import DialogActions from '@mui/material/DialogActions';
import { ButtonProps } from '@mui/material/Button';
import { Box, Stack } from '@mui/material';
import { useForm, UseFormReturn } from 'react-hook-form';
import RequestButton, { LoadingGroup } from './RequestButton';
import FormBuilder, { type FormBuilderProps } from './FormBuilder';

type ModalProps = Omit<DialogProps, 'open' | 'content' | 'onClose'> & {
  title?: React.ReactNode;
  content: React.ReactNode;
  contentProps?:DialogContentProps
  action?: React.ReactNode;
  onOk?: (form: UseFormReturn<any, any, undefined>) => Promise<void>;
  onClose?: () => Promise<void>;
  afterClose?: () => Promise<void>;
  okButtonProps?: ButtonProps
  cancelButtonProps?: ButtonProps
};

export const OpenModalContext = React.createContext<
// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-expect-error
// eslint-disable-next-line @typescript-eslint/no-throw-literal
{ close:() => Promise<void>, form: UseFormReturn }>({ close: async () => { throw 'not inside context'; }, form: {} });

const DefaultAction: React.FC<Pick<ModalProps, 'onOk' | 'okButtonProps' | 'cancelButtonProps'>> = ({
  onOk,
  okButtonProps = {},
  cancelButtonProps = {}
}) => {
  const { close, form } = useContext(OpenModalContext);
  const onSubmit = useCallback(async () => {
    onOk?.(form as any).then(close);
  }, []);
  return (
    <Stack flexDirection="row-reverse">
      <LoadingGroup>
        {onOk
          ? (
            <RequestButton {...{ variant: 'outlined', type: 'submit', children: 'Ok', ...okButtonProps }} onClick={onSubmit} />
          ) : null}
        <Box mr={1}>
          <RequestButton {...{ variant: 'outlined', children: 'Cancel', ...cancelButtonProps }} onClick={close} />
        </Box>
      </LoadingGroup>
    </Stack>
  );
};

const Modal: React.FC<ModalProps> = ({
  title, content, onClose, onOk, afterClose, action, contentProps = {},
  okButtonProps,
  cancelButtonProps,
  ...dialogProps
}) => {
  const defaultActions = action ?? <DefaultAction onOk={onOk} okButtonProps={okButtonProps} cancelButtonProps={cancelButtonProps} />;
  const defaultForm = useForm();
  const { form, content: contentWithForm } = useMemo(() => {
    if (React.isValidElement(content) && 'form' in (content as any).props) {
      return { form: (content.props as any).form as UseFormReturn, content };
    }
    return { form: defaultForm, content: React.cloneElement(content as any, { form: defaultForm }) };
  }, [content, defaultForm]);

  const ctx = useMemo(() => ({ close: onClose, form }), []);
  return (
    <OpenModalContext.Provider value={ctx}>
      <Dialog maxWidth="sm" fullWidth TransitionProps={{ onExited: afterClose }} {...dialogProps} open>
        {title && <DialogTitle>{title}</DialogTitle>}
        <DialogContent sx={{ px: 2, pb: 0 }} {...contentProps}>
          {contentWithForm}
        </DialogContent>
        {action === false ? null
          : (
            <DialogActions>
              {defaultActions}
            </DialogActions>
          )}
      </Dialog>
    </OpenModalContext.Provider>
  );
};

type SimpleFormSubmit<FormData> = {
  onOk: (form:UseFormReturn<FormData, any, undefined>) => Promise<any>
};

const openModal = (props:Omit<ModalProps, 'onClose'>) => {
  const modalRoot = document.createElement('div');
  document.body.appendChild(modalRoot);

  const root = ReactDOM.createRoot(modalRoot);

  const handleClose = async () => {
    root.unmount();
    document.body.removeChild(modalRoot);
  };

  root.render(<Modal {...props} onClose={handleClose} />);

  return { close: handleClose };
};

type OpenSimpleFormProps<FormData> = {
  modalProps: Omit<ModalProps, 'onClose' | 'content'>,
  formProps: FormBuilderProps<FormData> & SimpleFormSubmit<FormData>;
  dataBackfill?: FormData
};
export const openSimpleForm = <FormData,>({ modalProps, formProps }: OpenSimpleFormProps<FormData>) => openModal({
  ...modalProps,
  content: <FormBuilder {...formProps} />,
  action: modalProps.action
});

export default openModal;
