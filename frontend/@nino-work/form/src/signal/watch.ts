import {
  consumerBeginWork,
  consumerCompleteWork,
  consumerPollProducersForChange,
  destroyConsumer,
  isInNotificationPhase,
  markConsumerDirty,
  REACTIVE_NODE,
  ReactiveNode,
  SIGNAL,
} from './reactive';
import { defaultScheduler, Schedulable, Scheduler } from './scheduler';

const NOOP_CLEANUP_FN = () => {};

const WATCH_NODE: Partial<WatchNode> = /* @__PURE__ */ (() => ({
  ...REACTIVE_NODE,
  consumerIsAlwaysLive: true,
  consumerAllowSignalWrites: false,
  consumerMarkedDirty: (node: WatchNode) => {
    if (node.schedule !== null) {
      node.schedule(node.ref);
    }
  },
  hasRun: false,
  cleanupFn: NOOP_CLEANUP_FN,
}))();

export interface WatchNode extends ReactiveNode {
  hasRun: boolean;
  ref: Watch;
  fn: (() => VoidFunction) | VoidFunction;
  cleanupFn: VoidFunction;
  schedule: ((watch: Watch) => void) | null;
}

export interface Watch {
  notify: VoidFunction;
  run: VoidFunction;
  cleanup: VoidFunction;
  destroy: VoidFunction;
  [SIGNAL]: WatchNode;
}

export function createWatch(
  fn: WatchNode['fn'],
  schedule: WatchNode['schedule'],
  allowSignalWrites = false
): Watch {
  const node: WatchNode = Object.create(WATCH_NODE);

  if (allowSignalWrites) {
    node.consumerAllowSignalWrites = true;
  }

  node.fn = fn;
  node.schedule = schedule;

  function isDestroyed(n: WatchNode) {
    return n.fn === null && n.schedule === null;
  }

  function destroy(n: WatchNode) {
    if (!isDestroyed(n)) {
      destroyConsumer(n);
      node.fn = null;
      node.schedule = null;
      node.cleanupFn = NOOP_CLEANUP_FN;
    }
  }

  function run() {
    if (node.fn === null) {
      return;
    }

    if (isInNotificationPhase()) {
      throw new Error('in notification phase');
    }

    node.dirty = false;

    if (node.hasRun && !consumerPollProducersForChange(node)) {
      return;
    }
    node.hasRun = true;

    const prevConsumer = consumerBeginWork(node);
    try {
      node.cleanupFn();
      node.cleanupFn = NOOP_CLEANUP_FN;
      const possibleCleanUp = node.fn();
      if (typeof possibleCleanUp === 'function') {
        node.cleanupFn = possibleCleanUp;
      }
    } finally {
      consumerCompleteWork(node, prevConsumer);
    }
  }
  node.ref = {
    notify: () => markConsumerDirty(node),
    run,
    cleanup: () => node.cleanupFn(),
    destroy: () => destroy(node),
    [SIGNAL]: node,
  };

  return node.ref;
}

class EffectHandle implements Schedulable {
  readonly watcher: Watch;

  constructor(
    scheduler: Scheduler,
    private effectFn: WatchNode['fn'],
    allowSignalWrites: boolean
  ) {
    this.watcher = createWatch(effectFn, () => scheduler.schedule(this), allowSignalWrites);
  }

  run() {
    this.watcher.run();
  }

  destroy() {
    this.watcher.destroy();
  }
}

export function effect(
  effectFn: WatchNode['fn'],
  opt?: { allowSignalWrites?: boolean; scheduler?: Scheduler }
): { destroy: VoidFunction } {
  const handle = new EffectHandle(
    opt.scheduler ?? defaultScheduler,
    effectFn,
    opt?.allowSignalWrites ?? false
  );

  handle.watcher.notify();
  return handle;
}
