/* eslint-disable no-restricted-syntax */
import { ObjectModel } from './define';
import {
  AbstractControl, ControlStatus, FormRawValue, FormValue, IsAny, TypedOrUntyped
} from './model';

export type ExtractFormObjectValue<T extends { [K in keyof T]?: AbstractControl<any> }> = TypedOrUntyped<
T,
Partial<{ [K in keyof T]: FormValue<T[K]> }>,
{ [key: string]: any }
>;

export type ExtractFormObjectRawValue<T extends { [K in keyof T]?: AbstractControl<any> }> = TypedOrUntyped<
T,
{ [K in keyof T]: FormRawValue<T[K]> },
{ [key: string]: any }
>;

export class FormObject<TControl extends { [K in keyof TControl]: AbstractControl<any> } = any>
  extends AbstractControl<
  TypedOrUntyped<TControl, ExtractFormObjectValue<TControl>, any>,
  TypedOrUntyped<TControl, ExtractFormObjectRawValue<TControl>, any>
  > {
  constructor(controls: TControl, private proto: ObjectModel<TControl, any>) {
    super();
    this.controls = controls;
  }

  contains<K extends string & keyof TControl>(controlName: K): boolean {
    // eslint-disable-next-line no-prototype-builtins
    return this.controls.hasOwnProperty(controlName);
  }

  override _anyControls(fn: (c: AbstractControl) => boolean): boolean {
    for (const [controlName, control] of Object.entries(this.controls)) {
      if (this.contains(controlName as any) && fn(control as any)) {
        return true;
      }
    }
    return false;
  }

  override _allControlsDisabled(): boolean {
    for (const controlName of Object.keys(this.controls) as Array<keyof TControl>) {
      if ((this.controls as any)[controlName].status !== ControlStatus.DISABLED) {
        return false;
      }
    }
    return Object.keys(this.controls).length > 0 || this.status === ControlStatus.DISABLED;
  }

  outsideValues: any = {};

  defaultValue: TypedOrUntyped<TControl, ExtractFormObjectRawValue<TControl>, any>;

  controls: TypedOrUntyped<TControl, TControl, { [key: string]: AbstractControl<any> }>;

  override setValue(value: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, { [K in keyof TControl]: FormRawValue<TControl[K]>; }>, any>, options?: Object): void {
    Object.keys(value).forEach((name) => {
      const control = this.controls[name];
      if (control) {
        control.setValue(value[name]);
      } else {
        this.outsideValues[name] = value[name];
      }
    });
    this.updateValueAndValidity(options);
  }

  override patchValue(value: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, Partial<{ [K in keyof TControl]: FormValue<TControl[K]>; }>>, any>, options?: Object): void {
    if (value == null) return;
    (Object.keys(value) as Array<keyof TControl>).forEach((name) => {
      const control = (this.controls as any)[name];
      if (control) {
        control.patchValue(value[name as any]);
      }
    });
    this.updateValueAndValidity(options);
  }

  override reset(value?: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, Partial<{ [K in keyof TControl]: FormValue<TControl[K]>; }>>, any>, options?: Object): void {
    this._forEachChild((control: AbstractControl, name) => {
      control.reset(value ? (value as any)[name] : null);
    });
    this._updateDirty(options, this);
    this.outsideValues = {};
    this.updateValueAndValidity(options);
  }

  _deriveValue(): void {
    this.valueReactive.set(this._reduceValue() as any);
  }

  _reduceValue(): Partial<TControl> {
    const acc: Partial<TControl> = {};
    return this._reduceChildren(acc, (last, control, name) => {
      // eslint-disable-next-line no-param-reassign
      last[name] = control.value;
      return last;
    });
  }

  /** @internal */
  _reduceChildren<T, K extends keyof TControl>(
    initValue: T,
    fn: (acc: T, control: TControl[K], name: K) => T
  ): T {
    let res = initValue;
    this._forEachChild((control: TControl[K], name: K) => {
      res = fn(res, control, name);
    });
    return res;
  }

  setupControls(): void {
    this._forEachChild((c) => {
      c.setParent(this);
    });
  }

  override _forEachChild(cb: (c: any, key:any) => void): void {
    Object.keys(this.controls).forEach((key) => {
      const control = (this.controls as any)[key];
      if (control) {
        cb(control, key);
      }
    });
  }

  override _find(name: string | number): AbstractControl | null {
    return this.controls[name] ?? null;
  }

  override getRawValue():TypedOrUntyped<TControl, ExtractFormObjectRawValue<TControl>, any> {
    const next: any = { ...this.outsideValues };
    this._forEachChild((ctrl, key) => {
      next[key] = ctrl.getRawValue();
    });
    return next;
  }
}
