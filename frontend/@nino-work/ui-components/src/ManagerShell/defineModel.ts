import { TableCellProps } from '@mui/material';
import { UseFormReturn } from 'react-hook-form';

export type FormCommonLayout = {
  colSpan?: number;
  layout?: 'row' | 'column';
};

type WatcherParam<T> = {
  values: any[];
  form: UseFormReturn<any, any, undefined>;
  selfFormCellProps: Model<T>['formCellProps'];
  publisherFormCellProps: Model<any>['formCellProps'][];
};

export type CrossModelWatch<T = string> = {
  watch?: string[];
  callback?: (watcherParam: WatcherParam<T>) => Partial<Model<T>['formCellProps']>;
};

export type Model<T = string> = {
  label: string;
  field: string;
  formCellProps?: FormCommonLayout & {
    widget?: React.ComponentType<any>;
    widgetProps?: any;
    hidden?: boolean;
    noStyle?: boolean;
  } & CrossModelWatch<T>;
  headerCellProps?: TableCellProps;
  dataCellProps?: TableCellProps & {
    render?: (row: any, i: number) => React.ReactNode;
  };
};

const defineModel = <T>(m: Model<T>): Model<T> => m;
