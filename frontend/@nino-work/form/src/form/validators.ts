import type { AbstractStruct } from './model';

export type ValidationErrors = {
  [key: string]: any;
};

export interface ValidatorFn {
  (control: AbstractStruct):ValidationErrors | null |(Promise<ValidationErrors | null>) ;
}

export interface Observable<T> {
  subscribe(fn: (v: T) => void): void
  next(val:T)
}

export interface AsyncValidatorFn {
  // (control: AbstractStruct): Promise<ValidationErrors | null> | Observable<ValidationErrors | null>;
  (control: AbstractStruct): Promise<ValidationErrors | null> | Observable<ValidationErrors | null>;
}

function isPresent(o: any): boolean {
  return o != null;
}

type GenericValidatorFn = (control: AbstractStruct) => any;

function executeValidators<V extends GenericValidatorFn>(
  control: AbstractStruct,
  validators: V[]
): ReturnType<V>[] {
  return validators.map((validator) => validator(control));
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

  return function (control: AbstractStruct) {
    return mergeErrors(executeValidators<ValidatorFn>(control, presentValidators));
  };
}

export function composeValidators(validators: Array< ValidatorFn>): ValidatorFn | null {
  return validators != null ? compose(validators) : null;
}

export function coerceToValidator(validator: ValidatorFn | ValidatorFn[] | null): ValidatorFn | null {
  return Array.isArray(validator) ? composeValidators(validator) : validator || null;
}
