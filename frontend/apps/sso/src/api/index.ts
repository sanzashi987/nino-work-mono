import { defineApi } from './lib';

type LoginRequest = {
  username: string
  password:string
};

type LoginResponse = {};

const prefix = '/backend/v1';
export const login = defineApi<LoginRequest, LoginResponse>({
  url: `${prefix}/login`,
  method: 'POST'
});

type UserInfoResponse = {};

export const getUserInfo = defineApi<{}, UserInfoResponse>({
  url: `${prefix}/info`
});
