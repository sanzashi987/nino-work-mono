/* eslint-disable @typescript-eslint/no-namespace */
/**
 * ****************************************************
 * Basic API Response Type
 * ****************************************************
 */
export namespace Response {
  export type BasicPageResponseType = {
    /** 页面唯一标识 */
    id: string;
    code: string;
    /** 创建时间 */
    createTime: string;
    /** 更新时间 */
    updateTime: string;
    font: string;
    /** 分组code */
    groupCode: string;
    /** 名称 */
    name: string;
    /** 工作空间ID */
    projectCode: string | null;
    /** 用户 */
    userIdentify: string;
    /** 项目配置 */
    rootConfig: string | null;
    /** 项目对应的平台版本号，每次保存项目时更新 */
    version: null | string;
    /** 缩略图  */
    thumbnail: string;
  };

  /** PC版模板返回数据格式 */
  export type TemplateResponseType = BasicPageResponseType & {
    /** 能访问该模板的用户版本，大容量版、使用版等  */
    label: string;
    /** 分辨率，格式为"{width:100,height:200}" */
    resolutioRatio: string | null;
    /** 是否上架，仅为1时上架 */
    state: null | 0 | 1;
    permission: boolean;
  };

  /** PC版模块返回数据格式 */
  export type BlockResponseType = TemplateResponseType;

  /** PC版资产返回数据格式 */
  export type AssetsResponseType = {
    createTime: string;
    createUser: string;
    fileId: string;
    fileName: string;
    groupCode: string;
    groupName: string;
    mimeType: string;
    projectCode: string;
    size: number;
    suffix: string | null;
    updateTime: string;
  };

  /** PC版页面返回数据格式 */
  export type ScreenResponseType = BasicPageResponseType & {
    type: 'page';
    /** 0：未发布 1：公开发布 2：加密发布 3：token验证 */
    publishFlag: number;
    /** 加密发布密码 */
    publishSecretKey: string | null;
    /** 引用面板相关权限 0：禁止，1：仅自已引用，2：公开引用 */
    reference: number;
    /** 是否管理员 */
    isManager: boolean;
    /** @deprecated */
    config: null | string;
    /** @deprecated */
    dashboardComponents: null | string;
    /** @deprecated */
    interaction: null | string;
    /** @deprecated */
    layerZIndexList: null | string;
    /** @deprecated */
    filters: null | string;
  };

  // export type PageCardType =Partial<PC_BlockResponsiveType> & Partial<PC_TemplateResponsiveType> & Partial<PC_TemplateResponsiveType> & BasicPageResponsiveType;
  /** 页面数据源类型 */
  export type PageInfo = BlockResponseType & ScreenResponseType & TemplateResponseType;

  /** 分组信息 */
  export type GroupResponseType = {
    allCount: number;
    unCount: number;
    lists: Array<{
      code: string;
      count: number;
      name: string;
      createTime: string;
      updateTime: string;
    }>;
  };

  /** 单个自定义组件返回数据格式 */
  export type CustomComponentResponseType = {
    cn_name: string;
    category: string;
    // code: null;
    createTime: string;
    groupCode: string;
    icon: string;
    id: string;
    name: string;
    projectCode: string;
    type: 'com' /* | 'logical' */ | 'subcom';
    updateTime: string;
    userIdentify: string;
    version: string;
  };

  /** ----- 资产返回格式 ----- */
  export type AssetBasicType = {
    fileId: string;
    mimeType: string;
    name: string;
    size: number;
    suffix: string;
  };

  export type AssetsMetaType = Exclude<AssetBasicType, 'name'> & {
    createTime: string;
    updateTime: string;
    fileName: string;
    groupCode: string;
  };

  export type FontResponseType = {
    /** id */
    assetCode: string;
    fileId: string;
    /** 名称 */
    assetName: string;
    projectCode: string;
    createUser: string;
    updateUser: string;
    createTime: string;
    updateTime: string;
  };

  /** -----  Datasource ----- */
  export type ConnectUpdateType = {
    /** 数据源名称 */
    sourceName: string;
    /** 数据源类型 */
    sourceType: string;
    /**
     * 详细配置内容
     * @description 具体配置内容取决于sourceType
     */
    sourceInfo: string;
    /**
     * 数据源id
     * @description 仅修改时需要
     */
    sourceId?: string;
    /**
     * @description 静态文件数据源 文件上传
     */
    file?: Blob;
  };

  export type ConnectResponseType = Required<ConnectUpdateType> & {
    createTime: string;
    updateTime: string;
    userIdentify: string;
    projectCode: string | null;
  };
}
