import {
  Box, Button, Paper, Stack, Table, TableBody, TableCell, TableContainer,
  TableHead,
  TableRow
} from '@mui/material';
import React, { useState } from 'react';

type AppsManagementProps = {};

const AppsManagement: React.FC<AppsManagementProps> = (props) => {
  const [loading, setLoading] = useState();

  return (
    <Stack>
      <Button color="info" variant="contained" sx={{ width: 'fit-content' }}>
        + Create Application
      </Button>

      <TableContainer component={Paper} elevation={10} sx={{ mt: 2 }}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>Id</TableCell>
              <TableCell>Application Name</TableCell>
              <TableCell align="left">Code</TableCell>
              <TableCell align="left">Description</TableCell>
              <TableCell align="right">Status</TableCell>
              <TableCell align="right">Operation</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {[].map((row:any) => (
              <TableRow
                key={row.name}
                sx={{ '&:last-child td, &:last-child th': { border: 0 } }}
              >
                <TableCell component="th" scope="row">
                  {row.name}
                </TableCell>
                <TableCell align="right">{row.calories}</TableCell>
                <TableCell align="right">{row.fat}</TableCell>
                <TableCell align="right">{row.carbs}</TableCell>
                <TableCell align="right">{row.protein}</TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </Stack>
  );
};

export default AppsManagement;
