import {
  List, ListItem, ListItemButton, ListItemText,
  Paper
} from '@mui/material';
import React, { useContext } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';
import { UserContext } from '../RouterGuard';

const SideBar: React.FC = () => {
  const { menus } = useContext(UserContext);

  const localtion = useLocation();
  const navigate = useNavigate();

  return (
    <Paper elevation={9} sx={{ overflow: 'auto', height: '100%', width: 'fit-content' }}>
      <List>
        {
          menus.map((menu) => (
            <ListItem key={menu.code} disablePadding>
              <ListItemButton
                sx={{ py: 0.5 }}
                selected={localtion.pathname === menu.path}
                onClick={() => {
                  if (localtion.pathname !== menu.path) {
                    navigate(menu.path);
                  }
                }}
              >
                {/* <ListItemIcon>
                {index % 2 === 0 ? <InboxIcon /> : <MailIcon />}
                </ListItemIcon> */}
                <ListItemText primary={menu.name} />
              </ListItemButton>
            </ListItem>
          ))
        }
      </List>
    </Paper>

  );
};

export default SideBar;
