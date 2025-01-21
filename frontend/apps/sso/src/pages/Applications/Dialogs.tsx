import {
  Box, Button, DialogContent, DialogTitle, Divider, extendTheme, Input, Stack,
  TextField
} from '@mui/material';
import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { LoadingButton } from '@mui/lab';
import { createApp, CreateAppRequest } from '@/api';
import FormLabel from '@/components/FormLabel';

type UseForm = typeof useForm;

type CreateDialogProps<Payload extends object> = {
  title: string
  close: VoidFunction
  onSuccess: VoidFunction
  register: ReturnType<typeof useForm<Payload>>['register']
  handleSubmit: ReturnType<typeof useForm<Payload>>['handleSubmit']
};

type CreateDialogStates = {
  loading: boolean
};

abstract class BasicCreateDialog<T extends CreateDialogProps<any>>
  extends React.Component<T, CreateDialogStates> {
  constructor(props: T) {
    super(props);
    this.state = { loading: false };
  }

  onSubmit = (payload: T extends CreateDialogProps<infer Payload> ? Payload : never) => {
    this.setState({ loading: true });
    const { onSuccess } = this.props;
    this.requester()(payload).then(onSuccess).finally(() => {
      this.setState({ loading: false });
    });
  };

  abstract requester<
    F extends (
      p: T extends CreateDialogProps<infer Payload> ? Payload : never
    ) => any
  >(): F;

  render(): React.ReactNode {
    const { loading } = this.state;
    const { handleSubmit, register, close, title } = this.props;
    return (
      <Box minWidth={600}>
        <DialogTitle>{title}</DialogTitle>
        <DialogContent dividers>
          <form onSubmit={handleSubmit(this.onSubmit)}>
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
  }
}

export const CreateAppDialog: React.FC<CreateDialogProps> = ({ close, onSuccess }) => {
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

export const CreatePermissionDialog: React.FC<CreateDialogProps> = ({ close, onSuccess }) => {
  const [loading, setLoading] = useState(false);
};
