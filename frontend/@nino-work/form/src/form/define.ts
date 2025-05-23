/* eslint-disable @typescript-eslint/no-use-before-define */
import React from 'react';
import type { ValueEqualityFn } from '../signal/equality';
import type { ValidatorRule } from './validators';
import FormObject from './form_object';
import FormArray from './form_array';
import FormPrimitive from './form_primitive';

export type SupportedValidationValue = boolean | number | string | RegExp;

export type ValidationRule<
  ValidationValue extends SupportedValidationValue = SupportedValidationValue,
> = ValidationValue | ValidationWithMessage<ValidationValue>;

export type ValidationWithMessage<
  ValidationValue extends SupportedValidationValue = SupportedValidationValue,
> = {
  value: ValidationValue;
  message: string;
};

export interface WidgetMutate<W, P> {}

export type WidgetStandardProps<T> = {
  id?: string;
  value?: T;
  onChange?(next: T): void;
};

// type Validate

type WatchOption<Store extends {}> = {
  form?: FormObject<Store>;
};

export type IModel<T, S> = PrimitiveModel<T, S> | ObjectModel<T, S> | ArrayModel<T, S>;

export type BaseModel<
  ModelValue,
  StoreValue extends {},
  ToWatch extends React.Key[][] = [],
  StoreToWatch extends StoreValue = StoreValue,
> = {
  // label: React.ReactNode
  label?: string;
  field: string;
  watch?: readonly [...ToWatch];
  watchOptions?: WatchOption<StoreToWatch>;
  widget?: any;
  widgetProps?: any;
  formItemProps?: {
    initialValue?: ModelValue;
    rules?: ValidatorRule[];
    equality?: ValueEqualityFn<ModelValue>;
  };

  callback?(value: any, form: FormObject<StoreToWatch>): void;
};

type GetModel<V, S> = V extends (infer A)[]
  ? ArrayModel<A, S>
  : V extends object
    ? ObjectModel<V, S>
    : V extends never
      ? PrimitiveModel<string, S>
      : PrimitiveModel<V, S>;

export type PrimitiveModel<V, S> = BaseModel<V, S>;

export type ArrayModel<V, S> = {
  children?: Omit<GetModel<V, S>, 'field'>;
} & BaseModel<V[], S>;

export type ObjectModel<V, S> = {
  children?: DeriveChildrenForObject<V, S>;
} & BaseModel<V, S>;

export type ForceCompute<T> = T extends object
  ? {
      [K in keyof T]: T[K];
    }
  : never;

export type ForceDeepCompute<T> = T extends object ? { [K in keyof T]: ForceDeepCompute<T[K]> } : T;

// type aa = GetModel<number | { a:number }, any>;
// type MakeStaticArray<Key extends keyof T, T> = { [K in Key]: [ForceDeepCompute<
// Omit<GetModel<T[K]>, 'field'> & { field: K }
// >, ...(Exclude<Key, K> extends "" ? [] : MakeStaticArray<Exclude<Key, K>, T>)] }[Key]
type MakeObjectChildren<Key extends keyof T, T, S> = {
  [K in Key]: Omit<GetModel<T[K], S>, 'field'> & { field: K };
}[Key][];
type DeriveChildrenForObject<V, S = V> = V extends object
  ? MakeObjectChildren<keyof V, V, S>
  : never;
type MakeArrayChildren<V, S> = Omit<GetModel<V, S>, 'field'>[];
// type DeriveChildrenForArray<V, S> = V extends (infer A)[] ? MakeArrayChildren<A, S> : never;

type Form = {
  a: string;
  // // b: number
  b: { test: string };

  c: { test2: string }[];
  d: number[];
};

type Aa = ForceDeepCompute<DeriveChildrenForObject<Form>>;

export type WidgetIdentifier = keyof WidgetMutate<unknown, unknown>;

type MakeIntersect<T> = (T extends any ? (x: T) => void : never) extends (x: infer R) => void
  ? R
  : never;

export const model = (): IModel<any, any> &
  MakeIntersect<WidgetMutate<unknown, unknown>[WidgetIdentifier]> => ({}) as any;

export const defineModel = <const T>(models: DeriveChildrenForObject<T>) => models;

// const
export const res = defineModel<Form>([
  { field: 'a', label: 'b' },
  {
    label: 'bb',
    field: 'b',
    children: [{ field: 'test', label: 'test', formItemProps: { initialValue: '2' } }],
  },
  {
    label: 'cc',
    field: 'c',
    children: {
      label: '',
      children: [
        {
          label: '',
          field: 'test2',
          formItemProps: { initialValue: '233' },
        },
      ],
    },
    formItemProps: { initialValue: [{ test2: '' }] },
  },
  { label: 'd', field: 'd', children: { label: '2' } },
]);

const buildValueFromSchema = <T>(
  schema: DeriveChildrenForObject<T>[number],
  currentLevelData: any = {}
) => {
  currentLevelData[schema.field] =
    currentLevelData[schema.field] ?? schema.formItemProps?.initialValue;
  if ('children' in schema) {
    if (Array.isArray(schema.children)) {
      const data: any = currentLevelData[schema.field] ?? {};
      currentLevelData[schema.field] = data;
      buildFormDataFromSchemas<any>(schema.children, data);
    } else {
      const data: any[] = currentLevelData[schema.field] ?? [];
      currentLevelData[schema.field] = data;
      if (data.length) {
        data.forEach((v, i) => {
          const tempSchmea = { ...(schema.children as any), field: i };
          buildValueFromSchema<any>(tempSchmea as any, data);
        });
      }
    }
  }
};

const buildFormDataFromSchemas = <T = any>(schemas: DeriveChildrenForObject<T>, data: any = {}) => {
  schemas.forEach(schmea => {
    buildValueFromSchema<any>(schmea as any, data);
  });
  return data as T;
};

type FormOptions<T> = {
  initialValue?: Partial<T>;
};

export const createForm = <T>(schemas: DeriveChildrenForObject<T>, opts?: FormOptions<T>) => {
  const initialValue = buildFormDataFromSchemas(schemas, opts?.initialValue);

  console.log(initialValue);
};

export function decideControl(m: IModel<any, any>, data?: any) {
  const initValue = data ?? m.formItemProps.initialValue;

  if ('children' in m) {
    if (typeof m.children === 'object') {
      if (Array.isArray(m.children)) {
        return new FormObject(m as any, initValue);
      }
      return new FormArray(m as any, initValue);
    }
  }
  if (typeof initValue === 'object') {
    if (Array.isArray(initValue)) {
      return new FormArray(m as any, initValue);
    }
    return new FormObject(m as any, initValue);
  }
  return new FormPrimitive(m as any, initValue);
}
