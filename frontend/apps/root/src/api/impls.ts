import { defineApi as _d, DefineApiOptions } from '@nino-work/requester';
import { message } from '@nino-work/ui-components';

const showErrorNotification = async (payload:any) => {
  message.error(payload?.msg, 7000);
  return Promise.reject(payload);
};

export default function defineApi<Req, Res, Out = Res>(options: DefineApiOptions<Res, Out>) {
  const url = `/backend/root/v1/${options.url}`;

  return _d<Req, Res, Out>({ onError: showErrorNotification, ...options, url });
}
