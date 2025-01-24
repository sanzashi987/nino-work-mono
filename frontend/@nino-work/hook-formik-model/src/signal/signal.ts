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

export interface Reactive {
  [SIGNAL]: ReactiveNode;
}

export function isReactive(value: unknown): value is Reactive {
  return (value as Partial<Reactive>)[SIGNAL] !== undefined;
}

interface ReactiveNode {
  version: Version
  lastCleanEpoch: Version;
  dirty: boolean;

  producerNode?: ReactiveNode[];
  indexInThoseProducer?: number[]
  producerLastReadVersion?: Version[];
  nextProducerIndex: number;

  consumerNode?: ReactiveNode[];
  indexInThoseConsumer?: number[];

  consumerAllowSignalWrites: boolean;

  readonly consumerIsAlwaysLive: boolean;

  consumerMarkedDirty?(node: unknown): void

  /**
 * Called when a signal is read within this consumer.
 */
  consumerOnSignalRead(node: unknown): void;

  kind: string;
}

interface ConsumerNode extends ReactiveNode {
  producerNode: NonNullable<ReactiveNode['producerNode']>;
  indexInThoseProducer: NonNullable<ReactiveNode['indexInThoseProducer']>;
  producerLastReadVersion: NonNullable<ReactiveNode['producerLastReadVersion']>;
}

interface ProducerNode extends ReactiveNode {
  consumerNode: NonNullable<ReactiveNode['consumerNode']>;
  indexInThoseConsumer: NonNullable<ReactiveNode['indexInThoseConsumer']>;
}

export function producerNotifyConsumers(node: ReactiveNode): void {
  if (node.consumerNode === undefined) {
    return;
  }

  // Prevent signal reads when we're updating the graph
  const prev = inNotificationPhase;
  inNotificationPhase = true;
  try {
    for (const consumer of node.consumerNode) {
      if (!consumer.dirty) {
        consumerMarkDirty(consumer);
      }
    }
  } finally {
    inNotificationPhase = prev;
  }
}

export function consumerMarkDirty(node: ReactiveNode): void {
  node.dirty = true;
  producerNotifyConsumers(node);
  node.consumerMarkedDirty?.(node);
}

export function assertConsumer(node: ReactiveNode): asserts node is ConsumerNode {
  node.producerNode ??= [];
  node.indexInThoseProducer ??= [];
  node.producerLastReadVersion ??= [];
}
export function assertProducer(node: ReactiveNode): asserts node is ProducerNode {
  node.consumerNode ??= [];
  node.indexInThoseConsumer ??= [];
}

function consumerIsLive(node: ReactiveNode): boolean {
  return node.consumerIsAlwaysLive || (node?.consumerNode?.length ?? 0) > 0;
}

function isConsumer(node: ReactiveNode): node is ConsumerNode {
  return node.producerNode !== undefined;
}

function produceRemoveConsumer(producer: ReactiveNode, idx: number): void {
  assertProducer(producer);

  if (producer.consumerNode.length === 1 && isConsumer(producer)) {
    // if last active consumer is remove , producer itself is not active any more,
    // so remove itself from its producers' consumer list

  }
}
