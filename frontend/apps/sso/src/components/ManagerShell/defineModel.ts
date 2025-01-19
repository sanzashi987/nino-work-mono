import { TableCellProps } from '@mui/material';

export type Model<T = string> = {
  label: string
  field: string,
  headerCellProps?: TableCellProps
  dataCellProps?: TableCellProps & {
    render: (row: any, i: number) => React.ReactNode
  }
};

const defineModel = <T>(m: Model<T>): Model<T> => m;
