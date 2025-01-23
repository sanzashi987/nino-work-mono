import React from 'react';
import { createRoot } from 'react-dom/client';
import { ThemeProvider } from '@mui/material';
import { GEN_UID } from '@canvas/utilities';
import Message, { typeList, MessageInstance } from './Message';
import styles from './index.module.scss';
import type { BasicConfigWithType, MessageType, MessageConfig, MessageContent } from './type';
import { formatConfig } from './utils';
import StyledSnackbarProvider from './styled';
import { defaultTheme } from '../theme';

class EncMessage {
  private maxSnack = 5;

  messageRef = React.createRef<MessageInstance>();

  constructor() {
    this.initRoot();
    this.initMethod();
  }

  initRoot = () => {
    const dom = document.createElement('div');
    document.body.appendChild(dom);
    dom.id = 'canvas-message-container';
    const root = createRoot(dom);
    const Content = (
      <ThemeProvider theme={defaultTheme}>
        <StyledSnackbarProvider
          preventDuplicate
          maxSnack={this.maxSnack}
          anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
          hideIconVariant
          classes={{ containerRoot: styles['canvas-message-root'] }}
        >
          <Message ref={this.messageRef} />
        </StyledSnackbarProvider>
      </ThemeProvider>
    );
    root.render(Content);
  };

  initMethod = () => {
    typeList.forEach((type) => {
      (this as any)[type] = (content: MessageContent | MessageConfig, duration?: number) => {
        const config = formatConfig(content, type, duration);
        if (!config) return;
        this.showMessage(config);
      };
    });
  };

  destroy = (key?: string | number) => {
    this.messageRef?.current?.destroyMessage?.(key);
  };

  /** For possible overriding */
  // eslint-disable-next-line @typescript-eslint/no-unused-vars, class-methods-use-this
  handleMessage = (message: BasicConfigWithType) => { };

  showMessage = (config: BasicConfigWithType) => {
    const key = config.key ?? `message_${GEN_UID()}`;
    const newConfig: BasicConfigWithType = {
      ...config,
      key
    };
    this.handleMessage(newConfig);
    return this.messageRef?.current?.showMessage?.(newConfig);
  };
}

const message = new EncMessage() as unknown as MessageType;

export { MessageContent } from './Message';
export type { SnackbarKey } from 'notistack';
export default message;
