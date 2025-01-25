/* eslint-disable react/destructuring-assignment */
/* eslint-disable react/no-unused-class-component-methods */
import React, { CSSProperties, FC } from 'react';
import { CircularProgress, ClickAwayListener } from '@mui/material';
import { styled } from '@mui/material/styles';
import { enqueueSnackbar, closeSnackbar } from 'notistack';
import { CheckCircleOutline, ErrorOutline, WarningAmber, InfoOutlined } from '@mui/icons-material';
import type { BasicConfigWithType } from './type';

const messageIcon: Record<MessageKey, React.ElementType> = {
  success: CheckCircleOutline,
  error: ErrorOutline,
  info: InfoOutlined,
  warning: WarningAmber,
  loading: CircularProgress
};

const messageStyle: Record<MessageKey, CSSProperties> = {
  loading: { width: 16, height: 16 },
  error: { color: '#d32f2f' },
  success: { color: '#2e7d32' },
  info: { color: '#0288d1' },
  warning: { color: '#2e7d32' }
};

type MessageContentProps = {
  content: React.ReactNode;
  type: MessageKey;
  className?: string;
  onClickAway?: BasicConfigWithType['onClickAway'];
};

const MessageContentDiv = styled('div')({
  fontSize: 12,
  padding: '6px 16px',
  width: '100%',
  '& div': {
    minHeight: 24,
    '& .message-icon': { marginRight: 8 }
  }
});

const MessageContent: FC<MessageContentProps> = ({ content, type, onClickAway, className }) => {
  const Icon = messageIcon[type];
  const style = messageStyle[type];

  const Content = (
    <MessageContentDiv className={`frnc ${className ?? ''}`}>
      <div className="frnc">
        <Icon className="message-icon" style={style} />
        {content}
      </div>
    </MessageContentDiv>
  );
  return onClickAway ? (
    <ClickAwayListener onClickAway={onClickAway!}>
      {Content}
    </ClickAwayListener>
  ) : Content;
};

class Message extends React.Component {
  // eslint-disable-next-line class-methods-use-this
  public showMessage = (config: BasicConfigWithType) => {
    const { type, content, duration = 2000, onClickAway, variant = 'default', className = '', ...reset } = config;
    const Content = <MessageContent type={type} content={content} onClickAway={onClickAway} />;

    return enqueueSnackbar?.(Content, {
      ...reset,
      variant,
      className: `variant-${variant} ${className}`,
      autoHideDuration: duration
    });
  };

  // eslint-disable-next-line class-methods-use-this
  destroyMessage = (configKey?: string | number) => {
    closeSnackbar?.(configKey);
  };

  render() {
    return null;
  }
}

const typeList = ['success', 'error', 'info', 'loading', 'warning'] as const;
type MessageKey = typeof typeList[number];
type MessageInstance = InstanceType<typeof Message>;

export { MessageContent, typeList };
export type { MessageKey, MessageInstance };
export default Message;
