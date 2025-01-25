import { defineApi as _d, DefineApiOptions } from '@nino-work/requester';
import message from '@nino-work/message';

const showErrorNotification = async (input?: any) => {
  message.error(input, 7000);
  return Promise.reject(input);
};

export default function defineApi<Req, Res>(options: DefineApiOptions) {
  return _d<Req, Res>({ onError: showErrorNotification, ...options });
}
