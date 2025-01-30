import { openModal, FormLabel, OpenModalContext } from '@nino-work/ui-components';
import React, { useState, useContext, useCallback } from 'react';
import { LoadingButton } from '@mui/lab';
import { Box, TextField, Stack, Button } from '@mui/material';
import { useForm } from 'react-hook-form';
import { createBucket, BucketMeta } from '@/api';

type CreateProps = {
  onSuccess: VoidFunction
  requester: (params: BucketMeta) => Promise<any>
};

const CreateBucket = ({ onSuccess, requester }: CreateProps) => {
  const [loading, setLoading] = useState(false);
  const { register, handleSubmit } = useForm<BucketMeta>();
  const { close } = useContext(OpenModalContext);

  const onSubmit = useCallback((payload: BucketMeta) => {
    setLoading(true);
    requester(payload).then(onSuccess).then(close).finally(() => {
      setLoading(false);
    });
  }, []);

  return (
    <Box minWidth={600} p={2}>
      <form onSubmit={handleSubmit(onSubmit)}>
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

const openCreateBucket = (onSuccess: VoidFunction) => {
  openModal({
    title: 'Create Bucket',
    content: <CreateBucket onSuccess={onSuccess} requester={createBucket} />,
    action: false
  });
};

export default openCreateBucket;
