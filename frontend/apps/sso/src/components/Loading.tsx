import { Box, CircularProgress } from '@mui/material';
import React from 'react';

const loading = (
  <Box display="flex" justifyContent="center" alignItems="center" height="100%">
    <CircularProgress />
  </Box>
);

export default loading;
