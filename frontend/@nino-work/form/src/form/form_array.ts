/* eslint-disable no-restricted-syntax */
import { ArrayModel } from './define';
import {
  AbstractControl, ControlStatus, FormRawValue, FormValue, IsAny, TypedOrUntyped
} from './control';
import { Path } from './validators';

export type ExtractFormArrayValue<T extends AbstractControl<any>> = TypedOrUntyped<
T,
Array<FormValue<T>>,
any[]
>;
export type ExtractFormArrayRawValue<T extends AbstractControl<any>> = TypedOrUntyped<
T,
Array<FormRawValue<T>>,
any[]
>;

class FormArray<TControl extends AbstractControl<any> = any> extends AbstractControl<
TypedOrUntyped<TControl, ExtractFormArrayValue<TControl>, any>,
TypedOrUntyped<TControl, ExtractFormArrayRawValue<TControl>, any>
> {
  constructor(model:ArrayModel<ExtractFormArrayValue<TControl>, any>, initialValue?:ExtractFormArrayValue<TControl>) {
    super(model, initialValue);
  }

  override _forEachChild(cb: (c: AbstractControl, index:number) => void): void {
    this.controls.forEach((control: AbstractControl, index: number) => {
      cb(control, index);
    });
  }

  override _anyControls(fn: (c: AbstractControl) => boolean): boolean {
    return this.controls.some((control) => control.enabled && fn(control));
  }

  override _allControlsDisabled(): boolean {
    for (const control of this.controls) {
      if (control.enabled) return false;
    }
    return this.controls.length > 0 || this.status === ControlStatus.DISABLED;
  }

  override setValue(value: TypedOrUntyped<TControl, IsAny<TControl, any[], FormRawValue<TControl>[]>, any>, options: { onlySelf?: boolean } = {}): void {
    if (!Array.isArray(value)) {
      console.warn('expected to receive an array as value');
      return;
    }
    value.forEach((val, index) => {
      const control = this.at(index);
      if (control) {
        control.setValue(val, { onlySelf: true });
      }
    });
    this.updateValueAndValidity(options);
  }

  override patchValue(value: TypedOrUntyped<TControl, IsAny<TControl, any[], FormValue<TControl>[]>, any>, options?: Object): void {
    if (!Array.isArray(value)) {
      console.warn('expected to receive an array as value');
      return;
    }
    value.forEach((val, index) => {
      const control = this.at(index);
      if (control) {
        control.setValue(val, { onlySelf: true });
      }
    });

    this.updateValueAndValidity(options);
  }

  override reset(options: { onlySelf?: boolean } = {}): void {
    this.controls.forEach((c) => {
      c.reset({ onlySelf: true });
    });

    this._updateDirty(options, this);
    this.updateValueAndValidity(options);
  }

  override getRawValue() :any {
    return this.controls.map((c) => c.getRawValue());
  }

  controls: Array<AbstractControl<any>> = [];

  get length(): number {
    return this.controls.length;
  }

  private _adjustIndex(index: number): number {
    return index < 0 ? index + this.length : index;
  }

  at(index: number): TypedOrUntyped<TControl, TControl, AbstractControl<any>> {
    return (this.controls as any)[this._adjustIndex(index)];
  }

  override _find(name: string | number): AbstractControl | null {
    return this.controls[name] ?? null;
  }

  override _deriveValue(): void {
    this.valueReactive.set(
      this.controls.map((control) => control.value)
    );
  }

  override _findChildName(control: AbstractControl): Path | null {
    const index = this.controls.findIndex((c) => c === control);
    if (index === -1) return null;
    return index;
  }

  push() { }

  remove() { }
}

export default FormArray;
