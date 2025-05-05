import defineApi from './impls';

type LoginRequest = {
  username: string;
  password: string;
  expiry: number;
};

type LoginResponse = {
  jwt_token: string;
};

export const login = defineApi<LoginRequest, LoginResponse>({
  url: 'users/login',
  method: 'POST',
});
