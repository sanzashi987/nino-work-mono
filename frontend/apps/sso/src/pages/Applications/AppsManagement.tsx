import { Button, Dialog, IconButton } from '@mui/material';
import React, { useCallback, useMemo, useState } from 'react';
import { Delete, Settings } from '@mui/icons-material';
import { useLocation, useNavigate } from 'react-router-dom';
import { getAppList, PagninationRequest } from '@/api';
import { CreateAppDialog } from './Dialogs';
import ManagerShell, { useDeps } from '@/components/ManagerShell';

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' },
  { label: 'Description', field: 'description' },
  { label: 'Status', field: 'status' }
];

const AppsManagement: React.FC = () => {
  const [open, setOpen] = useState(false);
  const { pathname } = useLocation();

  const [deps, refresh] = useDeps();

  const naviagte = useNavigate();

  const handleClose = useCallback(() => {
    setOpen(false);
  }, []);
  const handleSuccess = useCallback(() => {
    refresh();
    setOpen(false);
  }, []);

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
            <IconButton onClick={() => {
              naviagte(`${pathname}/permission/${row.id}`);
            }}
            >
              <Settings />
            </IconButton>
            <IconButton>
              <Delete />
            </IconButton>
          </>
        )
      }
    }], [pathname, naviagte]);

  const requester = useCallback((req: PagninationRequest) => getAppList(req), []);

  return (
    <>
      <ManagerShell
        deps={deps}
        schema={schema}
        requester={requester}
        ActionNode={(
          <Button color="info" variant="contained" sx={{ width: 'fit-content' }} onClick={() => setOpen(true)}>
            + Create Application
          </Button>
        )}
      />
      <Dialog open={open}>
        <CreateAppDialog onSuccess={handleSuccess} close={handleClose} />
      </Dialog>
    </>
  );
};

export default AppsManagement;
