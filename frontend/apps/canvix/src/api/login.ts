/** 根据token获取用户的详细信息，用在通过单点登录获取token之后 */
export const getUserInfo = () => authRequest.GET('common/userInfo');
