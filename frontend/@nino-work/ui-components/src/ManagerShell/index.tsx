import React, { useCallback, useEffect, useMemo, useState } from 'react';
import {
  TableCell,
  Stack,
  TableRow,
  TableBody,
  TableHead,
  TableContainer,
  Box,
  Pagination,
  Paper,
  Table,
} from '@mui/material';
import { noop, PaginationResponse, PageSize } from '@nino-work/shared';
import { Model } from './defineModel';
import loading from '../Loading';

export const useDeps = () => {
  const [key, setKey] = useState(0);

  const refresh = useCallback(() => setKey(k => k + 1), []);

  return [useMemo(() => [key], [key]), refresh] as const;
};

type ManagerShellProps<Res, T = any> = {
  schema: Model<T>[];
  requester: (parms: PageSize) => Promise<PaginationResponse<Res>>;
  ActionNode?: React.ReactNode;
  deps?: any[];
};

const ManagerShell = <Res, T>({
  requester,
  schema,
  deps,
  ActionNode,
}: ManagerShellProps<Res, T>) => {
  const [pagination, setPagination] = useState<PageSize>({ page: 1, size: 10 });
  const [data, setData] = useState<PaginationResponse<Res> | null>(null);

  useEffect(() => {
    setData(null);
    requester(pagination).then(setData).catch(noop);
  }, [...deps, pagination]);

  const tableHeader = useMemo(
    () => (
      <TableHead>
        <TableRow>
          {schema.map(e => (
            <TableCell key={e.field} {...(e.headerCellProps ?? {})}>
              {e.label}
            </TableCell>
          ))}
        </TableRow>
      </TableHead>
    ),
    [schema]
  );

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

    const inner = data.data.map((row, i) => (
      <TableRow key={i as any} sx={{ '&:last-child td, &:last-child th': { border: 0 } }}>
        {schema.map(e => {
          const val =
            typeof e.dataCellProps?.render === 'function'
              ? e.dataCellProps.render(row, i)
              : (row as any)[e.field as any];
          return (
            <TableCell key={e.field} {...(e.dataCellProps ?? {})}>
              {val}
            </TableCell>
          );
        })}
      </TableRow>
    ));
    return <TableBody>{inner}</TableBody>;
  }, [data, schema]);

  return (
    <Stack p={3}>
      {ActionNode}
      {!data ? (
        <Box flexGrow={1} mt={2}>
          {loading}
        </Box>
      ) : (
        <>
          <TableContainer component={Paper} elevation={10} sx={{ my: 3 }}>
            <Table size="small">
              {tableHeader}
              {content}
            </Table>
          </TableContainer>
          <Pagination
            sx={{ mt: 2, '.MuiPagination-ul': { justifyContent: 'end' } }}
            shape="rounded"
            size="small"
          />
        </>
      )}
    </Stack>
  );
};

export default ManagerShell;
