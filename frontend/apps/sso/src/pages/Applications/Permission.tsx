import React, { useCallback, useMemo, useState } from 'react';
import { useParams } from 'react-router-dom';
import { IconButton, Stack, Typography } from '@mui/material';
import { Delete, Settings } from '@mui/icons-material';
import { listPermissions, ListPermissionsResponse } from '@/api';
import ManagerShell, { useDeps } from '@/components/ManagerShell';

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' }
];

const PermissionManagement: React.FC = () => {
  const { appId } = useParams();
  const [res, setRes] = useState<ListPermissionsResponse | null>(null);
  const [deps, refresh] = useDeps();
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

  const handleDeletePermission = useCallback((row: ListPermissionsResponse['permissions'][number]) => {
    refresh();
  }, []);

  const schema = useMemo(() => [
    ...staticSchema,
    {
      label: 'Operation',
      field: 'id',
      dataCellProps: {
        render(row: any) {
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
        <Settings />
        <Typography variant="h5" gutterBottom m={0} ml={1}>
          {res?.app_name}
        </Typography>
      </Stack>
      <ManagerShell
        deps={deps}
        schema={schema}
        requester={requester}
      />
    </>
  );
};

export default PermissionManagement;
