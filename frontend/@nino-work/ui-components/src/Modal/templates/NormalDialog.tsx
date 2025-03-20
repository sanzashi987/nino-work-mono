import React from 'react';
import { DialogContent } from '@mui/material';
import BasicDialog from './BasicDialog';
import { ConfirmActionsProps } from '../types';
import { contentClass } from '../utils';

class NormalDialog extends BasicDialog {
  renderFooter = (props: ConfirmActionsProps) => this.renderOkCancelActions(props);

  renderContent = () => {
    const { onClose, children } = this.props;
    const props = { onClose, renderFooter: this.renderFooter };
    return (
      <DialogContent className={contentClass}>
        {children ? React.cloneElement(children as React.ReactElement, props) : null}
      </DialogContent>
    );
  };
}

export default NormalDialog;
