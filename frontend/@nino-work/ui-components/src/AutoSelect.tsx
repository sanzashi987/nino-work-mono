/* eslint-disable no-nested-ternary */
import * as React from 'react';
import CircularProgress from '@mui/material/CircularProgress';
import { MenuItem, Select, SelectProps, Stack } from '@mui/material';

type AutoSelectProps<Value> = {
  requester?: () => Promise<{ label: React.ReactNode, value: Value }[]>
} & SelectProps;

export default function AutoSelect<Value>({ requester, ...others }: AutoSelectProps<Value>) {
  const [loading, setLoading] = React.useState(false);
  const [options, setOptions] = React.useState<{ label: React.ReactNode, value: Value }[]>([]);

  React.useEffect(() => {
    if (requester) {
      setLoading(true);
      requester().then(setOptions).finally(() => setLoading(false));
    }
  }, [requester]);

  return (
    <Select variant="standard" {...others}>
      {requester
        ? loading
          ? (
            <Stack direction="row" justifyContent="center">
              <CircularProgress />
            </Stack>
          )
          // @ts-ignore
          : options.map((e) => <MenuItem key={`${e.value}`} value={e.value}>{e.label}</MenuItem>)
        : others.children}
    </Select>
  );
}
