/* eslint-disable react/destructuring-assignment */
/* eslint-disable react/no-unused-class-component-methods */
import React, { FC } from 'react';
import { CircularProgress, ClickAwayListener } from '@mui/material';
import { withSnackbar, ProviderContext } from 'notistack';
import classNames from 'classnames';
import { CheckCircleOutline, ErrorOutline, WarningAmber, InfoOutlined } from '@mui/icons-material';
import styles from './index.module.scss';
import type { BasicConfigWithType } from '../type';

const {
  'canvas-message-content': contentClass,
  'canvas-message-wrapper': snackbarClass
} = styles;

const messageIcon: Record<MessageKey, React.ElementType> = {
  success: CheckCircleOutline,
  error: ErrorOutline,
  info: InfoOutlined,
  warning: WarningAmber,
  loading: CircularProgress
};

type MessageContentProps = {
  content: React.ReactNode;
  type: MessageKey;
  className?: string;
  onClickAway?: BasicConfigWithType['onClickAway'];
};
const MessageContent: FC<MessageContentProps> = ({ content, type, onClickAway, className }) => {
  const Icon = messageIcon[type];
  const Content = (
    <div className={classNames(className, contentClass, 'frnc')}>
      <div className="frnc">
        <Icon className={`canvas-message-icon --${type}`} />
        {content}
      </div>
    </div>
  );
  return onClickAway ? (
    <ClickAwayListener onClickAway={onClickAway!}>
      {Content}
    </ClickAwayListener>
  ) : Content;
};

type SnackbarProps = ProviderContext;
class Message extends React.Component<SnackbarProps> {
  public showMessage = (config: BasicConfigWithType) => {
    const { type, content, duration = 2000, onClickAway, variant = 'default', className = '', ...reset } = config;
    const Content = <MessageContent type={type} content={content} className={contentClass} onClickAway={onClickAway} />;
    return this.props.enqueueSnackbar?.(Content, {
      ...reset,
      variant,
      className: `${snackbarClass} variant-${variant} ${className}`,
      autoHideDuration: duration
    });
  };

  destroyMessage = (configKey?: string | number) => {
    this.props.closeSnackbar?.(configKey);
  };

  render() {
    return null;
  }
}

const MessageCtx = withSnackbar(Message);
const typeList = ['success', 'error', 'info', 'loading', 'warning'] as const;
type MessageKey = typeof typeList[number];
type MessageInstance = InstanceType<typeof Message>;

export { MessageContent, typeList };
export type { MessageKey, MessageInstance };
export default MessageCtx;
