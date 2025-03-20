import React from 'react';
import { Dialog, DialogTitle, IconButton } from '@mui/material';
import { Close } from '@mui/icons-material';
import { blockKeyEvent } from '@canvas/utilities';
import type { CModalProps, ConfirmActionsProps } from '../types';
import { renderOkCancelActions, titleClass } from '../utils';
import Style from '../index.module.scss';

abstract class BasicDialog<P extends CModalProps = CModalProps> extends React.PureComponent<P> {
  renderCloseIcon = () => {
    const { onClose, hideClose } = this.props;
    if (hideClose) return null;
    return (
      <IconButton className="close-btn" onClick={() => onClose('closeIconClick')}>
        <Close />
      </IconButton>
    );
  };

  renderOkCancelActions = renderOkCancelActions(this.props.onClose);

  renderTitle() {
    const { title } = this.props;
    return title ? <DialogTitle className={titleClass}>{title}</DialogTitle> : null;
  }

  abstract renderFooter(props: ConfirmActionsProps): React.ReactNode;
  abstract renderContent(): React.ReactNode;

  render() {
    const { afterClose, onClose, visible, className, dialogProps = {}, keepMounted } = this.props;
    return (
      <Dialog
        tabIndex={blockKeyEvent.tabIndex}
        {...dialogProps}
        keepMounted={keepMounted}
        className={`${Style['canvas-modal']} ${className || ''}`}
        open={visible}
        onClose={(event, reason) => onClose(reason)}
        TransitionProps={{ onExited: afterClose }}
      >
        {this.renderCloseIcon()}
        {this.renderTitle()}
        {this.renderContent()}
      </Dialog>
    );
  }
}

export default BasicDialog;
