import React, { useMemo } from 'react';
import { useForm, UseFormReturn } from 'react-hook-form';
import Grid from '@mui/material/Grid2';
import { Stack, TextField } from '@mui/material';
import { FormCommonLayout, Model } from './ManagerShell/defineModel';
import FormLabel from './FormLabel';

export type FormBuilderProps<FormData extends {}, T = any> = {
  schema: Model<T>[],
  onSubmit?: (d: FormData) => Promise<any>
  form?:UseFormReturn<FormData, any, undefined>,
  spacing?: number
} & FormCommonLayout;

// TODO add validator
const FormBuilder = <FormData, T = any>(props: FormBuilderProps<FormData, T>) => {
  const instance = useForm<FormData>();
  const { schema, form, colSpan = 12, spacing = 2, layout = 'column' } = props;
  const usedForm = form ?? instance;

  const widgets = useMemo(() => schema.map((s) => {
    const { label, field, formCellProps = {} } = s;
    const lay = formCellProps.layout ?? layout;
    const size = formCellProps.colSpan ?? colSpan;
    const widget = formCellProps.widget ?? TextField;
    const hidden = formCellProps.type === 'hidden';
    if (hidden) {
      /* @ts-ignore */
      return <input id={field} {...usedForm.register(field)} type="hidden" />;
    }

    return (
      <Grid key={field} size={size}>
        <Stack direction={lay}>
          <FormLabel title={label} field={field} />
          {/* @ts-ignore */ }
          {React.createElement(widget, { id: field, ...usedForm.register(field), ...(formCellProps.widgetProps ?? {}) })}
        </Stack>
      </Grid>
    );
  }), [colSpan, usedForm, layout, schema]);

  return (
    <form onSubmit={usedForm.handleSubmit(props.onSubmit)}>
      <Grid spacing={spacing ?? 2} container>
        {widgets}
      </Grid>
    </form>
  );
};

export default FormBuilder;
