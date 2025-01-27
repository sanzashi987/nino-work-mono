import React, { useMemo, useState } from 'react';
import { ManagerShell, useDeps } from '@nino-work/ui-components';
import { useLocation, useNavigate } from 'react-router-dom';
import { Button, IconButton } from '@mui/material';
import { Settings, Delete } from '@mui/icons-material';
import { listBucket } from '@/api';

type StorageListProps = {};

const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' }
];

const BucketList: React.FC<StorageListProps> = (props) => {
  const [open, setOpen] = useState(false);

  const { pathname } = useLocation();
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
              const p = pathname.endsWith('/') ? pathname.slice(0, -1) : pathname;
              navigate(`${p}/detail/${row.id}`);
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
    }], [pathname, navigate]);

  return (
    <ManagerShell
      schema={schema}
      deps={deps}
      requester={listBucket}
      ActionNode={(
        <Button color="info" variant="contained" sx={{ width: 'fit-content' }} onClick={() => setOpen(true)}>
          + Create Application
        </Button>
      )}
    />
  );
};

export default BucketList;
