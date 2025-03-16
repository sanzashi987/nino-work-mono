/* eslint-disable no-restricted-syntax */
import { decideControl, type IModel, type ObjectModel } from './define';
import {
  AbstractControl, ControlStatus, FormRawValue, FormValue, IsAny, TypedOrUntyped
} from './control';
import { Path } from './validators';

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

class FormObject<TValue extends object = any, TRawValue extends TValue = TValue>
  extends AbstractControl<TValue, TRawValue> {
  constructor(model: ObjectModel<TValue, any>, initialValue: any = {}) {
    super(model, initialValue);
    const myInitialValue = {};
    const otherValue = { ...initialValue };
    if (Array.isArray(model.children)) {
      this.controls = Object.fromEntries((model.children as IModel<any, any>[]).map((ctrl) => {
        const { field } = ctrl;
        if (field in initialValue) {
          myInitialValue[field] = initialValue[field];
          delete otherValue[field];
        }
        return [field, decideControl(ctrl, initialValue[field])];
      })) as any;
    }
    // @ts-ignore
    this.initialValue = myInitialValue;
    this.outsideValues = otherValue;
  }

  contains<K extends string & keyof TValue>(controlName: K): boolean {
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
    for (const controlName of Object.keys(this.controls) as Array<keyof TValue>) {
      if ((this.controls as any)[controlName].status !== ControlStatus.DISABLED) {
        return false;
      }
    }
    return Object.keys(this.controls).length > 0 || this.status === ControlStatus.DISABLED;
  }

  outsideValues: any = {};

  controls: { [K in keyof TValue]: AbstractControl<TValue[K]> };

  override setValue(value: Partial<TRawValue>, options?: object): void {
    if (!(value instanceof Array) && typeof value === 'object') {
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
  }

  override patchValue(value: Partial<TRawValue>, options?: object): void {
    if (value == null) return;
    (Object.keys(value) as Array<keyof TValue>).forEach((name) => {
      const control = (this.controls as any)[name];
      if (control) {
        control.patchValue(value[name as any]);
      }
    });
    this.updateValueAndValidity(options);
  }

  override reset(options: { onlySelf?: boolean } = {}): void {
    this._forEachChild((control: AbstractControl) => {
      control.reset({ onlySelf: true });
    });
    this._updateDirty(options, this);
    this.outsideValues = {};
    this.updateValueAndValidity(options);
  }

  _deriveValue(): void {
    this.valueReactive.set(this._reduceValue() as any);
  }

  _reduceValue(): Partial<TValue> {
    const acc: Partial<TValue> = {};
    return this._reduceChildren(acc, (last, control, name) => {
      last[name] = control.value;
      return last;
    });
  }

  /** @internal */
  _reduceChildren<T, K extends keyof TValue>(
    initValue: T,
    fn: (acc: T, control: AbstractControl<TValue[K], any>, name: K) => T
  ): T {
    let res = initValue;
    this._forEachChild((control: AbstractControl<TValue[K], any>, name: K) => {
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

  override getRawValue(): TRawValue {
    const next: any = { ...this.outsideValues };
    this._forEachChild((ctrl, key) => {
      next[key] = ctrl.getRawValue();
    });
    return next;
  }

  override _findChildName(control: AbstractControl): Path | null {
    const res = Object.entries(this.controls).find((c) => c[1] === control);
    if (!res) return null;
    return res[0];
  }
}

export default FormObject;
