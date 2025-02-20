/* eslint-disable no-restricted-syntax */
import { type ArrayModel, decideControl } from './define';
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
  constructor(model:ArrayModel<ExtractFormArrayValue<TControl>, any>, initialValue:ExtractFormArrayValue<TControl> = []) {
    super(model as any, initialValue);
    const arrayEnsured = Array.isArray(initialValue) ? initialValue : [];
    if (typeof model.children === 'object' && !Array.isArray(model.children)) {
      this.controls = arrayEnsured.map((v) => decideControl(model.children as any, v));
      this._forEachChild((control) => {
        control.setParent(this);
      });
    }
    // @ts-ignore
    this.initialValue = arrayEnsured;
    this.updateValueAndValidity({ onlySelf: true });
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

  private createChildren(value?:any) {
    const model = this.protoModel;
    if ('children' in model && typeof model.children === 'object' && !Array.isArray(model.children)) {
      const control = decideControl(model, value);
      control.setParent(this);
      return control;
    }
    return null;
  }

  push(value?: any) {
    const control = this.createChildren(value);
    if (control) {
      this.controls.push(control);
      this.updateValueAndValidity({});
    }
  }

  insert(index: number, value?: any) {
    const control = this.createChildren(value);
    if (control) {
      this.controls.splice(index, 0, control);
      this.updateValueAndValidity({});
    }
  }

  removeAt(index: number) {
    const adjustedIndex = this._adjustIndex(index);
    if (adjustedIndex >= 0) {
      const controls = this.controls.splice(adjustedIndex, 1);
      controls.forEach((control) => {
        control.setParent(null);
      });
      this.updateValueAndValidity({});
    }
  }
}

export default FormArray;
