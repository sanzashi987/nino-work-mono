import { AbstractStruct, FormRawValue, FormValue, TypedOrUntyped } from './model';

export type ExtractFormArrayValue<T extends AbstractStruct<any>> = TypedOrUntyped<
T,
Array<FormValue<T>>,
any[]
>;
export type ExtractFormArrayRawValue<T extends AbstractStruct<any>> = TypedOrUntyped<
T,
Array<FormRawValue<T>>,
any[]
>;

class FormArray<TControl extends AbstractStruct<any> = any> extends AbstractStruct<TValue, TRawValue> {
  controls: Array<AbstractStruct<any>> = [];

  get length(): number {
    return this.controls.length;
  }

  private _adjustIndex(index: number): number {
    return index < 0 ? index + this.length : index;
  }

  at(index: number): TypedOrUntyped<TControl, TControl, AbstractStruct<any>> {
    return (this.controls as any)[this._adjustIndex(index)];
  }

  override _find(name: string | number): AbstractStruct | null {
    return this.controls[name] ?? null;
  }
}
