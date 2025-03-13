export class ValueRW {
  protected value: Record<string, any> = {};

  setValue = (payload: Record<string, any>) => {
    Object.entries(payload).forEach(([key, val]) => {
      if (val === undefined) {
        delete this.value[key];
      } else {
        this.value[key] = val;
      }
    });
  };

  getValue = () => this.value;
  // TODO add value change event broadcast
}

const ReadOnlyHandler = {
  get(target: any, propKey: any) {
    return target[propKey];
  },
  set() {
    return false;
  }
};

export class SysValueRw extends ValueRW {
  constructor(init: Record<string, any>) {
    super();
    // const obj: Record<string, any> = Object.fromEntries(
    // new URLSearchParams(window.location.search).entries(),
    // );
    // obj['dev'] = dev;
    const obj = JSON.parse(JSON.stringify({ ...init }));

    this.value = new Proxy(obj, ReadOnlyHandler);

    // eslint-disable-next-line no-constructor-return
    return new Proxy(this, ReadOnlyHandler);
  }

  override setValue = () => {};
}

export type ValueContext = {
  getValue: () => Record<string, any>;
  setValue: (payload: Record<string, any>) => void;
};

export const VoidContext = {
  setValue: () => {},
  getValue: () => ({})
};
