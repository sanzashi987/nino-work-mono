/* eslint-disable no-plusplus */
/* eslint-disable no-restricted-syntax */
/* eslint-disable @typescript-eslint/no-use-before-define */
/* eslint-disable no-param-reassign */
let activeConsumer: ReactiveNode | null = null;
let inNotificationPhase = false;

type Version = number & { __brand: 'Version' };

let epoch: Version = 1 as Version;
export function producerIncrementEpoch(): void {
  // eslint-disable-next-line no-plusplus
  epoch++;
}

export const SIGNAL = /* @__PURE__ */Symbol('SIGNAL');

export function setActiveConsumer(consumer: ReactiveNode | null): ReactiveNode | null {
  const prev = activeConsumer;
  activeConsumer = consumer;
  return prev;
}

export function getActiveConsumer(): ReactiveNode | null {
  return activeConsumer;
}
export function isInNotificationPhase(): boolean {
  return inNotificationPhase;
}

export function producerUpdatesAllowed(): boolean {
  return activeConsumer?.consumerAllowSignalWrites !== false;
}

export interface Reactive {
  [SIGNAL]: ReactiveNode;
}

export function isReactive(value: unknown): value is Reactive {
  return (value as Partial<Reactive>)[SIGNAL] !== undefined;
}

export type Signal<T> = (() => T) & {
  [SIGNAL]:unknown
};

export interface ReactiveNode {
  version: Version
  lastCleanEpoch: Version;
  dirty: boolean;

  producers?: ReactiveNode[];
  indexInThoseProducers?: number[]
  producerLastReadVersion?: Version[];
  nextProducerIndex: number;

  consumers?: ReactiveNode[];
  indexInThoseConsumers?: number[];

  consumerAllowSignalWrites: boolean;

  readonly consumerIsAlwaysLive: boolean;

  producerMustRecompute(node: unknown): boolean;
  producerRecomputeValue(node: unknown): void;
  consumerMarkedDirty(node: unknown): void;

  /**
 * Called when a signal is read within this consumer.
 */
  consumerOnSignalRead(node: unknown): void;

  kind: string;
}

interface ConsumerNode extends ReactiveNode {
  producers: NonNullable<ReactiveNode['producers']>;
  indexInThoseProducers: NonNullable<ReactiveNode['indexInThoseProducers']>;
  producerLastReadVersion: NonNullable<ReactiveNode['producerLastReadVersion']>;
}

interface ProducerNode extends ReactiveNode {
  consumers: NonNullable<ReactiveNode['consumers']>;
  indexInThoseConsumers: NonNullable<ReactiveNode['indexInThoseConsumers']>;
}

export function producerNotifyConsumers(node: ReactiveNode): void {
  if (node.consumers === undefined) {
    return;
  }

  // Prevent signal reads when we're updating the graph
  const prev = inNotificationPhase;
  inNotificationPhase = true;
  try {
    for (const consumer of node.consumers) {
      if (!consumer.dirty) {
        markConsumerDirty(consumer);
      }
    }
  } finally {
    inNotificationPhase = prev;
  }
}

export function markConsumerDirty(node: ReactiveNode): void {
  node.dirty = true;
  producerNotifyConsumers(node);
  node.consumerMarkedDirty?.(node);
}

export function assertConsumer(node: ReactiveNode): asserts node is ConsumerNode {
  node.producers ??= [];
  node.indexInThoseProducers ??= [];
  node.producerLastReadVersion ??= [];
}
export function assertProducer(node: ReactiveNode): asserts node is ProducerNode {
  node.consumers ??= [];
  node.indexInThoseConsumers ??= [];
}

function isConsumerLive(node: ReactiveNode): boolean {
  return node.consumerIsAlwaysLive || (node?.consumers?.length ?? 0) > 0;
}

function isConsumer(node: ReactiveNode): node is ConsumerNode {
  return node.producers !== undefined;
}

function producerRemoveConsumer(producer: ReactiveNode, idx: number): void {
  assertProducer(producer);

  if (producer.consumers.length === 1 && isConsumer(producer)) {
    // if last active consumer is remove , producer itself is not active any more,
    // so remove itself from its producers' consumer list
    for (let i = 0; i < producer.producers.length; i++) {
      producerRemoveConsumer(producer.producers[i], producer.indexInThoseProducers[i]);
    }
  }

  const lastIndex = producer.consumers.length - 1;
  producer.consumers[idx] = producer.consumers[lastIndex];
  producer.indexInThoseConsumers[idx] = producer.indexInThoseConsumers[lastIndex];

  producer.consumers.length--;
  producer.indexInThoseConsumers.length--;

  if (idx < producer.consumers.length) {
    const idxProducer = producer.indexInThoseConsumers[idx];
    const consumer = producer.consumers[idx];
    assertConsumer(consumer);
    consumer.indexInThoseProducers[idxProducer] = idx;
  }
}

function producerAddConsumer(producer: ReactiveNode, consumer: ReactiveNode, indexOfConsumer: number): number {
  assertProducer(producer);
  if (producer.consumers.length === 0 && isConsumer(producer)) {
    for (let i = 0; i < producer.producers.length; i++) {
      producer.indexInThoseProducers[i] = producerAddConsumer(producer.producers[i], producer, i);
    }
  }

  producer.indexInThoseConsumers.push(indexOfConsumer);
  return producer.consumers.push(consumer) - 1;
}

export function destroyConsumer(consumer: ReactiveNode) {
  assertConsumer(consumer);
  if (isConsumerLive(consumer)) {
    for (let i = 0; i < consumer.producers.length; i++) {
      producerRemoveConsumer(consumer.producers[i], consumer.indexInThoseProducers[i]);
    }
  }

  // eslint-disable-next-line no-multi-assign
  consumer.producers.length = consumer.producerLastReadVersion.length = consumer.indexInThoseProducers.length = 0;
  if (consumer.consumers) {
    // eslint-disable-next-line no-multi-assign
    consumer.consumers.length = consumer.indexInThoseConsumers.length = 0;
  }
}

function markProducerClean(node: ReactiveNode) {
  node.dirty = false;
  node.lastCleanEpoch = epoch;
}

export function consumerPollProducersForChange(node: ReactiveNode): boolean {
  assertConsumer(node);

  for (let i = 0; i < node.producers.length; i++) {
    const producer = node.producers[i];
    const seenVersion = node.producerLastReadVersion[i];

    if (seenVersion !== producer.version) {
      return true;
    }

    updateProducerValueVersion(producer);

    if (seenVersion !== producer.version) {
      return true;
    }
  }

  return false;
}

export function updateProducerValueVersion(node: ReactiveNode) {
  if (isConsumerLive(node) && !node.dirty) {
    return;
  }
  if (!node.dirty && node.lastCleanEpoch === epoch) {
    return;
  }

  if (!node.producerMustRecompute(node) && !consumerPollProducersForChange(node)) {
    // None of our producers report a change since the last time they were read, so no
    // recomputation of our value is necessary, and we can consider ourselves clean.
    markProducerClean(node);
    return;
  }

  node.producerRecomputeValue(node);

  markProducerClean(node);
}

export function accessProducer(node: ReactiveNode) {
  if (inNotificationPhase) {
    throw new Error('Assertion error: signal read during notification phase');
  }
  if (activeConsumer === null) {
    // Accessed outside of a reactive context, so nothing to record.
    return;
  }
  activeConsumer.consumerOnSignalRead(node);

  // This producer is the `idx`th dependency of `activeConsumer`.
  const idx = activeConsumer.nextProducerIndex++;

  assertConsumer(activeConsumer);
  // deps changes
  if (idx < activeConsumer.producers.length && activeConsumer.producers[idx] !== node) {
    if (isConsumerLive(activeConsumer)) {
      const staleProducer = activeConsumer.producers[idx];
      producerRemoveConsumer(staleProducer, activeConsumer.indexInThoseProducers[idx]);
    }
  }

  if (activeConsumer.producers[idx] !== node) {
    activeConsumer.producers[idx] = node;
    activeConsumer.indexInThoseProducers[idx] = isConsumerLive(activeConsumer) ? producerAddConsumer(node, activeConsumer, idx) : 0;
  }

  activeConsumer.producerLastReadVersion[idx] = node.version;
}

export function consumerBeginWork(node: ReactiveNode | null): ReactiveNode | null {
  // eslint-disable-next-line @typescript-eslint/no-unused-expressions
  node && (node.nextProducerIndex = 0);
  return setActiveConsumer(node);
}

export function consumerCompleteWork(
  node: ReactiveNode | null,
  prevConsumer: ReactiveNode | null
) {
  setActiveConsumer(prevConsumer);

  if (!node || node.producers === undefined || node.indexInThoseProducers === undefined || node.producerLastReadVersion === undefined) {
    return;
  }

  if (isConsumerLive(node)) {
    // remove subs for all producers outside the computation border
    for (let i = node.nextProducerIndex; i < node.producers.length; i++) {
      producerRemoveConsumer(node.producers[i], node.indexInThoseProducers[i]);
    }
  }

  while (node.producers.length > node.nextProducerIndex) {
    node.producers.pop();
    node.producerLastReadVersion.pop();
    node.indexInThoseProducers.pop();
  }
}

export const REACTIVE_NODE: ReactiveNode = {
  version: 0 as Version,
  lastCleanEpoch: 0 as Version,
  dirty: false,
  producers: undefined,
  producerLastReadVersion: undefined,
  indexInThoseProducers: undefined,
  nextProducerIndex: 0,
  consumers: undefined,
  indexInThoseConsumers: undefined,
  consumerAllowSignalWrites: false,
  consumerIsAlwaysLive: false,
  kind: 'unknown',
  producerMustRecompute: () => false,
  producerRecomputeValue: () => {},
  consumerMarkedDirty: () => {},
  consumerOnSignalRead: () => {}
};
