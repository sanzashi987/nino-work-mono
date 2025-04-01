import { AbstractControl, ControlStatus } from './control';

class FormPrimitive <TValue = any> extends AbstractControl<TValue> {
  readonly initialValue: TValue;

  override _allControlsDisabled(): boolean {
    return this.status === ControlStatus.DISABLED;
  }

  override setValue(value: TValue, opts: object): void {
    this.valueReactive.set(value);
    this.updateValueAndValidity(opts);
  }

  override patchValue = this.setValue;

  override _deriveValue(): void { }

  override reset(opts: { onlySelf?: boolean } = {}): void {
    this.markAsPristine(opts);
    this.setValue(this.initialValue, opts);
  }

  override _forEachChild(): void {}

  override _anyControls(): boolean {
    return false;
  }
}

export default FormPrimitive;
