import {
  Box, Button, DialogContent, DialogTitle, Stack,
  TextField
} from '@mui/material';
import React, { useState } from 'react';
import { useForm } from 'react-hook-form';
import { LoadingButton } from '@mui/lab';
import { createApp, CreateAppRequest, CreateMeta, createPermission } from '@/api';
import FormLabel from '@/components/FormLabel';

type CreateDialogBasicProps = {
  close: VoidFunction
  onSuccess: VoidFunction
};

type CreateDialogProps<Payload extends object> = {

  register: ReturnType<typeof useForm<Payload>>['register']
  handleSubmit: ReturnType<typeof useForm<Payload>>['handleSubmit']
} & CreateDialogBasicProps;

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
    this.requester(payload).then(onSuccess).finally(() => {
      this.setState({ loading: false });
    });
  };

  abstract title: string;

  abstract requester: (
    p: T extends CreateDialogProps<infer Payload> ? Payload : never
  ) => any;

  render(): React.ReactNode {
    const { loading } = this.state;
    const { handleSubmit, register, close } = this.props;
    return (
      <Box minWidth={600}>
        <DialogTitle>{this.title}</DialogTitle>
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

class CreateAppDialogInner extends BasicCreateDialog<CreateDialogProps<CreateAppRequest>> {
  requester = createApp;

  title = 'Create Application';
}

export const CreateAppDialog = (props: CreateDialogBasicProps) => {
  const { register, handleSubmit } = useForm<CreateAppRequest>();

  return <CreateAppDialogInner register={register} handleSubmit={handleSubmit} {...props} />;
};

type CreatePermissionProps = CreateDialogBasicProps & {
  appId: number | string
};

class CreatePermissionDialogInner extends BasicCreateDialog<CreateDialogProps<CreateMeta> & CreatePermissionProps> {
  requester = (payload: CreateMeta) => createPermission({ app_id: this.props.appId, permissions: [payload] });

  title = 'Create Permission';
}

export const CreatePermissionDialog = (props: CreatePermissionProps) => {
  const { register, handleSubmit } = useForm<CreateAppRequest>();

  return <CreatePermissionDialogInner register={register} handleSubmit={handleSubmit} {...props} />;
};
