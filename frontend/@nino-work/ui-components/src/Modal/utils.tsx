/* eslint-disable @typescript-eslint/naming-convention */
import React from 'react';
import { DialogActions } from '@mui/material';
import type { ActionProps, ConfirmActionsProps } from './types';
import RequestButton from '../RequestButton';

export const contentClass = 'canvix-modal-content';
export const titleClass = 'canvix-modal-title';
export const footerClass = 'canvix-modal-footer';

export const renderActions = (actions: ActionProps[]): React.ReactNode => {
  if (!actions.length) return null;
  const footNodes = actions.map((action, index) => {
    const { text, ...props } = action;
    return (
      // eslint-disable-next-line react/no-array-index-key
      <RequestButton key={index} variant="outlined" {...props}>
        {text}
      </RequestButton>
    );
  });
  return <DialogActions className={footerClass}>{footNodes}</DialogActions>;
};

export const renderOkCancelActions = (onClose?: any) => (props: ConfirmActionsProps) => {
  const { okText, cancelText, onOk, onCancel, cancelButton = true, okButtonProps = {}, cancelButtonProps = {} } = props;
  const okCancel: ActionProps[] = [
    ...(cancelButton ? [{
      ...cancelButtonProps,
      text: cancelText || '取消',
      onClick: onCancel || onClose
    }] : []),
    {
      variant: 'contained',
      autoFocus: true,
      ...okButtonProps,
      text: okText || '确定',
      onClick: onOk
    }
  ];
  return renderActions(okCancel);
};
