import type { AxiosError } from 'axios';

export const getErrorInfo = (error: AxiosError & { resultMessage?: string }) => {
  return {
    name: error?.name,
    code: error?.code,
    message: error?.resultMessage || error?.message,
    data: error?.response?.data,
  };
};
