import {
  Box, Button, Stack,
  TextField
} from '@mui/material';
import React, { useCallback, useContext, useState } from 'react';
import { useForm } from 'react-hook-form';
import { LoadingButton } from '@mui/lab';
import { openModal, OpenModalContext } from '@nino-work/ui-components';
import { createApp, CreateMeta, createPermission } from '@/api';
import FormLabel from '@/components/FormLabel';

type CreateProps = {
  onSuccess: VoidFunction
  requester:(params:CreateMeta)=>Promise<any>
};

const BasicCreate = ({ onSuccess, requester }: CreateProps) => {
  const [loading, setLaoding] = useState(false);
  const { register, handleSubmit } = useForm<CreateMeta>();
  const { close } = useContext(OpenModalContext);

  const onSubmit = useCallback((payload: CreateMeta) => {
    setLaoding(true);
    requester(payload).then(onSuccess).finally(() => {
      setLaoding(false);
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
          <LoadingButton loading={loading} variant="contained" size="medium" type="submit">
            Create
          </LoadingButton>
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
  const requester = (payload: CreateMeta) => createPermission({ app_id: appId, permissions: [payload] });

  openModal({
    title: 'Create Permission',
    content: <BasicCreate onSuccess={onSuccess} requester={requester} />,
    action: false

  });
};
