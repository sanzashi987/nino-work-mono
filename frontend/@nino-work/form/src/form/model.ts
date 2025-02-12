/* eslint-disable no-underscore-dangle */
import { signal, untracked } from '../signal';

export const enum FormControlStatus {
  VALID = 'VALID',
  INVALID = 'INVALID',
  PENDING = 'PENDING',
  DISABLED = 'DISABLED'
}

export type FormHooks = 'change' | 'blur' | 'submit';

export type IsAny<T, Y, N> = 0 extends 1 & T ? Y : N;
export type TypedOrUntyped<T, Typed, Untyped> = IsAny<T, Untyped, Typed>;
export type FormValue<T extends AbstractStruct | undefined> =
  T extends AbstractStruct<any, any> ? T['value'] : never;
export type FormRawValue<T extends AbstractStruct | undefined> =
  T extends AbstractStruct<any, any>
    ? T['setValue'] extends (v: infer R) => void
      ? R
      : never
    : never;

export abstract class AbstractStruct<TValue = any, TRawValue extends TValue = TValue> {
  readonly defaultValue: TValue | null;

  private _parent: AbstractStruct | AbstractStruct | null;

  private valueReactive = signal<TValue | undefined>(undefined);

  private updateStrategy: FormHooks = 'change';

  abstract iterChild(cb:(c:AbstractStruct)=>void):void;

  abstract setValue(value: TRawValue, options?: Object): void;
  abstract patchValue(value: TValue, options?: Object): void;
  abstract reset(value?: TValue, options?: Object): void;

  setParent(p: AbstractStruct | null) {
    this._parent = p;
  }

  get parent() {
    return this._parent;
  }

  get value() {
    return untracked(() => this.valueReactive());
  }

  private readonly statusReactive = signal<FormControlStatus | undefined>(undefined);

  get status(): FormControlStatus {
    return untracked(this.statusReactive)!;
  }

  private set status(v: FormControlStatus) {
    untracked(() => this.statusReactive.set(v));
  }

  get valid() {
    return this.status === FormControlStatus.VALID;
  }

  get invalid() {
    return this.status === FormControlStatus.INVALID;
  }

  get disabled() {
    return this.status === FormControlStatus.DISABLED;
  }

  get enabled() {
    return this.status !== FormControlStatus.DISABLED;
  }

  get root() {
    let x = this;
    while (x.parent) {
      // @ts-ignore
      x = x.parent;
    }
    return x;
  }
}
