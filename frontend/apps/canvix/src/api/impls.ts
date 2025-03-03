import { CanvixResponse } from '@canvix/shared';
import { defineApi as _d, DefineApiOptions, StandardResponse } from '@nino-work/requester';
import { message } from '@nino-work/ui-components';

export default function defineApi<Req, Res>(options: DefineApiOptions) {
  const url = `/backend/storage/v1/${options.url}`;

  const showErrorNotification = async (input: string, payload: StandardResponse<Res>) => {
    message.error(input, 7000);

    // Adapt standard response into canvix response
    const canvixResponse: CanvixResponse<Res> = {
      resultCode: payload.code,
      data: payload.data,
      resultMessage: payload.msg
    };

    return Promise.reject(canvixResponse);
  };

  return _d<Req, Res>({ onError: showErrorNotification, ...options, url });
}
