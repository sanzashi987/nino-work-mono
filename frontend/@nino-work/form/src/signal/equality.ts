export type ValueEqualityFn<T> = (a: T, b: T) => boolean;

/**
 * The default equality function used for `signal` and `computed`, which uses referential equality.
 */
export function defaultEquals<T>(a: T, b: T) {
  return Object.is(a, b);
}
