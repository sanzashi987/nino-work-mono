import { Enum, ModelMeta, PaginationResponse, PageSize } from '@nino-work/shared';
import defineApi from './impls';

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

export const getAppList = defineApi<PageSize, AppListResponse>({
  url: 'apps/list',
  method: 'POST'
});

export type CreateAppRequest = Pick<AppModel, 'code' | 'name' | 'description'>;

export type CreateAppResponse = AppModel;

export const createApp = defineApi<CreateAppRequest, CreateAppResponse>({
  url: 'apps/create',
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
  { url: 'apps/permission/list' }
);

export type CreatePermissionRequest = {
  app_id: number | string
  permissions: ModelMeta[]
};

export const createPermission = defineApi<CreatePermissionRequest, void>({
  url: 'apps/permission/create',
  method: 'POST'
});

type RoleInfo = {
  id: number
  name: string,
  code: string
};

export const listRoles = defineApi<PageSize, PaginationResponse<RoleInfo>>({
  url: 'roles/list',
  method: 'POST'
});

export type CreateRoleRequest = {
  name: string
  code: string
  description?: string
  permission_ids?: number[]
};

export const createRole = defineApi<CreateRoleRequest, void>({
  url: 'roles/create',
  method: 'POST'
});

type UpdateRoleRequest = {
  id: number
  name?: string
  description?: string
  permission_ids?: number[]
};

export const updateRole = defineApi<UpdateRoleRequest, void>({
  url: 'roles/update',
  method: 'POST'
});

export const listAdminstratedPermissions = defineApi<void, Enum[]>({
  url: 'apps/permission/admined-permission',
  method: 'POST'
});

export type UserBio = {
  id: number
  username: string
};

export const listUsers = defineApi<PageSize, PaginationResponse<UserBio>>({
  url: 'users/list',
  method: 'POST'
});

export const getUserRoles = defineApi<{ id: number }, Enum<number>[]>({ url: 'users/user-roles' });

export type BindRoleRequest = {
  user_id: number
  role_ids:number[]
};
export const bindRoles = defineApi<BindRoleRequest, void>({
  url: 'users/bind-roles',
  method: 'POST'
});

export const listRolesAll = defineApi<void, Enum<number>[]>({
  url: 'roles/list-all',
  method: 'POST'
});
