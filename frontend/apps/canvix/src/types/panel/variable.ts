type VariableBasicType<Source, Detail = null> = {
  [B in {
    [K in keyof IVariableBasicType<Source, Detail>]: IVariableBasicType<
    Source,
    Detail
    >[K] extends null
      ? never
      : K;
  }[keyof IVariableBasicType<Source, Detail>]]: IVariableBasicType<Source, Detail>[B];
};

type IVariableBasicType<Source, Detail = null> = {
  id: string;
  label: string;
  source: Source;
  default: any;
  detail: Detail;
};

type KeyValueConfig = {
  key: string;
};

export type UrlVariableType = VariableBasicType<'url', KeyValueConfig>;

export type StaticVariableType = VariableBasicType<'static', null>;

export type SharedPreferenceVariableType = VariableBasicType<'shared', null>;

export type SqliteVariableType = VariableBasicType<
'sqlite',
{
  query: string;
}
>;

export type LocalVariableType = StaticVariableType;

export type GlobalVariableType =
  | LocalVariableType
  | SqliteVariableType
  | SharedPreferenceVariableType
  | UrlVariableType;

export type GlobalVariableSource = GlobalVariableType['source'];
export type LocalVariableSource = LocalVariableType['source'];

export type GlobalVariableCollection = GlobalVariableType[];
export type LocalVariableCollection = LocalVariableType[];

export type VariableConfigType = GlobalVariableType | LocalVariableType;

export type ScopedVariableConfigType = VariableConfigType & {
  scope: 'global' | 'local';
};

export type GetVariableType<T> = Extract<GlobalVariableType, { source: T }>;

export enum VariableSourceName {
  static = '临时变量',
  url = 'url查询参数',
  shared = '持久化变量',
  sqlite = 'sqlite数据查询',
}

type Updater = (x: any) => void;

export type VariableContextValue = {
  join(variableId: string, updater: Updater): void;
  leave(variableId: string, updater: Updater): void;
  getVariable(variableId: string): any;
  setVariable(variableId: string, value: any): void;
};

export type VariableContextType = React.Context<VariableContextValue>;

export type VariableDepotProps<T> = {
  configs: GetVariableType<T>[];
  children?: React.ReactNode;
};
export interface TypedVariableMutator<T> {
  getValue(conf: GetVariableType<T>): any;
  setValue(conf: GetVariableType<T>, value: any): void;
}

export interface TypedVariableMutatorCtor<T> {
  type: T;
  new (): TypedVariableMutator<T>;
}
