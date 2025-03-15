import { ComponentContextType, FontAssetContextType } from './context';
import { FilterType } from '../../panel/meta';
import { CProps } from './basic';
import { PackageConfigType } from '../../component/basic';
import { CanvixResponsePromise, Pagination } from '../../requests';
import { AssetsSearchParams } from '../../panel/assets';

type VariableProp = {
  // store the referred variable id
  config: {
    variableId: string;
    type: 'variable';
    global: boolean;
  };
};
type AssetProp = {
  config: {
    type: 'asset';
    assetId: string;
  };
};

type NullProp = {
  config: null;
};

export type InjectedPropConfig = {
  id: string;
  argName: string;
  order: number;
} & (ValidProp | NullProp);

export type ValidProp = VariableProp | AssetProp;
type ValidPropConfig = ValidProp['config'];

export type SingleFunctionConfig = FilterType & {
  order: number;
  injectedProps: Record<string, InjectedPropConfig>;
};

type Config = PackageConfigType<Record<string, SingleFunctionConfig>> & {
  functionName: string;
  maxFun: number;
  addable: boolean;
  title: string;
  injectable: boolean;
};
export type CFilterProps = CProps<Config>;
export type FilterProps = {
  injectable: boolean;
  filter: SingleFunctionConfig;
  functionArgs: string[];
  isNew?: boolean;
  deleteNewFilter?: () => void;
  handleChange: (payload: Partial<FilterType>) => void;
  handleDelete: () => void;
  changeWithChain: CFilterProps['onChange'];
};
export type State = {
  newFilter: SingleFunctionConfig | null;
};

export type PropsInjectorProps = {
  staticArgs: string[];
  configs: InjectedPropConfig[];
  configMap: Record<string, InjectedPropConfig>;
  onChange: CFilterProps['onChange'];
};

export type PropsInjectorStates = {
  activeId: string | false;
  expanded: boolean;
};

export type PropFormProps = {
  prop: InjectedPropConfig;
  updateArgName(name: string): void;
  updateSource(source: ValidPropConfig): void;
};

export type CFunctionRequiredInterface = {
  getVariableList: (global: boolean) => {
    id: string;
    label: string;
    source: string;
  }[];

  getDataList: (param: {
    page: number;
    size: number;
    sourceName: string;
  }) => PagedPickerItemResponse;

  getDataDetail: (sourceId: string) => CanvixResponsePromise<{ sourceName: string }>;

  getAssetsDetailById: (id: string) => CanvixResponsePromise<{
    name: string;
    fileName: string;
  }>;
};

type GetPickerMethod = (payload: AssetsSearchParams) => PagedPickerItemResponse;
type PagedPickerItemResponse = CanvixResponsePromise<{ data: PickerItemType[] } & Pagination>;

export type PropSelectModalProps = {
  componentContext: ComponentContextType;
  fontAssetContext: FontAssetContextType;
  lastSource: InjectedPropConfig['config'];
  updateSource(source: ValidPropConfig): void;
  getListPackage: Record<string, GetPickerMethod>;
};

export type PickerItemType = {
  sourceId: string;
  sourceName: string;
  type: string;
};

export type AssetItemProps = {
  onClick(itemId: string): void;
} & PickerItemType;

export type PickerProps<T extends 'asset' | 'variable' = 'asset' | 'variable'> = {
  sourceConfig: Extract<ValidPropConfig, { type: T }>;
  /**
   * @param payload the payload shall support fileName as query conditions
   */
  sourceListPagination: GetPickerMethod;
  updateSource(config: Extract<ValidPropConfig, { type: T }>): void;
};
