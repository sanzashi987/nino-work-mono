import React from 'react';
import { Box, Card, Grid2, useColorScheme } from '@mui/material';
import { ReactComponent as Logo } from '@nino-work/assets/logo.svg';
import AuthLogin from './Login';
import PageContainer from '@/components/PageContainer';

function Login() {
  const { mode } = useColorScheme();
  const stroke = mode === 'dark' ? '#fff' : '#000';
  return (
    <PageContainer title="Login" description="this is Login page">
      <Grid2 container spacing={0} justifyContent="center" sx={{ height: '100vh' }}>
        <Box display="flex" justifyContent="center" alignItems="center">
          <Card elevation={9} sx={{ p: 4, zIndex: 1, width: '100%', maxWidth: '500px' }}>
            <Box display="flex" alignItems="center" justifyContent="center" mb="20px">
              <Logo width="80" height="80" stroke={stroke} />
            </Box>
            <AuthLogin />
          </Card>
        </Box>
      </Grid2>
    </PageContainer>
  );
}

export default Login;
