import React from 'react';

export interface ValidatorFn {
  // (control: AbstractControl):ValidationErrors | null |(Promise<ValidationErrors | null>) ;
  (value: any): Promise<React.ReactNode>
}

export type Path = string | number;

type ErrorUI = string | React.ReactElement;

export interface ValidatorRule {
  // len?: number;
  message?: ErrorUI;
  min?: number;
  max?: number;
  pattern?: RegExp;
  required?: boolean;
  whitespace?: boolean
  warning?: boolean;
  validator?: ValidatorFn;
}

type ValidationError = {
  name:Path[]
  errors: ErrorUI[]
  warnings: ErrorUI[]
};

export type ValidationErrors = ValidationError[];

type ValidatorMap = (value:any, rule:ValidatorRule)=>Promise<React.ReactNode>;

type BuiltInValidatorMap = {

  [K in keyof ValidatorRule]: (value: any, rule: Omit<ValidatorRule, K> & Required< Pick<ValidatorRule, K>>) => Promise<React.ReactNode>;

};
const builtInValidators:BuiltInValidatorMap = {
  min: async (value, rule) => {
    if ('length' in value && typeof rule.min === 'number') {
      if (value.length < rule.min) {
        return Promise.reject(rule.message ?? `The input length should be more than ${rule.min}`);
      }
    }
    return null;
  },
  max: async (value, rule) => {
    if ('length' in value && typeof rule.max === 'number') {
      if (value.length > rule.max) {
        return Promise.reject(rule.message ?? `The input length should be less than ${rule.max}`);
      }
    }
    return null;
  },
  pattern: async (value, rule) => {
    if (typeof value === 'string' && rule.pattern instanceof RegExp) {
      const pass = rule.pattern.test(value);
      if (!pass) {
        return Promise.reject(rule.message ?? `Faile to pass the string test: ${rule.pattern}`);
      }
    }
    return null;
  },

  required: async (value, rule) => {
    if (!value && rule.required) {
      return Promise.reject(rule.message ?? 'Field is required');
    }
    return null;
  }

};

function parseRule(rule: ValidatorRule): ValidatorFn {
  return function (value:any) {
    const result = Object.keys(rule).map((key) => {
      const validator = builtInValidators[key as keyof ValidatorRule];
      return validator(key, rule as any);
    });
    return Promise.allSettled(result).then((r) => r.filter((p) => p.status === 'rejected'));
  };
}

export function composeValidator(rules: ValidatorRule[] = []):ValidatorFn | null {
  if (!rules.length) return null;

  const { errors, warnings } = rules.reduce<{ errors: ValidatorRule[], warnings: ValidatorRule[] }>((last, cur) => {
    last[cur.warning ? 'warnings' : 'errors'].push(cur);
    return last;
  }, { errors: [], warnings: [] });

  const errFns = errors.map((e) => parseRule(e));
  const warnFns = warnings.map((e) => parseRule(e));

  return function (value:any) {
    const errPromises = errFns.map((fn) => fn(value));
    const warnPromises = warnFns.map((fn) => fn(value));
  };
}
