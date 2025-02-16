export function mergeDefault<T extends {}>(value: Partial<T> | undefined, defaultVal: Partial<T>) {
  const next = value ?? {};
  return Object.keys(defaultVal).reduce((last, key, val) => {
    if (last[key] === undefined && val[key] !== undefined) {
      // eslint-disable-next-line no-param-reassign
      last[key] = val[key];
    }
    return last;
  }, next);
}

export function noop() {}
