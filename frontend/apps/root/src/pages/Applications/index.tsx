import { Button, IconButton } from '@mui/material';
import React, { useMemo } from 'react';
import { Delete, Settings } from '@mui/icons-material';
import { useNavigate } from 'react-router-dom';
import { useDeps, ManagerShell } from '@nino-work/ui-components';
import { getAppList } from '@/api';
import { openCreateApp } from './Dialogs';

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' },
  { label: 'Description', field: 'description' },
  { label: 'Status', field: 'status' },
];

const AppsManagement: React.FC = () => {
  const [deps, refresh] = useDeps();

  const navigate = useNavigate();

  const schema = useMemo(
    () => [
      ...staticSchema,
      {
        label: 'Operation',
        field: 'id',
        headerCellProps: { align: 'center' as const },
        dataCellProps: {
          align: 'center' as const,
          render: (row: any) => (
            <>
              <IconButton
                onClick={() => {
                  navigate(`./permission/${row.id}`);
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
    [navigate]
  );

  return (
    <ManagerShell
      deps={deps}
      schema={schema}
      requester={getAppList}
      ActionNode={
        <Button
          color="info"
          variant="contained"
          sx={{ width: 'fit-content' }}
          onClick={() => openCreateApp(refresh)}
        >
          + Create Application
        </Button>
      }
    />
  );
};

export default AppsManagement;
