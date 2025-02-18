import React from 'react';

type ErrorUI = string | React.ReactElement;

export interface ValidatorFn {
  // (control: AbstractControl):ValidationErrors | null |(Promise<ValidationErrors | null>) ;
  (value: any): Promise<ErrorUI>
}

export type Path = string | number;

export interface ValidatorRule {
  len?: number;
  message?: ErrorUI;
  min?: number;
  max?: number;
  pattern?: RegExp;
  required?: boolean;
  // whitespace?: boolean
  warning?: boolean;
  validator?: ValidatorFn;
}

export type ValidationError = {
  errors: ErrorUI[]
  warnings: ErrorUI[]
};

export type ValidationErrorWithName = {
  name: Path[]
} & ValidationError;

type ParsedValidatorFn = (value: any) => Promise<ErrorUI[]>;

type BuiltInValidatorMap = {
  [K in keyof ValidatorRule]: (value: any, rule: Omit<ValidatorRule, K> & Required<Pick<ValidatorRule, K>>) => Promise<ErrorUI>;
};
const builtInValidators:BuiltInValidatorMap = {
  len: async (value, rule) => {
    if ('length' in value && typeof rule.len === 'number') {
      if (value.length > rule.len) {
        return Promise.reject(rule.message ?? `The input length should be more than ${rule.min}`);
      }
    }
    return null;
  },
  min: async (value, rule) => {
    const toCompare = Number(value);
    if (!Number.isNaN(toCompare) && typeof rule.min === 'number') {
      if (toCompare < rule.min) {
        return Promise.reject(rule.message ?? `The input should be more than ${rule.min}`);
      }
    }
    return null;
  },
  max: async (value, rule) => {
    const toCompare = Number(value);
    if (!Number.isNaN(toCompare) && typeof rule.max === 'number') {
      if (toCompare > rule.max) {
        return Promise.reject(rule.message ?? `The input should be less than ${rule.max}`);
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
    if ((value === undefined) && rule.required) {
      return Promise.reject(rule.message ?? 'Field is required');
    }
    return null;
  },
  validator: async (value, rule) => {
    if (typeof rule.validator === 'function') {
      return rule.validator(value).catch((reason) => Promise.reject(reason ?? rule.message));
    }
    return null;
  }

};

function parseRule(rule: ValidatorRule): ParsedValidatorFn {
  return function (value:any) {
    const result = Object.keys(rule).map((key) => {
      const validator = builtInValidators[key as keyof ValidatorRule];
      return validator(value, rule as any);
    });
    return Promise.allSettled(result).then((r) => r.filter((p) => p.status === 'rejected').map((p) => p.reason));
  };
}

export type ComposedValidatorFn = (value: any) => Promise<Omit<ValidationError, 'name'>>;

export function composeValidator(rules: ValidatorRule[] = []): ComposedValidatorFn | null {
  if (!rules.length) return null;

  const { errorRules, warningRules } = rules.reduce<{ errorRules: ValidatorRule[], warningRules: ValidatorRule[] }>((last, cur) => {
    last[cur.warning ? 'warnings' : 'errors'].push(cur);
    return last;
  }, { errorRules: [], warningRules: [] });

  const errFns = errorRules.map((e) => parseRule(e));
  const warnFns = warningRules.map((e) => parseRule(e));

  return function (value:any) {
    const errPromises = errFns.map((fn) => fn(value));
    const mergedErrors = Promise.all(errPromises).then((errs) => errs.flat());

    const warnPromises = warnFns.map((fn) => fn(value));
    const mergedWarnings = Promise.all(warnPromises).then((errs) => errs.flat());

    return Promise.all([mergedErrors, mergedWarnings]).then(([errors, warnings]) => ({ errors, warnings }));
  };
}
