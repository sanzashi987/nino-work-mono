import { effect, signal, untracked } from '../signal';
import type { ValidationErrors, ValidatorFn } from './validators';
import { composeValidators } from './validators';

export const enum ControlStatus {
  VALID = 'VALID',
  INVALID = 'INVALID',
  PENDING = 'PENDING',
  DISABLED = 'DISABLED'
}

export type IsAny<T, Y, N> = 0 extends 1 & T ? Y : N;
export type TypedOrUntyped<T, Typed, Untyped> = IsAny<T, Untyped, Typed>;
export type FormValue<T extends AbstractControl | undefined> =
  T extends AbstractControl<any, any> ? T['value'] : never;
export type FormRawValue<T extends AbstractControl | undefined> =
  T extends AbstractControl<any, any>
    ? T['setValue'] extends (v: infer R) => void
      ? R
      : never
    : never;

export abstract class AbstractControl<TValue = any, TRawValue extends TValue = TValue> {
  name: string;

  /** validators */
  public errors: ValidationErrors | null = null;

  private _composedValidatorFn: ValidatorFn | null;

  private _rawValidators: ValidatorFn | ValidatorFn[] | null;

  private _composeValidators(validators: ValidatorFn | ValidatorFn[] | null) {
    const validatorArr = Array.isArray(validators) ? validators.slice() : validators;
    this._rawValidators = validatorArr;
    this._composedValidatorFn = Array.isArray(validatorArr) ? composeValidators(validatorArr) : validatorArr || null;
  }

  private runValidator(emitEvent?: boolean) {
    if (this._composedValidatorFn) {
      this.statusReactive.set(ControlStatus.PENDING);
      Promise.resolve().then(() => this._composedValidatorFn(this))
        .then((err) => {
          this.errors = err;
          this._updateErrors(emitEvent, this);
        })
        .catch(() => {
          this.statusReactive.set(ControlStatus.VALID);
        });
    }
  }

  private _updateErrors(emitEvent: boolean, changedControl: AbstractControl) {
    const status = this._deriveStatus();

    if (emitEvent) {
      this.statusReactive.set(status);
    } else {
      this.status = status;
    }

    if (this._parent) {
      this._parent._updateErrors(emitEvent, changedControl);
    }
  }

  markAsPending(opts: { onlySelf?: boolean, source?: AbstractControl }) {
    this.statusReactive.set(ControlStatus.PENDING);
    const control = opts.source ?? this;
    if (this._parent && !opts.onlySelf) {
      this._parent.markAsPending({ ...opts, source: control });
    }
  }

  clearValidator() {
    // eslint-disable-next-line no-multi-assign
    this._rawValidators = this._composedValidatorFn = null;
  }

  /** structure tree */

  abstract _forEachChild(cb: (c: AbstractControl) => void): void;
  abstract _anyControls(fn:(c: AbstractControl) => boolean):boolean;

  private _parent: AbstractControl | AbstractControl | null;

  setParent(p: AbstractControl | null) {
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

  get<P extends string | readonly (string | number)[]>(path: P): AbstractControl<any> | null;
  get<P extends string | (string | number)[]>(path: P): AbstractControl<any> | null {
    let currPath: Array<string | number> | string = path;
    if (currPath == null) return null;
    if (!Array.isArray(currPath)) currPath = currPath.split('.');
    if (currPath.length === 0) return null;
    return currPath.reduce(
      (control: AbstractControl | null, name) => control && control._find(name),
      this
    );
  }

  _find(name: string | number): AbstractControl | null {
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

  markAsDirty(opts: { onlySelf?: boolean, source?: AbstractControl }) {
    const changed = this.dirty === false;

    const sourceControl = opts.source ?? this;
    if (this._parent && !opts.onlySelf) {
      this._parent.markAsDirty({ ...opts, source: sourceControl });
    }

    if (changed) {
      this._dirtyReactive.set(true);
    }
  }

  markAsPristine(opts: { onlySelf?: boolean, source?: AbstractControl }) {
    const changed = this.dirty === true;

    const control = opts.source ?? this;
    this._forEachChild((c) => {
      c.markAsPristine({ onlySelf: true });
    });
    if (changed) {
      this._dirtyReactive.set(false);
    }
    if (this._parent && !opts.onlySelf) {
      this._parent._updateDirty(opts, control);
    }
  }

  _anyControlsDirty() {
    return this._anyControls((c) => c.dirty);
  }

  _updateDirty(opts: { onlySelf?:boolean }, source:AbstractControl) {
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

  private readonly statusReactive = signal<ControlStatus | undefined>(undefined);

  get enabled(): boolean {
    return this.status !== ControlStatus.DISABLED;
  }

  get status(): ControlStatus {
    return untracked(this.statusReactive);
  }

  private set status(v: ControlStatus) {
    untracked(() => this.statusReactive.set(v));
  }

  private _deriveStatus():ControlStatus {
    if (this._allControlsDisabled()) return ControlStatus.DISABLED;
    if (this.errors) return ControlStatus.INVALID;
    if (this._anyControlsHaveStatus(ControlStatus.VALID)) return ControlStatus.INVALID;
    return ControlStatus.VALID;
  }

  updateValueAndValidity(opts: { onlySelf?: boolean, source?:AbstractControl }) {
    const status = this._allControlsDisabled() ? ControlStatus.DISABLED : ControlStatus.VALID;
    this._deriveValue();
    const source = opts.source ?? this;

    if (status !== ControlStatus.DISABLED) {
      this.runValidator();
    }
    if (this._parent && !opts.onlySelf) {
      this._parent.updateValueAndValidity({ ...opts, source });
    }
  }

  enable(opts: { onlySelf?: boolean }) {
    this.statusReactive.set(ControlStatus.VALID);
    const dirty = this._isParentDirty();
    this._forEachChild((c) => {
      c.enable({ ...opts, onlySelf: true });
    });
    this.updateValueAndValidity({ onlySelf: true });
    this._updateParents({ ...opts, skipPristineCheck: dirty }, this);
  }

  disable(opts: { onlySelf?: boolean }) {
    this.statusReactive.set(ControlStatus.DISABLED);
    const dirty = this._isParentDirty();
    this._forEachChild((c) => {
      c.disable({ ...opts, onlySelf: true });
    });
    this.updateValueAndValidity({ onlySelf: true });
    this._updateParents({ ...opts, skipPristineCheck: dirty }, this);
  }

  private _isParentDirty(onlySelf = false) {
    const parentDirty = this._parent && this._parent.dirty;
    return onlySelf && !!parentDirty && this.parent._anyControlsDirty();
  }

  _anyControlsHaveStatus(status: ControlStatus): boolean {
    return this._anyControls((c) => c.status === status);
  }

  private _updateParents(
    opts: { onlySelf?: boolean; skipPristineCheck?: boolean },
    source: AbstractControl
  ) {
    if (this._parent) {
      this._parent.updateValueAndValidity({});
      if (!opts.skipPristineCheck) {
        this._parent._updateDirty({}, source);
      }
    }
  }

  watchValue(cb: (v: TValue) => void): VoidFunction {
    const { destroy } = effect(() => {
      const value = this.valueReactive();
      cb(value);
    });
    return destroy;
  }
}
