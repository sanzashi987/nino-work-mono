import { AppBar, Box, IconButton, Menu, MenuItem, Stack, Toolbar, useColorScheme } from '@mui/material';
import React, { useContext } from 'react';
import { ReactComponent as Logo } from '@nino-work/assets/logo.svg';
import { AccountCircle } from '@mui/icons-material';
import Cookies from 'js-cookie';
import { Outlet, useNavigate } from 'react-router-dom';
import { MicroFrontendContext } from '@nino-work/mf';
import PageContainer from '@/components/PageContainer';
import SideBar from './SideBar';

// type DashboardProps = {};

const Dashboard: React.FC = () => {
  const [anchorEl, setAnchorEl] = React.useState<null | HTMLElement>(null);
  const { mode } = useColorScheme();
  const isLight = mode !== 'dark';
  const navigate = useNavigate();
  const openMenu = (event: React.MouseEvent<HTMLElement>) => {
    setAnchorEl(event.currentTarget);
  };

  const handleClose = () => {
    setAnchorEl(null);
  };

  const handleLogout = () => {
    Cookies.remove('login_token');
    navigate('/login');
    handleClose();
  };

  const { matched } = useContext(MicroFrontendContext);
  const style = matched?.fullpage ? { display: 'none' } : undefined;

  return (
    <Stack height="100%">
      <AppBar position="relative" style={style}>
        <Toolbar style={{ minHeight: 48 }}>
          <Logo width="35" height="35" />
          <div style={{ flexGrow: 1 }} />
          <IconButton onClick={openMenu}>
            {mode}
            <AccountCircle fontSize="medium" style={{ color: isLight ? 'white' : 'grey' }} />
          </IconButton>
          <Menu
            id="menu-appbar"
            anchorEl={anchorEl}
            keepMounted
            open={!!anchorEl}
            onClose={handleClose}
          >
            {/* <MenuItem onClick={handleClose}>Profile</MenuItem> */}
            <MenuItem onClick={handleLogout}>Logout</MenuItem>
          </Menu>
        </Toolbar>
      </AppBar>
      <Stack minHeight={0} flexGrow={1} direction="row">
        <SideBar style={style} />
        <Box flexGrow={1} overflow="auto">
          <Outlet />
          <div id="nino-sub-app" />
        </Box>
      </Stack>
    </Stack>
  );
};

export default Dashboard;
