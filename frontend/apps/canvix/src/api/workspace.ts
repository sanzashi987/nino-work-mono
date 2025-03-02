import { Request } from '@canvas/types';

/** 获取空间列表 */
export const getProjects = () => authRequest.GET('/project/myProject');
/** 获取单个空间详细信息 */
export const getProjectDetail = (projectCode: string) => authRequest.GET(`/project/info/${projectCode}`);
/** 新建空间 */
export const createProject = (name: string) => authRequest.POST('/project/create', { name });
/** 获取单个空间详细信息 */
export const searchUser = (params: Request.SearchUserParams) => authRequest.POST('common/pageUserlist', params);
/** 删除空间 */
export const deleteProject = (code: string) => authRequest.DELETE(`/project/delete/${code}`);
/** 更新空间信息（名称、配额） */
export const updateProject = (params: Request.UpdateParams) => authRequest.POST('/project/update', params);
/** 添加空间成员 */
export const addUser = (params: Request.AddUserParams) => authRequest.POST('/project/addUser', params);
/** 获取选中空间成员 */
export const getProjectUsers = (params: Request.ProjectUsersParams) => authRequest.POST('/project/projectUsers', params);
/** 修改空间成员权限 */
export const modifyUserPromission = (params: Request.ModifyUserParams) => authRequest.POST('/project/editUser', params);
/** 删除空间成员 */
export const deleteUser = (params: Request.DeleteUserParams) => authRequest.POST('/project/removeUser', params);

export const changeProject = (code: string) => authRequest.GET(`/project/changeProject/${code}`);
