/* eslint-disable no-restricted-syntax */
import {
  AbstractControl, ControlStatus, FormRawValue, FormValue, IsAny, TypedOrUntyped
} from './model';

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
  override _forEachChild(cb: (c: AbstractControl, index:number) => void): void {
    this.controls.forEach((control: AbstractControl, index: number) => {
      cb(control, index);
    });
  }

  _anyControls(fn: (c: AbstractControl) => boolean): boolean {
    return this.controls.some((control) => control.enabled && fn(control));
  }

  _allControlsDisabled(): boolean {
    for (const control of this.controls) {
      if (control.enabled) return false;
    }
    return this.controls.length > 0 || this.status === ControlStatus.DISABLED;
  }

  setValue(value: TypedOrUntyped<TControl, IsAny<TControl, any[], FormRawValue<TControl>[]>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  patchValue(value: TypedOrUntyped<TControl, IsAny<TControl, any[], FormValue<TControl>[]>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  reset(value?: TypedOrUntyped<TControl, IsAny<TControl, any[], FormValue<TControl>[]>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
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
}
