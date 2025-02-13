import { AbstractStruct } from './model';

class FormPrimitive <TValue = any> extends AbstractStruct<TValue> {
  readonly defaultValue: TValue;

  override _allControlsDisabled(): boolean {
    return this.disabled;
  }
}
