import React, { useMemo, useState } from 'react';
import { ManagerShell, useDeps } from '@nino-work/ui-components';
import { useNavigate } from 'react-router-dom';
import { Button, IconButton } from '@mui/material';
import { Settings, Delete } from '@mui/icons-material';
import { listBucket } from '@/api';
import openCreateBucket from './openCreateBucket';

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Code', field: 'code' },
  { label: 'Description', field: 'description' }
];

const BucketList: React.FC = () => {
  const [deps, refresh] = useDeps();

  const navigate = useNavigate();

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
              navigate(`../detail/${row.id}`);
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
    }], [navigate]);

  return (
    <ManagerShell
      schema={schema}
      deps={deps}
      requester={listBucket}
      ActionNode={(
        <Button
          color="info"
          variant="contained"
          sx={{ width: 'fit-content' }}
          onClick={() => {
            openCreateBucket(refresh);
          }}
        >
          + Create Bucket
        </Button>
      )}
    />
  );
};

export default BucketList;
