import { Enum } from '@nino-work/shared';
import { FormLabel, openModal, OpenModalContext } from '@nino-work/ui-components';
import React, { useContext, useEffect, useState } from 'react';
import { Box, Button, MenuItem, Paper, Select, Stack } from '@mui/material';
import { bindRoles, getUserRoles, listRolesAll } from '@/api';

type UserDetailProps = {
  userId: number
  refresh:VoidFunction
};

const UserDetail: React.FC<UserDetailProps> = ({ userId, refresh }) => {
  const { close } = useContext(OpenModalContext);
  const [roles, setRoles] = useState<number[]>([]);
  const [roleOptions, setRoleOptions] = useState<Enum<number>[]>([]);

  const [loading, setLoading] = useState(true);

  const submit = async () => {
    const payload = {
      role_ids: roles,
      user_id: userId
    };
    setLoading(true);
    return bindRoles(payload).then(close).finally(() => {
      setLoading(false);
    });
  };

  useEffect(() => {
    Promise.all([
      listRolesAll().then(setRoleOptions),
      getUserRoles({ id: userId }).then((res) => {
        setRoles(res.map((e) => e.value));
      })
    ]).then(() => {
      setLoading(false);
    });
  }, []);

  return (
    <Box>
      <Stack>
        <FormLabel title="Roles" field="role_ids" />
        <Select
          variant="standard"
          disabled={loading}
          value={roles}
          multiple
          onChange={(e) => {
            setRoles(e.target.value as number[]);
          }}
        >
          {roleOptions.map((e) => <MenuItem key={e.value} value={e.value}>{e.name }</MenuItem>)}
        </Select>
      </Stack>

      <Stack direction="row">
        <Box ml="auto" mt={2}>
          <Button onClick={close} sx={{ marginRight: '8px' }}>Cancel</Button>
          <Button loading={loading} onClick={submit} variant="contained">Ok</Button>
        </Box>
      </Stack>
    </Box>
  );
};

const openUserDetail = (userId: number, refresh:VoidFunction) => {
  openModal({
    title: 'Bind Roles',
    content: <UserDetail userId={userId} refresh={refresh} />,
    action: false
  });
};

export default openUserDetail;
