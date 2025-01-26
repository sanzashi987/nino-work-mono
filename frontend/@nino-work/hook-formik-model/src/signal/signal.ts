/* eslint-disable no-param-reassign */
import { defaultEquals, ValueEqualityFn } from './equality';
import { throwInvalidWriteToSignalError } from './errors';
import {
  accessProducer, producerIncrementEpoch, producerNotifyConsumers, producerUpdatesAllowed, REACTIVE_NODE, ReactiveNode, SIGNAL
} from './reactive';

let postSignalSetFn: (() => void) | null = null;
export function setPostSignalSetFn(fn: (() => void) | null): (() => void) | null {
  const prev = postSignalSetFn;
  postSignalSetFn = fn;
  return prev;
}

export function runPostSignalSetFn(): void {
  postSignalSetFn?.();
}

export interface SignalNode<T> extends ReactiveNode {
  value: T;
  equal: ValueEqualityFn<T>;
}

type SignalBaseGetter<T> = (() => T) & { readonly [SIGNAL]: unknown };

export interface SignalGetter<T> extends SignalBaseGetter<T> {
  readonly [SIGNAL]: SignalNode<T>;
}

export const SIGNAL_NODE: SignalNode<unknown> = /* @__PURE__ */ (() => ({
  ...REACTIVE_NODE,
  equal: defaultEquals,
  value: undefined,
  kind: 'signal'
}))();

export function createSignal<T>(value: T): SignalGetter<T> {
  const node: SignalNode<T> = Object.create(SIGNAL_NODE);
  node.value = value;
  const getter = (() => {
    accessProducer(node);
    return node.value;
  }) as SignalGetter<T>;
  (getter as any)[SIGNAL] = node;
  return getter;
}

function signalValueChanged<T>(node: SignalNode<T>): void {
  // eslint-disable-next-line no-plusplus
  node.version++;
  producerIncrementEpoch();
  producerNotifyConsumers(node);
  postSignalSetFn?.();
}

export function signalSetFn<T>(node: SignalNode<T>, newValue: T) {
  if (!producerUpdatesAllowed()) {
    throwInvalidWriteToSignalError();
  }

  if (!node.equal(node.value, newValue)) {
    node.value = newValue;
    signalValueChanged(node);
  }
}

export function signalUpdateFn<T>(node: SignalNode<T>, updater: (value: T) => T): void {
  if (!producerUpdatesAllowed()) {
    throwInvalidWriteToSignalError();
  }

  signalSetFn(node, updater(node.value));
}
