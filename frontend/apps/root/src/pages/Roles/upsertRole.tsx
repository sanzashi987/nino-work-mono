import { AutoSelect, Model, openSimpleForm } from '@nino-work/ui-components';
import { createRole, CreateRoleRequest, listAdminstratedPermissions } from '@/api';

const upsertModels: Model[] = [
  { label: 'Name', field: 'name', formCellProps: { widgetProps: { variant: 'standard' } } },
  { label: 'Code', field: 'code', formCellProps: { widgetProps: { variant: 'standard' } } },
  {
    label: 'Permissions',
    field: 'permission_ids',
    formCellProps: {
      widget: AutoSelect,
      widgetProps: {
        size: 'small',
        multiple: true,
        defaultValue: [],
        requester: () => listAdminstratedPermissions().then((res) => res.map((e) => ({ label: e.name, value: e.value })))
      }
    }
  },
  { label: 'Description', field: 'description', formCellProps: { widgetProps: { multiline: true, minRows: 3 } } }
];

function openUpsertRole(refresh: VoidFunction, backfill?: CreateRoleRequest) {
  openSimpleForm({
    modalProps: { title: backfill ? 'Update Role' : 'Create Role' },
    formProps: {
      schema: upsertModels,
      async onOk(form) {
        const pass = await form.trigger();
        if (!pass) return Promise.reject();

        const val = form.getValues();

        if (backfill) {
          throw new Error('todo');
        } else {
          return createRole(val).then(refresh);
        }
      }
    },
    dataBackfill: backfill
  });
}

export default openUpsertRole;
