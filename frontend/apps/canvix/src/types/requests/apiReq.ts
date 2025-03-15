/* eslint-disable @typescript-eslint/no-namespace */
export type PageType = 'page' | 'template' | 'assets' | 'customComponent' | 'block';
export type User = {
  loginName: string;
  userName: string;
};
export enum Role {
  NoPermission = 0,
  HavePermission = 1,
}
/** 模板/组件 上架状态 0 已上架 1 已下架 */
export enum StateType {
  On = 0,
  Off = 1,
}

/** 0 stands for system built-in palettes, 1 stands for user created palette */
export type PaletteFlag = 0 | 1;

export namespace Request {
  /** ----- Template ----- */
  export type TemplateListPayload = {
    groupCode?: string;
    name: string;
    page: number;
    size: number;
  };

  export type TemplateMetaPayload = {
    id: string;
    groupCode?: string;
    groupName?: string;
    name?: string;
    rootConfig?: string;
    thumbnail?: string;
    font?: string;
    version?: string;
  };

  export type TemplateUpdateGroupPayload = { groupCode: string; groupName: string };
  export type TemplateCreatePayload = { name: string } & Omit<TemplateMetaPayload, 'id'>;
  export type TemplateBatchMovePayload = { groupCode?: string; ids: string[]; groupName?: string };
  export type TemplateImportPayload = { zip: Blob } & Partial<TemplateUpdateGroupPayload>;
  export type TemplateUpdatePayload = { code: string } & TemplateMetaPayload;
  export type TemplateStateUpdatePayload = { ids: string[]; state: StateType };

  /** ----- Page ----- */
  export type PageCreatePayload = TemplateCreatePayload & PageGroupPayload;
  export type PageCreateByTemplatePayload = TemplateCreatePayload & { id: string };
  export type PublishViewParams = { secret: string; publishToken: string; parentId: string };
  export type PagePublishPayload = {
    id: string;
    publishFlag?: number;
    publicSecretKey?: string;
    reference?: number;
  };
  export type PageUpdatePayload = TemplateUpdatePayload;
  export type PageListPayload = TemplateListPayload & PageGroupPayload;
  export type PageGroupPayload = { type: PageType };
  export type PageGroupUpdatePayload = { code: string; name: string };
  export type PageGroupCreatePayload = Pick<PageGroupUpdatePayload, 'name'> & PageGroupPayload;

  /** ----- Workspace ----- */
  export type ModifyUserParams = {
    code: string;
    loginName: string;
    dataDelete?: Role;
    dataSelect?: Role;
    dataUpdate?: Role;
    manager?: Role;
  };

  export type DeleteUserParams = {
    code: string;
    loginName: string;
  };

  export type SearchUserParams = {
    pageSize: number;
    pageNum: number;
    condition: string;
  };
  export type UpdateParams = {
    code?: string;
    name?: string;
    pageNumber?: number;
  };

  export type AddUserParams = {
    code: string;
    projectUsers: User[];
  };

  export type ProjectUsersParams = {
    code: string;
    page: number;
    size: number;
  };

  /** ----- Assets ----- */
  export type AssetsUploadPayload = Partial<TemplateUpdateGroupPayload> & {
    type?: 'design' | 'font' | 'data' | 'cover';
  };
  export type AssetsUpdateGroupPayload = Partial<TemplateUpdateGroupPayload> & {
    fileIds: string[];
  };
  export type AssetsReplacePayload = { file: Blob; fileId: string };

  /** -----  Components / Custom Components ----- */
  export type ComponentGroupListPayload = {
    state?: StateType;
    type?: 'com' | 'subcom';
  };
  export type ComponentUpdatePayload = Partial<TemplateUpdateGroupPayload> & {
    id: string;
    zip?: Blob;
  };
  export type ComponentListPayload = Omit<TemplateListPayload, 'name'> & { cn_name?: string };
  export type ComponentImportPayload = Pick<TemplateImportPayload, 'zip'>;

  /** -----  Fonts ----- */
  export type FontListPayload = { assetName: string; page: number; size: number };
  export type FontAllListPayload = Pick<FontListPayload, 'assetName'>;
  export type FontUpdatePayload = { file?: Blob; assetCode: string; assetName?: string };
  export type FontCreatePayload = { file: Blob; assetName: string };

  /** -----  DataSource ----- */
  export type DataSourceSourceListPayload = { sourceType: string };
  export type DataSourceListPayload = {
    page: number;
    size: number;
    sourceName: string;
    sourceType?: string;
  };
  export type DataSourceSearchPayload = { sourceType: string[]; search: string };
  export type DataSourceReplacePayload = { search: string; target: string; sourceId: string[] };
  export type DataSourceFileDataPayload = { sourceId: string; user: string };
  export type DataSourceFindTablesParams = DataSourceFindPatternParams & {
    partternName?: string; // pg/sqlserver only
  };
  export type DataSourceFindPatternParams = {
    dbName: string;
    sourceId: string;
  };
  /** Theme & Palette */
  export type PaletteCreatePayload = {
    name: string;
    theme: string;
    flag: PaletteFlag;
  };
  export type PaletteUpdatePayload = {
    id: number;
  } & Partial<PaletteCreatePayload>;
}
