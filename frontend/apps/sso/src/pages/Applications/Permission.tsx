import React, { useCallback, useState } from 'react';
import { useParams } from 'react-router-dom';
import { listPermissions, ListPermissionsResponse } from '@/api';
import ManagerShell from '@/components/ManagerShell';
import { Typography } from '@mui/material';

const schema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' },
  { label: 'Operation', field: 'id' }
];

const PermissionManagement: React.FC = () => {
  const { appId } = useParams();
  const [res, setRes] = useState<ListPermissionsResponse | null>(null);
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

  return (
    <>
      <Typography variant='h4' gutterBottom>
        {res?.app_name}
      </Typography>
      <ManagerShell
        schema={schema}
        requester={requester}
      />
    </>
  );
};

export default PermissionManagement;
