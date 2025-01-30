import React, { JSX, useEffect, useMemo, useState } from 'react';
import LoadingButton from '@mui/lab/LoadingButton';
import {
  Box,
  Typography,
  FormControlLabel,
  Stack,
  Checkbox,
  Input
} from '@mui/material';
import IconButton from '@mui/material/IconButton';
import InputAdornment from '@mui/material/InputAdornment';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';
import { useForm } from 'react-hook-form';
import { useNavigate } from 'react-router-dom';
import Cookie from 'js-cookie';
import { noop } from '@nino-work/shared';
import { FormLabel } from '@nino-work/ui-components';
import { login } from '@/api';

interface LoginProps {
  title?: string;
  subtitle?: JSX.Element | JSX.Element[];
}

const AuthLogin: React.FC<LoginProps> = ({ title, subtitle }) => {
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const { register, handleSubmit } = useForm();

  const onSubmit = ({ username, password, remember }: any) => {
    const payload = { username, password, expiry: remember ? 30 : 1 };
    setLoading(true);
    login(payload).then(({ jwt_token }) => {
      Cookie.set('login_token', jwt_token);
      navigate('/dashboard');
    }).catch(noop).finally(() => {
      setLoading(false);
    });
  };

  useEffect(() => {
    const hasToken = Cookie.get('login_token');
    if (hasToken) {
      navigate('/dashboard');
    }
  }, []);

  const indornent = useMemo(
    () => (
      <InputAdornment position="end">
        <IconButton
          aria-label={
            showPassword ? 'hide the password' : 'display the password'
          }
          onClick={() => setShowPassword((last) => !last)}
        >
          {showPassword ? <VisibilityOff /> : <Visibility />}
        </IconButton>
      </InputAdornment>
    ),
    [showPassword]
  );

  return (
    <>
      {title ? (
        <Typography fontWeight="700" variant="h2" mb={1}>
          {title}
        </Typography>
      ) : null}
      <form onSubmit={handleSubmit(onSubmit)}>
        <Stack>
          <Box>
            <FormLabel title="Username" field="username" />
            <Input id="username" fullWidth {...register('username', { required: true })} />
          </Box>
          <Box mt="25px">
            <FormLabel title="Password" field="password" />
            <Input
              id="password"
              {...register('password', { required: true })}
              fullWidth
              type={showPassword ? 'text' : 'password'}
              endAdornment={indornent}
            />
          </Box>
          <Stack
            justifyContent="space-between"
            direction="row"
            alignItems="center"
            my={2}
          >
            <FormControlLabel
              control={(
                <Checkbox
                  id="remember"
                  {...register('remember')}
                  defaultChecked
                />
              )}
              label="Remember me in 30 days"
            />
            {/* <Typography
            component="a"
            href="/"
            fontWeight="500"
          >
            Forgot Password ?
          </Typography> */}
          </Stack>
        </Stack>
        <Box>
          <LoadingButton
            loading={loading}
            variant="contained"
            size="large"
            fullWidth
            type="submit"
          >
            Sign In
          </LoadingButton>
        </Box>
      </form>

      {subtitle}
    </>
  );
};

export default AuthLogin;
