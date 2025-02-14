import { AbstractStruct, FormControlStatus } from './model';

class FormPrimitive <TValue = any> extends AbstractStruct<TValue> {
  readonly defaultValue: TValue;

  override _allControlsDisabled(): boolean {
    return this.status === FormControlStatus.DISABLED;
  }

  override setValue(value: TValue, opts: Object): void {
    this.valueReactive.set(value);
    this.updateValueAndValidity(opts);
  }

  override patchValue = this.setValue;

  override reset(value: TValue = this.defaultValue, opts?: Object): void {
    this.setValue(value, opts);
  }
}
