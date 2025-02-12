import { ObjectModel } from './define';
import {
  AbstractStruct, FormRawValue, FormValue, IsAny, TypedOrUntyped
} from './model';

export type ExtractFormObjectValue<T extends { [K in keyof T]?: AbstractStruct<any> }> = TypedOrUntyped<
T,
Partial<{ [K in keyof T]: FormValue<T[K]> }>,
{ [key: string]: any }
>;

export type ExtractFormObjectRawValue<T extends { [K in keyof T]?: AbstractStruct<any> }> = TypedOrUntyped<
T,
{ [K in keyof T]: FormRawValue<T[K]> },
{ [key: string]: any }
>;

export class FormObject<TControl extends { [K in keyof TControl]: AbstractStruct<any> } = any>
  extends AbstractStruct<
  TypedOrUntyped<TControl, ExtractFormObjectValue<TControl>, any>,
  TypedOrUntyped<TControl, ExtractFormObjectRawValue<TControl>, any>
  > {
  setValue(value: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, { [K in keyof TControl]: FormRawValue<TControl[K]>; }>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  patchValue(value: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, Partial<{ [K in keyof TControl]: FormValue<TControl[K]>; }>>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  reset(value?: TypedOrUntyped<TControl, IsAny<TControl, { [key: string]: any; }, Partial<{ [K in keyof TControl]: FormValue<TControl[K]>; }>>, any>, options?: Object): void {
    throw new Error('Method not implemented.');
  }

  controls: TypedOrUntyped<TControl, TControl, { [key: string]: AbstractStruct<any> }>;

  constructor(controls: TControl, private proto: ObjectModel<TControl, any>) {
    super();
    this.controls = controls;
  }

  setupControls(): void {
    this.iterChild((c) => {
      c.setParent(c);
    });
  }

  override iterChild(cb: (c: AbstractStruct) => void): void {
    Object.keys(this.controls).forEach((key) => {
      const control = (this.controls as any)[key];
      if (control) {
        cb(control);
      }
    });
  }
}
