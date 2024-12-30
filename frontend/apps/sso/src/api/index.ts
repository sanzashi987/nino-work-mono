import { defineApi } from './lib';

type LoginRequest = {
  username: string
  password: string
  expiry: number
};

type LoginResponse = {
  jwt_token: string
};

const prefix = '/backend/v1';
export const login = defineApi<LoginRequest, LoginResponse>({
  url: `${prefix}/login`,
  method: 'POST'
});

type UserInfoResponse = {};

export const getUserInfo = defineApi<{}, UserInfoResponse>({
  url: `${prefix}/info`
});

export const testToken = defineApi<{}, void>({
  url: `${prefix}/token`
});
