import type { AbstractStruct } from './model';

export abstract class ControlEvent<T = any> {
  public abstract readonly source: AbstractStruct<unknown>;
}

export class ValueChangeEvent<T> extends ControlEvent<T> {
  constructor(
    public readonly value: T,
    public readonly source: AbstractStruct
  ) {
    super();
  }
}

export class DityhangeEvent extends ControlEvent {
  constructor(
    public readonly dirty: boolean,
    public readonly source: AbstractStruct
  ) {
    super();
  }
}
