import { defineApi as _d, DefineApiOptions } from '@nino-work/requester';
import { message } from '@nino-work/ui-components';

const showErrorNotification =  async (input:string, payload:any ) => {
  message.error(input, 7000);
  return Promise.reject(payload);
};

export default function defineApi<Req, Res>(options: DefineApiOptions) {
  const url = `/backend/storage/v1/${options.url}`;

  return _d<Req, Res>({ onError: showErrorNotification, ...options, url });
}
