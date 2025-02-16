import React from 'react';
import type { AbstractControl } from './model';

export interface ValidatorFn {
  // (control: AbstractControl):ValidationErrors | null |(Promise<ValidationErrors | null>) ;
  (value: any): Promise<React.ReactNode>
}

export type Path = string | number;

type ErrorUI = string | React.ReactElement;

export interface Rules {
  len?: number;
  max?: number;
  message?: ErrorUI;
  min?: number;
  pattern?: RegExp;
  required?: boolean;
  whitespace?: boolean
  warning?: boolean;
  validator?: ValidatorFn;
}

type ValidationError = {
  name:Path[]
  errors: ErrorUI[]
  warnings:ErrorUI[]
};

export type ValidationErrors = ValidationError[];

export interface Observable<T> {
  subscribe(fn: (v: T) => void): void
  next(val:T)
}

export interface AsyncValidatorFn {
  // (control: AbstractStruct): Promise<ValidationErrors | null> | Observable<ValidationErrors | null>;
  (control: AbstractControl): Promise<ValidationErrors | null> | Observable<ValidationErrors | null>;
}

function isPresent(o: any): boolean {
  return o != null;
}

type GenericValidatorFn = (value:any) => any;

function runValidators<V extends GenericValidatorFn>(
  control: AbstractControl,
  validators: V[]
): Promise<ErrorUI[]> {
  const val = validators.map((validator) => Promise.resolve().then(() => validator(control)));
  return Promise.allSettled(val)
    .then((res) => res.filter((p) => p.status === 'rejected')
      .map((e) => e.reason));
}

function mergeErrors(arrayOfErrors: (ValidationErrors | null)[]): ValidationErrors | null {
  let res: { [key: string]: any } = {};
  arrayOfErrors.forEach((errors: ValidationErrors | null) => {
    res = errors != null ? { ...res!, ...errors } : res!;
  });

  return Object.keys(res).length === 0 ? null : res;
}

function compose(validators: (ValidatorFn | null | undefined)[] | null): ValidatorFn | null {
  if (!validators) return null;
  const presentValidators: ValidatorFn[] = validators.filter(isPresent) as any;
  if (presentValidators.length === 0) return null;

  return function (control: AbstractControl) {
    return mergeErrors(runValidators<ValidatorFn>(control, presentValidators));
  };
}

export function composeValidators(validators: Array< ValidatorFn>): ValidatorFn | null {
  return validators != null ? compose(validators) : null;
}
