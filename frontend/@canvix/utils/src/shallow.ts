function is(x: any, y: any) {
  return (
    (x === y && (x !== 0 || 1 / x === 1 / y)) || (x !== x && y !== y) // eslint-disable-line no-self-compare
  );
}

const objectIs = typeof Object.is === 'function' ? Object.is : is;

const hasOwnProperty$2 = Object.prototype.hasOwnProperty;

export function shallowEqual(objA: any, objB: any, depth = 0) {
  if (objectIs(objA, objB)) {
    return true;
  }

  if (typeof objA !== 'object' || objA === null || typeof objB !== 'object' || objB === null) {
    return false;
  }

  const keysA = Object.keys(objA);
  const keysB = Object.keys(objB);

  if (keysA.length !== keysB.length) {
    // console.log('length');
    return false;
  } // Test for A's keys different from B.

  for (let i = 0; i < keysA.length; i++) {
    if (depth) {
      if (
        !hasOwnProperty$2.call(objB, keysA[i])
        || !shallowEqual(objA[keysA[i]], objB[keysA[i]], depth - 1)
      ) {
        return false;
      }
    } else if (!hasOwnProperty$2.call(objB, keysA[i]) || !objectIs(objA[keysA[i]], objB[keysA[i]])) {
      return false;
    }
  }

  return true;
}

export const shallowClone = (data: any) => {
  if (data && typeof data === 'object') {
    if (Array.isArray(data)) {
      return [...data];
    }
    return { ...data };
  }
  return data;
};

export function strictEquality<T>(a: T, b: T) {
  return a === b;
}
