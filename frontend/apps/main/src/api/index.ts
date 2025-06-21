import defineRequester from './impls';

type LoginRequest = {
  username: string;
  password: string;
  expiry: number;
};

type LoginResponse = {
  jwt_token: string;
};

export const login = defineRequester<LoginRequest, LoginResponse>({
  url: 'users/login',
  method: 'POST',
});
