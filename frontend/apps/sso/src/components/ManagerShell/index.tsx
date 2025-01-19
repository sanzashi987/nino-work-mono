import React, { useMemo, useState } from 'react';
import { Settings, Delete } from '@mui/icons-material';
import {
  TableCell, Stack, TableRow, IconButton, TableBody,
  TableHead,
  TableContainer,
  Box,
  Pagination,
  Paper,
  Table
} from '@mui/material';
import { PagninationRequest } from '@/api';
import { Model } from './defineModel';
import { usePromise } from '@/utils';
import loading from '../Loading';

type ManagerShellProps<Res, T = any> = {
  schema: Model<T>[],
  requester: (parms: PagninationRequest) => Promise<Res>
  ActionNode: React.ReactNode
};

const ManagerShell = <Res extends any[], T>({
  requester,
  schema,
  ActionNode
}: ManagerShellProps<Res, T>) => {
  const [pagination, setPagination] = useState<PagninationRequest>({ page: 1, size: 10 });
  const { data, refetch } = usePromise(() => requester(pagination));

  const tableHeader = useMemo(() => (
    <TableHead>
      <TableRow>
        {schema.map((e) => (
          <TableCell key={e.field} {...e.headerCellProps ?? {}}>
            {e.label}
          </TableCell>
        ))}
      </TableRow>
    </TableHead>
  ), [schema]);

  const content = useMemo(() => {
    if (!data) {
      return null;
    }
    if (data.length === 0) {
      return (
        <TableCell colSpan={6}>
          <Stack sx={{ width: '100%' }} py={1} alignItems="center" justifyContent="center">
            No Data
          </Stack>
        </TableCell>
      );
    }

    const inner = data.map((row, i) => (
      <TableRow
        key={row.name}
        sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
      >
        {schema.map((e, schemaIdx) => {
          const val = typeof e.dataCellProps?.render === 'function' ? e.dataCellProps.render(row, i) : row[e.field];
          return (
            <TableCell key={e.field} {...e.dataCellProps ?? {}}>
              {val}
            </TableCell>
          );
        })}
      </TableRow>
    ));
    return <TableBody>{inner}</TableBody>;
  }, [data, schema]);

  return (
    <Stack>
      {ActionNode}
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
                {tableHeader}
                {content}
              </Table>
            </TableContainer>
            <Pagination sx={{ mt: 2, '.MuiPagination-ul': { justifyContent: 'end' } }} shape="rounded" size="small" />
          </>
        )}
    </Stack>
  );
};

export default ManagerShell;
