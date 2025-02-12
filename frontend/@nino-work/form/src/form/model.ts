import { signal, untracked } from '../signal';

export const enum FormControlStatus {
  VALID = 'VALID',
  INVALID = 'INVALID',
  PENDING = 'PENDING',
  DISABLED = 'DISABLED'
}

export type FormHooks = 'change' | 'blur' | 'submit';

abstract class AbstractControl<TValue, TRawValue extends TValue = TValue> {
  public parent: FormGroup | FormArray | null;

  private valueReactive = signal<TValue | undefined>(undefined);

  get value() {
    return untracked(() => this.valueReactive());
  }

  private readonly statusReactive = signal<FormControlStatus | undefined>(undefined);

  get status(): FormControlStatus {
    return untracked(this.statusReactive)!;
  }

  private set status(v: FormControlStatus) {
    untracked(() => this.statusReactive.set(v));
  }

  get valid() {
    return this.status === FormControlStatus.VALID;
  }

  get invalid() {
    return this.status === FormControlStatus.INVALID;
  }

  get disabled() {
    return this.status === FormControlStatus.DISABLED;
  }

  get enabled() {
    return this.status !== FormControlStatus.DISABLED;
  }

  get root() {
    let x = this;
    while (x.parent) {
      x = x.parent;
    }
    return x;
  }
}
