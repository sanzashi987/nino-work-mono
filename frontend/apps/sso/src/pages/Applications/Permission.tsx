import React, { useCallback, useMemo, useState } from 'react';
import { useNavigate, useParams } from 'react-router-dom';
import {
  Button, Dialog, IconButton, Stack, Typography
} from '@mui/material';
import { ArrowBack, Delete } from '@mui/icons-material';
import { useDeps, ManagerShell } from '@nino-work/ui-components';
import { listPermissions, ListPermissionsResponse } from '@/api';
import { CreatePermissionDialog } from './Dialogs';

type PermissionMeta = ListPermissionsResponse['permissions'][number];

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' }
];

const PermissionManagement: React.FC = () => {
  const [open, setOpen] = useState(false);
  const { appId } = useParams();
  const [res, setRes] = useState<ListPermissionsResponse | null>(null);
  const [deps, refresh] = useDeps();
  const naviagte = useNavigate();

  const handleClose = useCallback(() => {
    setOpen(false);
  }, []);
  const handleSuccess = useCallback(() => {
    refresh();
    setOpen(false);
  }, []);

  const requester = useCallback(
    () => listPermissions({ app_id: Number(appId) }).then((response) => {
      setRes(response);
      const len = response.permissions.length;
      return {
        data: response.permissions,
        page_index: 1,
        page_size: len,
        page_total: len,
        record_total: len
      };
    }),
    [appId]
  );

  const handleDeletePermission = useCallback((row: PermissionMeta) => {
    refresh();
  }, []);

  const schema = useMemo(() => [
    ...staticSchema,
    {
      label: 'Operation',
      field: 'id',
      dataCellProps: {
        render(row: PermissionMeta) {
          return (
            <IconButton onClick={() => handleDeletePermission(row)}>
              <Delete />
            </IconButton>
          );
        }
      }
    }
  ], []);

  return (
    <>
      <Stack direction="row" alignItems="center">
        <IconButton onClick={() => { naviagte('..'); }}>
          <ArrowBack />
        </IconButton>
        <Typography variant="h5" gutterBottom m={0} ml={1}>
          {res?.app_name ?? '...'}
        </Typography>
      </Stack>
      <ManagerShell
        deps={deps}
        schema={schema}
        requester={requester}
        ActionNode={(
          <Button color="info" variant="contained" sx={{ width: 'fit-content' }} onClick={() => setOpen(true)}>
            + Create Permission
          </Button>
        )}
      />
      <Dialog open={open}>
        <CreatePermissionDialog onSuccess={handleSuccess} close={handleClose} appId={appId!} />
      </Dialog>
    </>
  );
};

export default PermissionManagement;
