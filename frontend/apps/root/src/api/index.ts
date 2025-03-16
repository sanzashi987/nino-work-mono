import { Enum, ModelMeta, PaginationResponse, PagninationRequest } from '@nino-work/shared';
import defineApi from './impls';

const prefix = '/backend/root/v1';

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

export const getAppList = defineApi<PagninationRequest, AppListResponse>({
  url: `${prefix}/apps/list`,
  method: 'POST'
});

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
  { url: `${prefix}/apps/permission/list` }
);

export type CreatePermissionRequest = {
  app_id: number | string
  permissions: ModelMeta[]
};

export const createPermission = defineApi<CreatePermissionRequest, void>({
  url: `${prefix}/apps/permission/create`,
  method: 'POST'
});

type RoleInfo = {
  id: number
  name: string,
  code: string
};

export const listRoles = defineApi<PagninationRequest, PaginationResponse<RoleInfo>>({
  url: `${prefix}/roles/list`,
  method: 'POST'
});

export type CreateRoleRequest = {
  name: string
  code: string
  description?: string
  permission_ids?: number[]
};

export const createRole = defineApi<CreateRoleRequest, void>({
  url: `${prefix}/roles/create`,
  method: 'POST'
});

type UpdateRoleRequest = {
  id: number
  name?: string
  description?: string
  permission_ids?: number[]
};

export const updateRole = defineApi<UpdateRoleRequest, void>({
  url: `${prefix}/roles/update`,
  method: 'POST'
});

export const listAdminstratedPermissions = defineApi<void, Enum[]>({
  url: `${prefix}/apps/permission/admined-permission`,
  method: 'POST'
});

export type UserBio = {
  id: number
  username: string
};

export const listUsers = defineApi<PagninationRequest, PaginationResponse<UserBio>>({
  url: `${prefix}/users/list`,
  method: 'POST'
});

export const getUserRoles = defineApi<{ id: number }, Enum<number>[]>({ url: `${prefix}/users/user-roles` });

export type BindRoleRequest = {
  user_id: number
  role_ids:number[]
};
export const bindRoles = defineApi<BindRoleRequest, void>({
  url: `${prefix}/users/bind-roles`,
  method: 'POST'
});

export const listRolesAll = defineApi<void, Enum<number>[]>({
  url: `${prefix}/roles/list-all`,
  method: 'POST'
});
