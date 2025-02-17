import { AbstractControl, ControlStatus } from './model';

class FormPrimitive <TValue = any> extends AbstractControl<TValue> {
  readonly defaultValue: TValue;

  override _allControlsDisabled(): boolean {
    return this.status === ControlStatus.DISABLED;
  }

  override setValue(value: TValue, opts: Object): void {
    this.valueReactive.set(value);
    this.updateValueAndValidity(opts);
  }

  override patchValue = this.setValue;

  override _deriveValue(): void { }

  override reset(value: TValue = this.defaultValue, opts?: Object): void {
    this.markAsPristine(opts);
    this.setValue(value, opts);
  }

  override _forEachChild(): void {}

  override _anyControls(): boolean {
    return false;
  }
}

export default FormPrimitive;
