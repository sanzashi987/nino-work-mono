import { List, ListItem, ListItemButton, ListItemText, Paper } from '@mui/material';
import { MicroFrontendContext } from '@nino-work/mf';
import React, { useContext } from 'react';
import { useLocation, useNavigate } from 'react-router-dom';

type SideBarProps = {
  style?:React.CSSProperties
};

const SideBar: React.FC<SideBarProps> = ({ style }) => {
  const { menus } = useContext(MicroFrontendContext);

  const localtion = useLocation();
  const navigate = useNavigate();

  return (
    <Paper elevation={9} style={style} sx={{ overflow: 'auto', height: '100%', width: 'fit-content' }}>
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
