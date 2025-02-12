/* eslint-disable no-param-reassign */
import React from 'react';

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

export interface WidgetMutate<W, P> {

}

type WidgetStandardProps<T> = {
  value?: T
  onChange?(next: T): void
};

// type Validate

type WatchOption<Store> = {
  form?: FormInstance<Store>
};

type FormInstance<T> = {

};

export type IModel<T, S> = PrimitiveModel<T, S> | ObjectModel<T, S> | ArrayModel<T, S>;

type BaseModel<
  ModelValue,
  StoreValue,
  ToWatch extends React.Key[][] = [],
  StoreToWatch = StoreValue
> = {
  // label: React.ReactNode
  label: string
  field: string
  watch?: readonly [...ToWatch],
  watchOptions?: WatchOption<StoreToWatch>
  callback?(value: any, form: FormInstance<StoreToWatch>): void,
  initialValue?: ModelValue
};

type IsAny<T> = T extends never ? false : true;

type GetModel<V, S> = V extends (infer A)[]
  ? ArrayModel<A, S>
  : V extends object
    ? ObjectModel<V, S>
    : V extends never ? PrimitiveModel<string, S> : PrimitiveModel<V, S>;

export type PrimitiveModel<V, S> = {
  initialValue?: V
} & BaseModel<V, S>;

export type ArrayModel<V, S> = {
  children?: Omit<GetModel<V, S>, 'field'>
} & BaseModel<V[], S>;

export type ObjectModel<V, S> = {
  children?: DefineArrayModels<V>
  initialValue?: V
} & BaseModel<V, S>;

export type ForceCompute<T> = T extends object
  ? {
    [K in keyof T]: T[K]
  }
  : never;

export type ForceDeepCompute<T> = T extends object ? { [K in keyof T]: ForceDeepCompute<T[K]> } : T;

// type MakeStaticArray<Key extends keyof T, T> = { [K in Key]: [ForceDeepCompute<
// Omit<GetModel<T[K]>, 'field'> & { field: K }
// >, ...(Exclude<Key, K> extends "" ? [] : MakeStaticArray<Exclude<Key, K>, T>)] }[Key]
type MakeStaticArray<Key extends keyof T, T> = { [K in Key]: Omit<GetModel<T[K], T>, 'field'> & { field: K } }[Key][];
type DefineArrayModels<T> = T extends object ? MakeStaticArray<keyof T, T> : never;

type Form = {
  a: string
  // // b: number
  b: { test: string }

  c: { test2: string }[]
  // d: number[]
};

type Aa = ForceDeepCompute<DefineArrayModels<Form>>;

export type WidgetIdentifier = keyof WidgetMutate<unknown, unknown>;

type MakeIntersect<T> = (T extends any ? (x: T) => void : never) extends (x: infer R) => void ?
  R : never;

export const model = (): IModel<any, any> & MakeIntersect<WidgetMutate<unknown, unknown>[WidgetIdentifier]> => ({} as any);

export const defineModel = <const T>(models: DefineArrayModels<T>) => models;

// const
export const res = defineModel<Form>([
  { field: 'a', label: 'b' },
  {
    label: 'bb',
    field: 'b',
    children:
      [
        { field: 'test', label: 'test', initialValue: '2' }
      ]
  }, {
    label: 'cc',
    field: 'c',
    children: {
      label: '',
      children: [{
        label: '',
        field: 'test2',
        initialValue: '233'
      }]
    },
    initialValue: [{ test2: '' }]
  }

]);

const buildValueFromSchema = <T>(
  schema: DefineArrayModels<T>[number],
  currentLevelData: any = {}
) => {
  currentLevelData[schema.field] = currentLevelData[schema.field] ?? schema.initialValue;
  if ('children' in schema) {
    if (Array.isArray(schema.children)) {
      const data: any = currentLevelData[schema.field] ?? {};
      currentLevelData[schema.field] = data;
      schema.children.forEach((s: any) => {
        buildValueFromSchema<any>(s, data);
      });
    } else {
      const data: any[] = currentLevelData[schema.field] ?? [];
      currentLevelData[schema.field] = data;
      if (data.length) {
        data.forEach((v, i) => {
          const tempSchmea = { ...schema.children as any, field: i };
          buildValueFromSchema<any>(tempSchmea as any, data);
        });
      }
    }
  }
};

const buildFormDataFromSchemas = <T = any>(schemas: DefineArrayModels<T>, data: any = {}) => {
  schemas.forEach((schmea) => {
    buildValueFromSchema<any>(schmea as any, data);
  });
  return data as T;
};

type FormOptions<T> = {
  initialValue?: Partial<T>
};

export const createForm = <T>(schemas: DefineArrayModels<T>, opts?: FormOptions<T>) => {
  const initialValue = buildFormDataFromSchemas(schemas, opts?.initialValue);

  console.log(initialValue);
};
