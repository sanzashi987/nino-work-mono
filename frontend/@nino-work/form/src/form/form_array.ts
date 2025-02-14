import { AbstractControl, FormRawValue, FormValue, TypedOrUntyped } from './model';

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

class FormArray<TControl extends AbstractControl<any> = any> extends AbstractControl<TValue, TRawValue> {
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
}
