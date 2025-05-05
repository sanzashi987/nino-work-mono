import { Settings, Delete } from '@mui/icons-material';
import { IconButton, Button } from '@mui/material';
import { useDeps, ManagerShell } from '@nino-work/ui-components';
import React, { useMemo } from 'react';
import { listUsers, UserBio } from '@/api';
import openUserDetail from './detail';

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Username', field: 'username' },
  { label: 'Email', field: 'email' },
];

const UserManagement: React.FC = () => {
  const [deps, refresh] = useDeps();

  const schema = useMemo(
    () => [
      ...staticSchema,
      {
        label: 'Operation',
        field: 'id',
        headerCellProps: { align: 'center' as const },
        dataCellProps: {
          align: 'center' as const,
          render: (row: UserBio) => (
            <>
              <IconButton
                onClick={() => {
                  openUserDetail(row.id, refresh);
                }}
              >
                <Settings />
              </IconButton>
              <IconButton>
                <Delete />
              </IconButton>
            </>
          ),
        },
      },
    ],
    []
  );

  return (
    <ManagerShell
      deps={deps}
      schema={schema}
      requester={listUsers}
      // ActionNode={(
      //   <Button color="info" variant="contained" sx={{ width: 'fit-content' }} onClick={() => openUpsertRole(refresh)}>
      //     + Create User
      //   </Button>
      // )}
    />
  );
};

export default UserManagement;
