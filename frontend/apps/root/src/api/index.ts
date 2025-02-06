import { ModelMeta, PaginationResponse, PagninationRequest } from '@nino-work/shared';
import defineApi from './impls';

const prefix = '/backend/v1';

export enum AppStatus {
  ENABLE = 0,
  DISABLE = 1,
}

type AppModel = {
  id: number
  name: string
  code: string
  description: string
  status: AppStatus
};

export type AppListResponse = PaginationResponse<AppModel>;

export const getAppList = defineApi<PagninationRequest, AppListResponse>({ url: `${prefix}/apps/list`, method: 'POST' });

export type CreateAppRequest = Pick<AppModel, 'code' | 'name' | 'description'>;

export type CreateAppResponse = AppModel;

export const createApp = defineApi<CreateAppRequest, CreateAppResponse>({
  url: `${prefix}/apps/create`,
  method: 'POST'
});

export type ListPermissionsRequest = {
  app_id: number
};
export type ListPermissionsResponse = {
  permissions: {
    id: number
    name: string
    code: string
  }[]
  super_admin_id: number
  admin_id: number
  is_admin: boolean,
  is_super: boolean
  app_name: string
};

export const listPermissions = defineApi<ListPermissionsRequest, ListPermissionsResponse>(
  { url: `${prefix}/apps/list-permission` }
);

export type CreatePermissionRequest = {
  app_id: number | string
  permissions: ModelMeta[]
};

export const createPermission = defineApi<CreatePermissionRequest, void>({
  url: `${prefix}/permission/create`,
  method: 'POST'
});
