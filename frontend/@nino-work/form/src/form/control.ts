import { effect, signal, untracked } from '../signal';
import { BaseModel } from './define';
import type FormArray from './form_array';
import type FormObject from './form_object';
import type { ComposedValidatorFn, Path, ValidationError, ValidatorRule } from './validators';
import { composeValidator } from './validators';

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

export type LegalParent = FormObject | FormArray;
let nextInstanceId = 0;
export abstract class AbstractControl<TValue = any, TRawValue extends TValue = TValue> {
  protected readonly instanceId: string;

  protected readonly protoModel: BaseModel<any, any>;

  constructor(model: BaseModel<TValue, any>, initialValue?: TValue) {
    this.initialValue = initialValue ?? model.formItemProps.initialValue;
    this.protoModel = model;
    // eslint-disable-next-line no-plusplus
    this.instanceId = `nino-control-${nextInstanceId++}`;
    if (model.formItemProps?.rules) {
      this._composeValidator(model.formItemProps.rules);
    }
  }

  /** validators */
  public errors: ValidationError | null = null;

  private _composedValidatorFn: ComposedValidatorFn | null;

  // private _rawRules: ValidatorRule[] | null;

  private _composeValidator(rules?: ValidatorRule[]) {
    if (rules) {
      // this._rawRules = rules;
      this._composedValidatorFn = composeValidator(rules);
    }
  }

  private runValidator() {
    if (this._composedValidatorFn) {
      this.statusReactive.set(ControlStatus.PENDING);
      this._composedValidatorFn(this.value).then(({ errors, warnings }) => {
        this.errors = (errors.length === 0 && warnings.length === 0) ? null : { errors, warnings };
        this._updateErrors(this);
      });
    }
  }

  private _updateErrors(changedControl: AbstractControl) {
    const status = this._deriveStatus();
    this.statusReactive.set(status);
    if (this._parent) {
      this._parent._updateErrors(changedControl);
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
    this._composedValidatorFn = null;
  }

  /** structure tree */

  abstract _forEachChild(cb: (c: AbstractControl) => void): void;
  abstract _anyControls(fn:(c: AbstractControl) => boolean):boolean;
  private _parent: LegalParent | null;

  setParent(p: LegalParent | null) {
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

  get<P extends string | readonly Path[]>(path: P): AbstractControl<any> | null;
  get<P extends string | Path[]>(path: P): AbstractControl<any> | null {
    let currPath: Array<Path> | string = path;
    if (currPath == null) return null;
    if (!Array.isArray(currPath)) currPath = currPath.split('.');
    if (currPath.length === 0) return null;
    return currPath.reduce(
      (control: AbstractControl | null, name) => control && control._find(name),
      this
    );
  }

  _find(name: Path): AbstractControl | null {
    return null;
  }

  _findChildName(control: AbstractControl): Path | null {
    return null;
  }

  /* the name always stores at its parent, root will be an empty arr */
  get name(): Path[] {
    if (this._parent) {
      const parentBase = this._parent.name;
      const localField = this._parent._findChildName(this);
      if (localField === null) {
        // unexpected case ,
        // a control that isnt registered to the parent
        return [];
      }
      return parentBase.concat(localField);
    }
    return [];
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

  _updateDirty(opts: { onlySelf?: boolean }, source: AbstractControl) {
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
  readonly initialValue: TValue | null;

  protected valueReactive = signal<TValue | undefined>(undefined);

  get value() {
    return untracked(() => this.valueReactive());
  }

  getRawValue(): any {
    return this.value;
  }
  abstract setValue(value: TRawValue, options?: Object): void;
  abstract patchValue(value: TValue, options?: Object): void;
  abstract reset(options?: Object): void;
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
    if (this._anyControlsHaveStatus(ControlStatus.PENDING)) return ControlStatus.PENDING;
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
    if (this._parent && !opts.onlySelf) {
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
