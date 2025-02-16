import {
  AbstractControl, FormRawValue, FormValue, IsAny, TypedOrUntyped
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
  _forEachChild(cb: (c: AbstractControl) => void): void {
    throw new Error('Method not implemented.');
  }

  _anyControls(fn: (c: AbstractControl) => boolean): boolean {
    throw new Error('Method not implemented.');
  }

  _allControlsDisabled(): boolean {
    throw new Error('Method not implemented.');
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
    this.controls;
  }
}
