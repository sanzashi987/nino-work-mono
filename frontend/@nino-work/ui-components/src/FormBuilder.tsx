import React, { useEffect, useMemo, useRef, useState } from 'react';
import { FieldValues, useForm, UseFormReturn } from 'react-hook-form';
import Grid from '@mui/material/Grid2';
import { Stack, TextField } from '@mui/material';
import { FormCommonLayout, Model } from './ManagerShell/defineModel';
import FormLabel from './FormLabel';

export type FormBuilderProps<FormData extends FieldValues, T = any> = {
  schema: Model<T>[];
  onSubmit?: (d?: FormData) => Promise<any>;
  form?: UseFormReturn<FormData, any, undefined>;
  spacing?: number;
} & FormCommonLayout;

// TODO add validator
const FormBuilder = <FormData extends FieldValues, T = any>(
  props: FormBuilderProps<FormData, T>
) => {
  const { schema, form, colSpan = 12, spacing = 2, layout = 'column' } = props;
  const instance = useForm<FormData>();
  const [runtimeSchema, setRuntimeSchema] = useState(schema);
  const runtimeForm = form ?? instance;
  const { subscribe } = runtimeForm;
  const runtimeSchemaRef = useRef(runtimeSchema);
  runtimeSchemaRef.current = runtimeSchema;

  const fieldIdxMap = useMemo(() => {
    const map = new Map<string, number>();
    schema.forEach((model, idx) => {
      map.set(model.field, idx);
    });
    return map;
  }, [schema]);

  const widgets = useMemo(
    () =>
      runtimeSchema.map(s => {
        const { label, field, formCellProps = {} } = s;
        const lay = formCellProps.layout ?? layout;
        const size = formCellProps.colSpan ?? colSpan;
        const widget = formCellProps.widget ?? TextField;
        if (formCellProps.hidden) {
          return null;
        }

        const dom = React.createElement(widget, {
          id: field,
          /* @ts-expect-error the field will be the path string */
          ...runtimeForm.register(field),
          ...(formCellProps.widgetProps ?? {}),
        });
        if (formCellProps.noStyle) {
          return dom;
        }

        return (
          <Grid key={field} size={size}>
            <Stack direction={lay}>
              <FormLabel title={label} field={field} />
              {dom}
            </Stack>
          </Grid>
        );
      }),
    [colSpan, runtimeForm, layout, runtimeSchema]
  );

  useEffect(() => {
    const toUbsubscribe = schema.map((model, idx) => {
      return subscribe({
        name: model.formCellProps!.watch as any[],
        callback: ({ values }) => {
          const valueArr = model.formCellProps!.watch!.map(w => values[w]);
          const publisherFormCellProps = model.formCellProps!.watch!.map(cur => {
            const curIdx = fieldIdxMap.get(cur)!;
            return runtimeSchemaRef.current![curIdx].formCellProps;
          });

          const res = model.formCellProps!.callback?.({
            values: valueArr,
            form: runtimeSchemaRef as any,
            selfFormCellProps: model.formCellProps,
            publisherFormCellProps,
          });
          setRuntimeSchema(last => {
            const newSchema = [...last];
            newSchema[idx] = { ...model, formCellProps: { ...model.formCellProps, ...res } };
            return newSchema;
          });
        },
      });
    });

    return () => {
      toUbsubscribe.forEach(ub => ub());
    };
  }, [schema]);

  return (
    <form onSubmit={runtimeForm.handleSubmit(props.onSubmit!)}>
      <Grid spacing={spacing ?? 2} container>
        {widgets}
      </Grid>
    </form>
  );
};

export default FormBuilder;
