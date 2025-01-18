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

type AppsManagementProps = {};

const AppsManagement: React.FC<AppsManagementProps> = (props) => {
  const [open, setOpen] = React.useState(false);
  const [pagination, setPagination] = useState<PagninationRequest>({ page: 1, size: 10 });
  const { data, refetch } = usePromise(() => getAppList(pagination));
  const { pathname } = useLocation();
  const naviagte = useNavigate();

  const handleClose = useCallback(() => {
    setOpen(false);
  }, []);
  const handleSuccess = useCallback(() => {
    refetch();
    setOpen(false);
  }, []);

  const content = useMemo(() => {
    if (!data) {
      return null;
    }
    if (data.data.length === 0) {
      return (
        <TableCell colSpan={6}>
          <Stack sx={{ width: '100%' }} py={1} alignItems="center" justifyContent="center">
            No Data
          </Stack>
        </TableCell>
      );
    }

    const inner = data.data.map((row) => (
      <TableRow
        key={row.name}
        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
      >
        <TableCell component="th" scope="row">
          {row.id}
        </TableCell>
        <TableCell>{row.name}</TableCell>
        <TableCell>{row.code}</TableCell>
        <TableCell>{row.description}</TableCell>
        <TableCell>{row.status}</TableCell>
        <TableCell align="center">
          <IconButton onClick={() => {
            naviagte(`${pathname}/permission/${row.id}`);
          }}
          >
            <Settings />
          </IconButton>
          <IconButton>
            <Delete />
          </IconButton>
        </TableCell>
      </TableRow>
    ));
    return <TableBody>{inner}</TableBody>;
  }, [data]);

  return (
    <Stack>
      <Button color="info" variant="contained" sx={{ width: 'fit-content' }} onClick={() => setOpen(true)}>
        + Create Application
      </Button>
      {!data
        ? (
          <Box flexGrow={1} mt={2}>
            {loading}
          </Box>
        )
        : (
          <>
            <TableContainer component={Paper} elevation={10} sx={{ mt: 2 }}>
              <Table>
                <TableHead>
                  <TableRow>
                    <TableCell>Id</TableCell>
                    <TableCell>Application Name</TableCell>
                    <TableCell align="left">Code</TableCell>
                    <TableCell align="left">Description</TableCell>
                    <TableCell align="left">Status</TableCell>
                    <TableCell align="center">Operation</TableCell>
                  </TableRow>
                </TableHead>
                {content}
              </Table>
            </TableContainer>
            <Pagination sx={{ mt: 2, '.MuiPagination-ul': { justifyContent: 'end' } }} shape="rounded" size="small" />
          </>
        )}
      <Dialog open={open}>
        <CreateAppDialog onSuccess={handleSuccess} close={handleClose} />
      </Dialog>
    </Stack>
  );
};

export default AppsManagement;
