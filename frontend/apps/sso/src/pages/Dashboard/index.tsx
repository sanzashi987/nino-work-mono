import { Box, Stack } from '@mui/material';
import React from 'react';
import PageContainer from '@/components/PageContainer';
import SideBar from './SideBar';

type DashboardProps = {};

const Dashboard: React.FC<DashboardProps> = (props) => (
  <PageContainer title="Dashboard" description="user dashboard">
    <Stack height="100%" direction="row">
      <SideBar />
    </Stack>
  </PageContainer>
);

export default Dashboard;
