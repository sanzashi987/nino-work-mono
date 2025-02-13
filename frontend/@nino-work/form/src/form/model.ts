import { effect, signal, untracked } from '../signal';
import type { ValidationErrors, ValidatorFn } from './validators';
import { coerceToValidator } from './validators';

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
  /** validators */
  public errors: ValidationErrors | null = null;

  private _composedValidatorFn: ValidatorFn | null;

  private _rawValidators: ValidatorFn | ValidatorFn[] | null;

  private _composeValidators(validators: ValidatorFn | ValidatorFn[] | null) {
    this._rawValidators = Array.isArray(validators) ? validators.slice() : validators;
    this._composedValidatorFn = coerceToValidator(this._rawValidators);
  }

  private _parent: AbstractStruct | AbstractStruct | null;

  private updateStrategy: FormHooks = 'change';

  abstract _forEachChild(cb:(c:AbstractStruct)=>void):void;
  abstract _allControlsDisabled(): boolean;

  /** dirty */
  private _dirtyReactive = signal(false);

  get dirty() {
    return untracked(this._dirtyReactive);
  }

  set dirty(val: boolean) {
    untracked(() => this._dirtyReactive.set(val));
  }

  setParent(p: AbstractStruct | null) {
    this._parent = p;
  }

  get parent() {
    return this._parent;
  }

  /** values */
  readonly defaultValue: TValue | null;

  private valueReactive = signal<TValue | undefined>(undefined);

  get value() {
    return untracked(() => this.valueReactive());
  }

  getRawValue(): any {
    return this.value;
  }
  abstract setValue(value: TRawValue, options?: Object): void;
  abstract patchValue(value: TValue, options?: Object): void;
  abstract reset(value?: TValue, options?: Object): void;
  abstract _updateValue(): void;

  _find(name: string | number): AbstractStruct | null {
    return null;
  }

  private readonly statusReactive = signal<FormControlStatus | undefined>(undefined);

  get status(): FormControlStatus {
    return untracked(this.statusReactive)!;
  }

  private set status(v: FormControlStatus) {
    untracked(() => this.statusReactive.set(v));
  }

  abstract _anyControls(fn:(c: AbstractStruct) => boolean):boolean;

  get root() {
    let x = this;
    while (x.parent) {
      // @ts-ignore
      x = x.parent;
    }
    return x;
  }

  watchValue(cb: (v: TValue) => void): VoidFunction {
    const { destroy } = effect(() => {
      const value = this.valueReactive();
      cb(value);
    });
    return destroy;
  }

  get<P extends string | readonly (string | number)[]>(path: P): AbstractStruct<any> | null;
  get<P extends string | (string | number)[]>(path: P): AbstractStruct<any> | null {
    let currPath: Array<string | number> | string = path;
    if (currPath == null) return null;
    if (!Array.isArray(currPath)) currPath = currPath.split('.');
    if (currPath.length === 0) return null;
    return currPath.reduce(
      (control: AbstractStruct | null, name) => control && control._find(name),
      this
    );
  }

  private _setInitialStatus() {
    this.status = this._allControlsDisabled() ? FormControlStatus.DISABLED : FormControlStatus.VALID;
  }
}
function coerceToAsyncValidator(_rawAsyncValidators: any): ValidatorFn {
  throw new Error('Function not implemented.');
}
