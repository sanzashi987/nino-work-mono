import { AbstractStruct } from './model';

class FormArray<TValue, TRawValue extends TValue = TValue> extends AbstractStruct<TValue, TRawValue> {
  controls: Array<AbstractStruct<any>> = [];

  get length(): number {
    return this.controls.length;
  }

  // at(index:number) { }
}
