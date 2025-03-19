import React from 'react';
import { createRoot } from 'react-dom/client';
import { createTheme, styled, ThemeOptions, ThemeProvider } from '@mui/material';
import { theme, nanoid } from '@nino-work/shared';
import Message, { typeList, MessageInstance } from './Message';
import type { BasicConfigWithType, MessageType, MessageConfig, MessageContent } from './type';
import { formatConfig } from './utils';
import StyledSnackbarProvider from './styled';

const StyledMessageContainer = styled('div')({
  '.frnc': {
    display: 'flex',
    alignItems: 'center'
  },
  '.SnackbarContainer-bottom.SnackbarContainer-right': {
    bottom: 24,
    right: 8
  },
  '.SnackbarItem-contentRoot.SnackbarContent-root': {
    padding: 0,
    '.SnackbarItem-message': {
      padding: 0,
      width: '100%'
    }
  }
});

class NinoMessage {
  private maxSnack = 5;

  messageRef = React.createRef<MessageInstance>();

  constructor() {
    this.initRoot();
    this.initMethod();
  }

  theme = createTheme(theme as ThemeOptions);

  initRoot = () => {
    const dom = document.createElement('div');
    document.body.appendChild(dom);
    dom.id = 'message-container';
    const root = createRoot(dom);
    const Content = (
      <StyledMessageContainer>
        <ThemeProvider theme={this.theme}>
          <StyledSnackbarProvider
            preventDuplicate
            maxSnack={this.maxSnack}
            anchorOrigin={{ horizontal: 'right', vertical: 'bottom' }}
            hideIconVariant
          >
            <Message ref={this.messageRef} />
          </StyledSnackbarProvider>
        </ThemeProvider>
      </StyledMessageContainer>
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
    const key = config.key ?? `message_${nanoid()}`;
    const newConfig: BasicConfigWithType = {
      ...config,
      key
    };
    this.handleMessage(newConfig);
    return this.messageRef?.current?.showMessage?.(newConfig);
  };
}

const message = new NinoMessage() as unknown as MessageType;

export { MessageContent } from './Message';
export type { SnackbarKey } from 'notistack';
export default message;
