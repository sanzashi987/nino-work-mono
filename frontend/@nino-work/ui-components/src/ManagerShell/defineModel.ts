import { TableCellProps } from '@mui/material';

export type FormCommonLayout = {
  colSpan?: number
  layout?:'row' | 'column'
};

export type Model<T = string> = {
  label: string
  field: string,
  formCellProps?: FormCommonLayout & {
    widget?: React.ComponentType,
    widgetProps?: any
    type?: 'hidden'
  },
  headerCellProps?: TableCellProps
  dataCellProps?: TableCellProps & {
    render?: (row: any, i: number) => React.ReactNode
  }
};

const defineModel = <T>(m: Model<T>): Model<T> => m;
