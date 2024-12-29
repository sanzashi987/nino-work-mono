import { Box } from '@mui/material';
import React from 'react';
import PageContainer from '@/components/PageContainer';
import { usePromise } from '@/utils';
import { getUserInfo } from '@/api';

type DashboardProps = {};

const Dashboard: React.FC<DashboardProps> = (props) => {
  const s = usePromise(() => getUserInfo());

  return (
    <PageContainer title="Dashboard" description="user dashboard">
      <Box />
    </PageContainer>
  );
};

export default Dashboard;
