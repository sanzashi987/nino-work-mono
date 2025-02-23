import { Model, openSimpleForm } from '@nino-work/ui-components';
import React from 'react';
import { CreateRoleRequest } from '@/api';

const upsertModels: Model[] = [
  { label: 'Name', field: 'name', formCellProps: { widgetProps: { variant: 'standard' } } },
  { label: 'Code', field: 'code', formCellProps: { widgetProps: { variant: 'standard' } } },
  {
    label: 'Permissions',
    field: 'permission_ids',
    formCellProps: {}
  },
  { label: 'Description', field: 'description', formCellProps: { widgetProps: { multiline: true, minRows: 3 } } }
];

function openUpsertRole(refresh: VoidFunction, backfill?: CreateRoleRequest) {
  openSimpleForm({
    modalProps: { title: backfill ? 'Update Role' : 'Create Role' },
    formProps: {
      schema: upsertModels,
      async onOk(form) {

      }
    },
    dataBackfill: backfill
  });
}
