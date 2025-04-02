import { uuid } from '@nino-work/shared';
import activate, { collections } from './activate';
import { ConfirmDialog, NormalDialog } from './templates';
import type { CModalConfig, CModalFuncProps, CModalFuncTypes, CModalTemplatesTypes } from './types';
import Style from './index.module.scss';

function createModalProvider(templates: CModalTemplatesTypes) {
  const show = (options: CModalConfig) => {
    const { id = `cmodal_${uuid()}`, dialogProps = {} } = options;
    const dialogConfig = { fullWidth: true, ...dialogProps };
    const props = { ...options, dialogProps: dialogConfig, id };
    return activate(templates.NormalDialog, props);
  };

  const confirm = (type: CModalFuncTypes) => (options: CModalFuncProps) => {
    const { id = 'cmodal_confirm', dialogProps = {}, ...other } = options;
    const className = `${Style['canvix-modal-confirm']} ${other?.className || ''}`;
    const props = {
      ...other, dialogProps, id, className, type
    };
    return activate(templates.ConfirmDialog, props);
  };

  const closeAll = () => {
    Object.values(collections)
      .filter((modal) => modal.visible)
      .forEach((modal) => { modal.onClose(); });
  };

  const close = (id?: string) => {
    if (!id) {
      closeAll();
      return;
    }
    const modal = collections[id];
    modal?.onClose();
  };

  return { show, close, confirm: confirm('confirm'), warning: confirm('warning') };
}
// success: confirm('success'),
// error: confirm('error'),
// info: confirm('info'),
// warning: confirm('warning')

const ModalBase = createModalProvider({
  NormalDialog,
  ConfirmDialog
});

export * from './templates';
export * from './types';
export { createModalProvider };
export default ModalBase;
