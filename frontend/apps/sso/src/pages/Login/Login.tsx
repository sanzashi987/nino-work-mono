import React, { JSX, useMemo } from 'react';
import {
  Box,
  Typography,
  FormGroup,
  FormControlLabel,
  Button,
  Stack,
  Checkbox,
  Input
} from '@mui/material';
import IconButton from '@mui/material/IconButton';
import InputAdornment from '@mui/material/InputAdornment';
import Visibility from '@mui/icons-material/Visibility';
import VisibilityOff from '@mui/icons-material/VisibilityOff';

interface LoginProps {
  title?: string;
  subtitle?: JSX.Element | JSX.Element[];
  subtext?: JSX.Element | JSX.Element[];
}

const AuthLogin: React.FC<LoginProps> = ({ title, subtitle, subtext }) => {
  const [showPassword, setShowPassword] = React.useState(false);

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
          <Input id="username" fullWidth />
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
          <FormGroup>
            <FormControlLabel
              control={<Checkbox defaultChecked />}
              label="Remember me in 30 Days"
            />
          </FormGroup>
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
        <Button
          variant="contained"
          size="large"
          fullWidth
          href="/"
          type="submit"
        >
          Sign In
        </Button>
      </Box>
      {subtitle}
    </>
  );
};

export default AuthLogin;
