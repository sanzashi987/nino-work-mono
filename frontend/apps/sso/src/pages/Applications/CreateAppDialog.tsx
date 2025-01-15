import {
  Box, Button, DialogContent, DialogTitle, Divider, Input, Stack,
  TextField
} from '@mui/material';
import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { LoadingButton } from '@mui/lab';
import { createApp, CreateAppRequest } from '@/api';
import FormLabel from '@/components/FormLabel';

type CreateAppDialogProps = {
  close: VoidFunction
  onSuccess: VoidFunction
};

const CreateAppDialog: React.FC<CreateAppDialogProps> = ({ close, onSuccess }) => {
  const [loading, setLoading] = useState(false);

  const { register, handleSubmit } = useForm<CreateAppRequest>();
  const onSubmit = (payload: CreateAppRequest) => {
    setLoading(true);
    createApp(payload).then(onSuccess).finally(() => {
      setLoading(false);
    });
  };

  return (
    <Box minWidth={600}>
      <DialogTitle>Create Application</DialogTitle>
      <DialogContent dividers>
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
      </DialogContent>

    </Box>
  );
};

export default CreateAppDialog;
