import { ComponentType, PropsWithChildren } from 'react';
import type { DialogProps } from '@mui/material';
import { RequestButtonProps } from '../RequestButton';

export type BaseConfig = {
  /**
   * The unique id of dialog. Id will be auto generated if not supplied.
   */
  id: string;
  /**
   * The title of dialog.
   */
  title?: React.ReactNode;
  /**
   * The content of dialog.
   */
  content?: React.ReactNode;
  /**
   * Whether to unmount child components on onClose.
   * @default false
   */
  keepMounted?: boolean;
  /**
   * The class name of the container of the dialog.
   */
  className?: string;
  /**
   * Whether a close (x) button is visible on top right of the dialog or not.
   * @default false
   */
  hideClose?: boolean;
  /**
   * If true, clicking backdrop will not fire the onClose callback.
   * @default false
   */
  disableBackdropClick?: boolean;
  /**
   * Callback fired when the dialog requests to be closed.
   * @example function(reason: string) => void
   * @param {string} reason Can be: `"escapeKeyDown"`, `"backdropClick"` `"closeIconClick"`.
   * Result can be the value passed from `onClose` in content component.
   * @example function(result: any) => void
   */
  onClose?: (...args: any[]) => void;
  /**
   * Callback fired when dialog is closed completely
   */
  afterClose?: () => void;
};
export type BaseConfigRuntime = BaseConfig & { visible: boolean };
export type BaseModalProps = Omit<BaseConfigRuntime, 'content'> & Required<Pick<BaseConfig, 'onClose' | 'afterClose'>>;
export type ModalComponentProps<T = BaseModalProps> = PropsWithChildren<T>;
export type ModalComponent<T = BaseModalProps> = ComponentType<ModalComponentProps<T>>;

type StandardProps = Omit<DialogProps, keyof BaseConfig | 'open' | 'children'>;
export interface CModalConfig<T = StandardProps> extends Partial<BaseConfig> { dialogProps?: T }
export interface CModalProps<T = StandardProps> extends ModalComponentProps { dialogProps?: T }
export type CModalTemplatesTypes = Record<string, ModalComponent<CModalProps>>;
export type ContainerProps<T = {}> = T & {
  /**
   * Callback passed into content component to control dialog close.
   */
  onClose?: BaseConfig['onClose'];
  /**
   * Footer render function passed into content component for template actions rendering.
   */
  renderFooter?: (props: ConfirmActionsProps) => React.ReactNode;
};

export type ActionHandlers = { setLoading: (value: boolean) => void; };
export type ActionClickCallback = (
  e: React.MouseEvent,
  handlers: ActionHandlers
) => void | Promise<any>;
export interface ActionProps extends Omit<RequestButtonProps, 'onClick'> {
  text?: React.ReactNode;
  onClick?: ActionClickCallback;
}
export interface ConfirmActionsProps extends Omit<ConfirmProps, 'onOk' | 'onCancel'> {
  onOk?: ActionClickCallback;
  onCancel?: ActionClickCallback;
}

type ConfirmProps = {
  okText?: string;
  cancelText?: string;
  cancelButton?: boolean;
  okButtonProps?: ActionProps;
  cancelButtonProps?: ActionProps;
  onOk?: (e: React.MouseEvent) => void | Promise<any>;
  onCancel?: (e: React.MouseEvent) => void | Promise<any>;
};

export type CModalFuncTypes = 'info' | 'success' | 'error' | 'warning' | 'confirm';
export type CModalFuncProps = CModalConfig & ConfirmProps;
export type ConfirmModalProps = CModalProps & ConfirmProps & { type?: CModalFuncTypes };
