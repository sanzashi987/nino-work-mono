import { Settings, Delete } from '@mui/icons-material';
import { IconButton, Button } from '@mui/material';
import { useDeps, ManagerShell } from '@nino-work/ui-components';
import React, { useMemo } from 'react';
import { listRoles } from '@/api';
import { openCreateApp } from '../Applications/Dialogs';

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Role Name', field: 'name' },
  { label: 'Code', field: 'code' },
  { label: 'Description', field: 'description' }
];

const RoleManagement: React.FC = () => {
  const [deps, refresh] = useDeps();

  const schema = useMemo(() => [
    ...staticSchema,
    {
      label: 'Operation',
      field: 'id',
      headerCellProps: { align: 'center' as const },
      dataCellProps: {
        align: 'center' as const,
        render: (row: any) => (
          <>
            <IconButton onClick={() => {}}>
              <Settings />
            </IconButton>
            <IconButton>
              <Delete />
            </IconButton>
          </>
        )
      }
    }], []);

  return (
    <ManagerShell
      deps={deps}
      schema={schema}
      requester={listRoles}
      ActionNode={(
        <Button color="info" variant="contained" sx={{ width: 'fit-content' }} onClick={() => openCreateApp(refresh)}>
          + Create Role
        </Button>
      )}
    />
  );
};

export default RoleManagement;
