import React from 'react';
import { Box, Card, Grid2 } from '@mui/material';
// components
import { ReactComponent as Logo } from '@nino-work/assets/logo.svg';
import AuthLogin from './Login';
import PageContainer from '@/components/PageContainer';

function Login() {
  return (
    <PageContainer title="Login" description="this is Login page">
      <Box
        sx={{
          position: 'relative',
          '&:before': {
            content: '""',
            background: 'radial-gradient(#d2f1df, #d3d7fa, #bad8f4)',
            backgroundSize: '400% 400%',
            animation: 'gradient 15s ease infinite',
            position: 'absolute',
            height: '100%',
            width: '100%',
            opacity: '0.3'
          }
        }}
      >
        <Grid2
          container
          spacing={0}
          justifyContent="center"
          sx={{
            height: '100vh'
          }}
        >
          <Box
            display="flex"
            justifyContent="center"
            alignItems="center"

          >
            <Card
              elevation={9}
              sx={{
                p: 4, zIndex: 1, width: '100%', maxWidth: '500px'
              }}
            >
              <Box display="flex" alignItems="center" justifyContent="center" mb="20px">
                <Logo width="80" height="80" />
              </Box>
              <AuthLogin />
            </Card>
          </Box>
        </Grid2>
      </Box>
    </PageContainer>
  );
}

export default Login;
