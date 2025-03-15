import type {
  GetVariableType,
  SharedPreferenceVariableType,
  StaticVariableType,
  UrlVariableType
} from '@/types';

export interface TypedVariableMutatorCtor<T> {
  type: T;
  new (): TypedVariableMutator<T>;
}

export interface TypedVariableMutator<T> {
  getValue(conf: GetVariableType<T>): any;
  setValue(conf: GetVariableType<T>, value: any): void;
}

/** 临时变量 */
class StaticVariableMutator implements TypedVariableMutator<'static'> {
  static type = 'static' as const;

  cache = new Map<string, any>();

  getValue(conf: StaticVariableType) {
    const { id, default: defaultVal } = conf;
    if (!this.cache.has(id)) {
      this.cache.set(id, defaultVal);
      // this.cache.set(id, JSON.parse(defaultVal));
    }
    return this.cache.get(id);
  }

  setValue(conf: StaticVariableType, value: any): void {
    const { id } = conf;
    this.cache.set(id, value);
  }
}

/** url查询参数 */
class UrlVariableMutator implements TypedVariableMutator<'url'> {
  static type = 'url' as const;

  cache = new Map<string, any>(new URLSearchParams(window.location.search));

  getValue(conf: UrlVariableType) {
    const { default: defaultVal, detail } = conf;
    if (!this.cache.has(detail.key)) {
      // this.cache.set(detail.key, JSON.parse(defaultVal));
      this.cache.set(detail.key, defaultVal);
    }
    return this.cache.get(detail.key);
  }

  setValue(): void {
    throw new Error('Url Variable shall not be set.');
  }
}

/** 持久化变量 */
class SharedVariableMutator implements TypedVariableMutator<'shared'> {
  static type = 'shared' as const;

  storageKey = 'canvas_shared_variable';

  // TODO 多端适配 SharedPreference
  getValue(conf: SharedPreferenceVariableType) {
    const { id, default: defaultVal } = conf;
    const config = JSON.parse(localStorage.getItem(this.storageKey) || '{}');
    if (Object.hasOwnProperty.call(config, id)) {
      return config[id];
    }
    this.setValue(conf, defaultVal);
    return defaultVal;
    // throw new Error('Method not implemented.');
  }

  setValue(conf: SharedPreferenceVariableType, value: any): void {
    const config = JSON.parse(localStorage.getItem(this.storageKey) || '{}');
    localStorage.setItem(
      this.storageKey,
      JSON.stringify({
        ...config,
        [conf.id]: value
      })
    );
    // throw new Error('Method not implemented.');
  }
}

export { StaticVariableMutator, SharedVariableMutator, UrlVariableMutator };
