import { FormStatus } from 'react-dom';
import { effect, signal, untracked } from '../signal';
import type { ValidationErrors, ValidatorFn } from './validators';
import { composeValidators } from './validators';

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

  private _pendingValidator = null;

  private _composeValidators(validators: ValidatorFn | ValidatorFn[] | null) {
    const validatorArr = Array.isArray(validators) ? validators.slice() : validators;
    this._rawValidators = validatorArr;
    this._composedValidatorFn = Array.isArray(validatorArr) ? composeValidators(validatorArr) : validatorArr || null;
  }

  private runValidator(shouldHaveEmitted: boolean, emitEvent?: boolean) {
    if (this._composedValidatorFn) {
      this.status = FormControlStatus.PENDING;
      Promise.resolve().then(() => this._composedValidatorFn(this))
        .then((err) => {
          this.errors = err;
          this._updateErrors(emitEvent, this, shouldHaveEmitted);
        })
        .finally();
    }
  }

  private _updateErrors(emitEvent: boolean, changedControl: AbstractStruct, shouldHaveEmitted?: boolean) {
    const status = this._deriveStatus();

    if (emitEvent || shouldHaveEmitted) {
      this.statusReactive.set(status);
    } else {
      this.status = status;
    }

    if (this._parent) {
      this._parent._updateErrors(emitEvent, changedControl, shouldHaveEmitted);
    }
  }

  markAsPending(opts: { onlySelf?: boolean, source?: AbstractStruct }) {
    const control = opts.source ?? this;
    if (this._parent && !opts.onlySelf) {
      this._parent.markAsPending({ ...opts, source: control });
    }
  }

  clearValidator() {
    // eslint-disable-next-line no-multi-assign
    this._rawValidators = this._composedValidatorFn = null;
  }

  private updateStrategy: FormHooks = 'change';

  /** structure tree */

  abstract _forEachChild(cb: (c: AbstractStruct) => void): void;
  abstract _anyControls(fn:(c: AbstractStruct) => boolean):boolean;

  private _parent: AbstractStruct | AbstractStruct | null;

  setParent(p: AbstractStruct | null) {
    this._parent = p;
  }

  get parent() {
    return this._parent;
  }

  get root() {
    let x = this;
    while (x.parent) {
      // @ts-ignore
      x = x.parent;
    }
    return x;
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

  _find(name: string | number): AbstractStruct | null {
    return null;
  }

  abstract _allControlsDisabled(): boolean;

  /** dirty */
  private _dirtyReactive = signal(false);

  get dirty() {
    return untracked(this._dirtyReactive);
  }

  set dirty(val: boolean) {
    untracked(() => this._dirtyReactive.set(val));
  }

  markAsDirty(opts: { onlySelf?: boolean, source?: AbstractStruct }) {
    const changed = this.dirty === false;

    const sourceControl = opts.source ?? this;
    if (this._parent && !opts.onlySelf) {
      this._parent.markAsDirty({ ...opts, source: sourceControl });
    }

    if (changed) {
      this._dirtyReactive.set(true);
    }
  }

  markAsPristine(opts: { onlySelf?: boolean, source?: AbstractStruct }) {
    const changed = this.dirty === true;
    this.dirty = false;
    const control = opts.source ?? this;
    this._forEachChild((c) => {
      c.markAsPristine({ onlySelf: true });
    });
    if (this._parent && !opts.onlySelf) {
      this._parent._updateDirty(opts, control);
    }
    if (changed) {
      this._dirtyReactive.set(false);
    }
  }

  _anyControlsDirty() {
    return this._anyControls((c) => c.dirty);
  }

  _updateDirty(opts: { onlySelf?:boolean }, source:AbstractStruct) {
    const nextDirty = this._anyControlsDirty();
    const changed = this.dirty !== nextDirty;
    if (changed) {
      this._dirtyReactive.set(nextDirty);
    }
    if (this._parent && !opts.onlySelf) {
      this._parent._updateDirty(opts, source);
    }
  }

  /** values */
  readonly defaultValue: TValue | null;

  protected valueReactive = signal<TValue | undefined>(undefined);

  get value() {
    return untracked(() => this.valueReactive());
  }

  getRawValue(): any {
    return this.value;
  }
  abstract setValue(value: TRawValue, options?: Object): void;
  abstract patchValue(value: TValue, options?: Object): void;
  abstract reset(value?: TValue, options?: Object): void;
  abstract _deriveValue(): void;

  private readonly statusReactive = signal<FormControlStatus | undefined>(undefined);

  get enabled(): boolean {
    return this.status !== FormControlStatus.DISABLED;
  }

  get status(): FormControlStatus {
    return untracked(this.statusReactive);
  }

  private set status(v: FormControlStatus) {
    untracked(() => this.statusReactive.set(v));
  }

  private _deriveStatus():FormControlStatus {
    if (this._allControlsDisabled()) return FormControlStatus.DISABLED;
    if (this.errors) return FormControlStatus.INVALID;
    if (this._anyControlsHaveStatus(FormControlStatus.VALID)) return FormControlStatus.INVALID;
    return FormControlStatus.VALID;
  }

  updateValueAndValidity(opts: { onlySelf?: boolean, source?:AbstractStruct }) {
    this.status = this._allControlsDisabled() ? FormControlStatus.DISABLED : FormControlStatus.VALID;
    this._deriveValue();
    const source = opts.source ?? this;

    if (this._parent && !opts.onlySelf) {
      this._parent.updateValueAndValidity({ ...opts, source });
    }
  }

  enable() {
    this.status = FormControlStatus.VALID;
    this._forEachChild((c) => {
      c.enable();
    });
  }

  disable() { }

  private _isParentDirty(onlySelf = false) {
    const parentDirty = this._parent && this._parent.dirty;
    return onlySelf && !!parentDirty && this.parent._anyControlsDirty();
  }

  _anyControlsHaveStatus(status: FormControlStatus): boolean {
    return this._anyControls((c) => c.status === status);
  }

  private _updateParents() {

  }

  watchValue(cb: (v: TValue) => void): VoidFunction {
    const { destroy } = effect(() => {
      const value = this.valueReactive();
      cb(value);
    });
    return destroy;
  }
}
