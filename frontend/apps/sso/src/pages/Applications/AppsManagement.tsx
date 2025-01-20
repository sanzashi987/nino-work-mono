import {
  Box,
  Button, Dialog, IconButton, Pagination, Paper, Stack, Table, TableBody, TableCell, TableContainer,
  TableHead,
  TableRow
} from '@mui/material';
import React, { useCallback, useMemo, useState } from 'react';
import { Delete, Details, Link, Settings } from '@mui/icons-material';
import { useLocation, useNavigate } from 'react-router-dom';
import { getAppList, PagninationRequest } from '@/api';
import { usePromise } from '@/utils';
import loading from '@/components/Loading';
import CreateAppDialog from './CreateAppDialog';
import ManagerShell from '@/components/ManagerShell';


const staticSchema = [
  { label: 'Id', field: 'id' },
  { label: 'Name', field: 'name' },
  { label: 'Code', field: 'code' },
  { label: 'Description', field: 'description' },
  { label: 'Status', field: 'status' },
]




const AppsManagement: React.FC = () => {
  const [open, setOpen] = useState(false);
  const { pathname } = useLocation();
  const naviagte = useNavigate();

  const handleClose = useCallback(() => {
    setOpen(false);
  }, []);
  const handleSuccess = useCallback(() => {
    // refetch();
    setOpen(false);
  }, []);

  const schema = useMemo(() => {
    return [
      ...staticSchema,
      {
        label: 'Operation', field: 'id',
        headerCellProps: { align: 'center' as const },
        dataCellProps: {
          align: 'center' as const,
          render: (row: any) => <>
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
        }
      }]
  }, [pathname, naviagte])

  const requester = useCallback((req: PagninationRequest) => {
    return getAppList(req)
  }, [])

  return (
    <>
      <ManagerShell schema={schema} requester={requester} />
      <Dialog open={open}>
        <CreateAppDialog onSuccess={handleSuccess} close={handleClose} />
      </Dialog>
    </>
  );
};

export default AppsManagement;
