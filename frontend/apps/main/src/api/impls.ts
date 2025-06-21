import { defineRequester as _d, DefineRequesterOptions } from '@nino-work/requester';
import { message } from '@nino-work/ui-components';

const showErrorNotification = async (payload: any) => {
  message.error(payload?.msg, 7000);
  return Promise.reject(payload);
};

export default function defineRequester<Req, Res, Out = Res>(
  options: DefineRequesterOptions<Res, Out>
) {
  const url = `/backend/root/v1${options.url}`;
  return _d<Req, Res, Out>({ onError: showErrorNotification, ...options, url });
}
