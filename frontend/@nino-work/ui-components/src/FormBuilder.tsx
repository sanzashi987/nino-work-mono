/* eslint-disable no-param-reassign */
import React, { useEffect, useMemo, useState } from 'react';
import { useForm, UseFormReturn, FormProvider } from 'react-hook-form';
import { noop } from '@nino-work/shared';
import Grid from '@mui/material/Grid2';
import { Stack, TextField } from '@mui/material';
import { FormCommonLayout, Model } from './ManagerShell/defineModel';
import FormLabel from './FormLabel';

type FormBuilderProps<FormData, T = any> = {
  schema: Model<T>[],
  // onSubmit: (d: FormData) => Promise<any>
  formRef?: React.RefObject<UseFormReturn<FormData, any, undefined>>
  spacing?: number

} & FormCommonLayout;

const FormBuilder = <FormData, T = any>(props:FormBuilderProps<FormData, T>) => {
  const instance = useForm<FormData>();
  const { schema, formRef, colSpan = 12, spacing = 2, layout = 'column' } = props;
  useEffect(() => {
    if (formRef) {
      formRef.current = instance;
      return () => { formRef.current = null; };
    }
    return noop;
  }, [formRef]);

  const widgets = useMemo(() => schema.map((s) => {
    const { label, field, formCellProps } = s;
    const lay = formCellProps.layout ?? layout;
    const size = formCellProps.colSpan ?? colSpan;
    const widget = formCellProps.widget ?? TextField;
    return (
      <Grid key={field} size={size}>
        <Stack direction={lay}>
          <FormLabel title={label} field={field} />
          {/* @ts-ignore */ }
          {React.createElement(widget, { id: field, ...instance.register(field) })}
        </Stack>
      </Grid>
    );
  }), [colSpan, instance, layout, schema]);

  return (
    <Grid spacing={spacing ?? 2} container>
      {widgets}
    </Grid>
  );
};

export default FormBuilder;
