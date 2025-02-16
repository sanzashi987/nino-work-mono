import { ObjectModel } from './define';
import {
  AbstractControl, FormRawValue, FormValue, IsAny, TypedOrUntyped
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
  _anyControls(fn: (c: AbstractControl) => boolean): boolean {
    throw new Error('Method not implemented.');
  }

  _allControlsDisabled(): boolean {
    throw new Error('Method not implemented.');
  }

  _deriveValue(): void {
    throw new Error('Method not implemented.');
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
  }

  patchValue(value: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, Partial<{ [K in keyof TControl]: FormValue<TControl[K]>; }>>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  reset(value?: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, Partial<{ [K in keyof TControl]: FormValue<TControl[K]>; }>>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  constructor(controls: TControl, private proto: ObjectModel<TControl, any>) {
    super();
    this.controls = controls;
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
    const next: any = {};
    this._forEachChild((ctrl, key) => {
      next[key] = ctrl.getRawValue();
    });
    return next;
  }
}
