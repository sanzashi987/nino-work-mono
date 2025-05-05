import { Model, openSimpleForm } from '@nino-work/ui-components';
import { ModelMeta } from '@nino-work/shared';
import { createApp, createPermission } from '@/api';

const CreateModels: Model[] = [
  {
    label: 'Name',
    field: 'name',
    formCellProps: { widgetProps: { required: true, variant: 'standard' } },
  },
  {
    label: 'Code',
    field: 'code',
    formCellProps: { widgetProps: { required: true, variant: 'standard' } },
  },
  {
    label: 'Description',
    field: 'description',
    formCellProps: { widgetProps: { multiline: true, minRows: 3 } },
  },
];

export const openCreateApp = (onSuccess: VoidFunction) => {
  openSimpleForm({
    modalProps: { title: 'Create Application' },
    formProps: {
      schema: CreateModels,
      async onOk(form) {
        const pass = await form.trigger();
        if (!pass) return Promise.reject();
        const val = form.getValues();
        return createApp(val as any).then(() => onSuccess());
      },
    },
  });
};
export const openCreatePermission = (appId: number, onSuccess: VoidFunction) => {
  const requester = (payload: ModelMeta) =>
    createPermission({ app_id: appId, permissions: [payload] });

  openSimpleForm({
    modalProps: { title: 'Create Permission' },
    formProps: {
      schema: CreateModels,
      async onOk(form) {
        const pass = await form.trigger();
        if (!pass) return Promise.reject();
        const val = form.getValues();
        return requester(val as any).then(() => onSuccess());
      },
    },
  });
};
