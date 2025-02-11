import { Box, Button, Stack, TextField } from '@mui/material';
import React, { useCallback, useContext, useState } from 'react';
import { useForm } from 'react-hook-form';
import { openModal, OpenModalContext, FormLabel } from '@nino-work/ui-components';
import { ModelMeta } from '@nino-work/shared';
import { createApp, createPermission } from '@/api';

type CreateProps = {
  onSuccess: VoidFunction
  requester:(params:ModelMeta)=>Promise<any>
};

const BasicCreate = ({ onSuccess, requester }: CreateProps) => {
  const [loading, setLoading] = useState(false);
  const { register, handleSubmit } = useForm<ModelMeta>();
  const { close } = useContext(OpenModalContext);

  const onSubmit = useCallback((payload: ModelMeta) => {
    setLoading(true);
    requester(payload).then(onSuccess).then(close).finally(() => {
      setLoading(false);
    });
  }, []);

  return (
    <Box minWidth={600} boxSizing="border-box" p={2}>
      <form onSubmit={handleSubmit(onSubmit)}>
        <Box>
          <FormLabel title="Name" field="name" />
          <TextField id="name" fullWidth {...register('name')} required variant="standard" />
        </Box>
        <Box mt={2}>
          <FormLabel title="Code" field="code" />
          <TextField id="code" fullWidth {...register('code')} required variant="standard" />
        </Box>
        <Box my={2}>
          <FormLabel title="Description" field="description" />
          <TextField id="description" fullWidth {...register('description')} multiline minRows={3} />
        </Box>
        <Stack flexDirection="row-reverse">
          <Button loading={loading} variant="contained" size="medium" type="submit">
            Create
          </Button>
          <Box mr={1}>
            <Button variant="outlined" onClick={close}>Cancel</Button>
          </Box>
        </Stack>
      </form>
    </Box>
  );
};

export const openCreateApp = (onSuccess: VoidFunction) => {
  openModal({
    title: 'Create Application',
    content: <BasicCreate onSuccess={onSuccess} requester={createApp} />,
    action: false
  });
};

export const openCreatePermission = (appId:number, onSuccess: VoidFunction) => {
  const requester = (payload: ModelMeta) => createPermission({ app_id: appId, permissions: [payload] });

  openModal({
    title: 'Create Permission',
    content: <BasicCreate onSuccess={onSuccess} requester={requester} />,
    action: false
  });
};
