import React from "react";
import { GEN_UID } from "@canvas/utilities";
import activate, { collections } from "./activate";
import { ConfirmDialog, NormalDialog } from "./templates";
import type { CModalConfig, CModalFuncProps, CModalFuncTypes, CModalTemplatesTypes } from "./types";
import Style from './index.module.scss';

const DialogTemplates = {
  NormalDialog,
  ConfirmDialog
};

function createModalProvider(templates: CModalTemplatesTypes) {
  const { NormalDialog, ConfirmDialog } = templates;

  const show = (options: CModalConfig) => {
    const { id = `cmodal_${GEN_UID()}`, dialogProps = {} } = options;
    const dialogConfig = { fullWidth: true, ...dialogProps };
    const props = { ...options, dialogProps: dialogConfig, id };
    return activate(NormalDialog, props);
  };

  const confirm = (type: CModalFuncTypes) => (options: CModalFuncProps) => {
    const { id = `cmodal_confirm`, dialogProps = {}, ...other } = options;
    const className = `${Style['canvas-modal-confirm']} ${other?.className || ''}`;
    const props = { ...other, dialogProps, id, className, type };
    return activate(ConfirmDialog, props);
  };

  const closeAll = () => {
    Object.values(collections)
      .filter(modal => modal.visible)
      .forEach(modal => { modal.onClose(); });
  };

  const close = (id?: string) => {
    if (!id) { return closeAll(); }
    const modal = collections[id];
    modal?.onClose();
  };

  return { show, close, confirm: confirm('confirm'), warning: confirm('warning') };
}
// success: confirm('success'),
// error: confirm('error'),
// info: confirm('info'),
// warning: confirm('warning')

const ModalBase = createModalProvider(DialogTemplates);

export * from './templates';
export * from './types';
export { createModalProvider };
export default ModalBase;
