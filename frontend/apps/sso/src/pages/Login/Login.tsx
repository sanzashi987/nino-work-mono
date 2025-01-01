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
import { useFormik } from 'formik';
import { useNavigate } from 'react-router-dom';
import Cookie from 'js-cookie';
import { login } from '@/api';

interface LoginProps {
  title?: string;
  subtitle?: JSX.Element | JSX.Element[];
  subtext?: JSX.Element | JSX.Element[];
}

const AuthLogin: React.FC<LoginProps> = ({ title, subtitle, subtext }) => {
  const [showPassword, setShowPassword] = useState(false);
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const formik = useFormik({
    initialValues: { username: '', password: '', remember: true },
    onSubmit: ({ username, password, remember }) => {
      const paylaod = { username, password, expiry: remember ? 30 : 1 };
      setLoading(true);
      login(paylaod).then(({ jwt_token }) => {
        Cookie.set('login_token', jwt_token);
        navigate('/dashboard');
      }).finally(() => {
        setLoading(false);
      });
    }
  });

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

      {subtext}
      <form onSubmit={formik.handleSubmit}>
        <Stack>
          <Box>
            <Typography
              variant="subtitle1"
              fontWeight={600}
              component="label"
              htmlFor="username"
              mb="5px"
              mr="5px"
            >
              Username
            </Typography>
            <Input
              id="username"
              name="username"
              fullWidth
              value={formik.values.username}
              onChange={formik.handleChange}
            />
          </Box>
          <Box mt="25px">
            <Typography
              variant="subtitle1"
              fontWeight={600}
              component="label"
              htmlFor="password"
              mb="5px"
              mr="5px"
            >
              Password
            </Typography>
            <Input
              id="password"
              name="password"
              fullWidth
              type={showPassword ? 'text' : 'password'}
              endAdornment={indornent}
              value={formik.values.password}
              onChange={formik.handleChange}
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
                  name="remember"
                  onChange={formik.handleChange}
                  checked={formik.values.remember}
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
