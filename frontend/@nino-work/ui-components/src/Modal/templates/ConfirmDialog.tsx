import React from 'react';
import { DialogContent } from '@mui/material';
import { CheckCircleOutline, ErrorOutline, InfoOutlined, WarningAmber } from '@mui/icons-material';
import BasicDialog from './BasicDialog';
import type { ActionClickCallback, ConfirmActionsProps, ConfirmModalProps } from '../types';
import { contentClass } from '../utils';

const modalFuncIcons = {
  warning: WarningAmber,
  info: InfoOutlined,
  error: ErrorOutline,
  success: CheckCircleOutline,
  confirm: ErrorOutline
};

const loadingHandler = (
  fn?: (...args: any[]) => any | Promise<any>,
  setLoading?: (value: boolean) => void,
  close?: (...args: any) => void
) => (...args: any[]) => {
  if (!fn) { close?.(); return; }
  const returnValue = fn(...args);
  if (!returnValue || typeof returnValue.then !== 'function') return close?.();
  setLoading?.(true);
  return returnValue.then((...result: any[]) => {
    close?.(...result);
  }).finally(() => { setLoading?.(false); });
};

class ConfirmDialog extends BasicDialog<ConfirmModalProps> {
  handleOkClick: ActionClickCallback = (e, { setLoading }) => {
    const { onClose, onOk } = this.props;
    loadingHandler(onOk, setLoading, onClose)(e);
  };

  renderFooter = (props: ConfirmActionsProps) => this.renderOkCancelActions({ ...props, onOk: this.handleOkClick });

  renderContent = () => {
    const { type, children } = this.props;
    const Icon = modalFuncIcons[type || 'info'];
    return (
      <>
        <DialogContent className={contentClass}>
          <div className="confirm-content">
            <Icon className={`confirm-icon --${type}`} />
            <div className="confirm-text-wp">{children}</div>
          </div>
        </DialogContent>
        {this.renderFooter(this.props)}
      </>
    );
  };
}

export default ConfirmDialog;
