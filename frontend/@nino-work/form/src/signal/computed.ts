/* eslint-disable no-plusplus */
/* eslint-disable no-param-reassign */
import { defaultEquals, ValueEqualityFn } from './equality';
import {
  accessProducer,
  consumerBeginWork, consumerCompleteWork, REACTIVE_NODE, ReactiveNode, setActiveConsumer, Signal, SIGNAL,
  updateProducerValueVersion
} from './reactive';

export interface ComputedNode<T> extends ReactiveNode {
  value: T;
  error: unknown;

  computation: () => T;

  equal: ValueEqualityFn<T>;
}

export type ComputedGetter<T> = (() => T) & {
  [SIGNAL]: ComputedNode<T>;
};

export const UNSET: any = /* @__PURE__ */ Symbol('UNSET');

export const COMPUTING: any = /* @__PURE__ */ Symbol('COMPUTING');

export const ERRORED: any = /* @__PURE__ */ Symbol('ERRORED');

const COMPUTED_NODE = /* @__PURE__ */ (() => ({
  ...REACTIVE_NODE,
  value: UNSET,
  dirty: true,
  error: null,
  equal: defaultEquals,
  kind: 'computed',

  producerMustRecompute(node: ComputedNode<unknown>): boolean {
    return node.value === UNSET || node.value === COMPUTING;
  },

  producerRecomputeValue(node: ComputedNode<unknown>): void {
    if (node.value === COMPUTING) {
      // Our computation somehow led to a cyclic read of itself.
      throw new Error('Detected cycle in computations.');
    }

    const oldValue = node.value;
    node.value = COMPUTING;

    const prevConsumer = consumerBeginWork(node);
    let newValue: unknown;
    let wasEqual = false;
    try {
      newValue = node.computation();
      // We want to mark this node as errored if calling `equal` throws; however, we don't want
      // to track any reactive reads inside `equal`.
      setActiveConsumer(null);
      wasEqual = oldValue !== UNSET
          && oldValue !== ERRORED
          && newValue !== ERRORED
          && node.equal(oldValue, newValue);
    } catch (err) {
      newValue = ERRORED;
      node.error = err;
    } finally {
      consumerCompleteWork(node, prevConsumer);
    }

    if (wasEqual) {
      // No change to `valueVersion` - old and new values are
      // semantically equivalent.
      node.value = oldValue;
      return;
    }

    node.value = newValue;
    node.version++;
  }
}))();

export function createComputed<T>(computation: () => T): ComputedGetter<T> {
  const node: ComputedNode<T> = Object.create(COMPUTED_NODE);
  node.computation = computation;

  const getter = (() => {
    updateProducerValueVersion(node);
    accessProducer(node);
    if (node.value === ERRORED) {
      throw node.error;
    }
    return node.value;
  }) as ComputedGetter<T>;

  (getter as ComputedGetter<T>)[SIGNAL] = node;

  return getter;
}

export function computed<T>(computation: () => T, opt?: { equal: ValueEqualityFn<T> }): Signal<T> {
  const getter = createComputed(computation);
  if (opt?.equal) {
    getter[SIGNAL].equal = opt.equal;
  }
  return getter;
}
