/* eslint-disable no-restricted-syntax */

export interface Schedulable {
  run: VoidFunction;
}

export class Scheduler {
  private queue = new Set<Schedulable>();

  private hasTask = false;

  schedule(handle: Schedulable) {
    this.enqueue(handle);
    if (this.hasTask === false) {
      this.hasTask = true;
      queueMicrotask(() => {
        this.flush();
      });
    }
  }

  remove(handle: Schedulable): void {
    const { queue } = this;
    if (!queue.has(handle)) {
      return;
    }
    queue.delete(handle);
  }

  private enqueue(handle: Schedulable): void {
    const { queue } = this;
    if (queue.has(handle)) {
      return;
    }
    queue.add(handle);
  }

  flush() {
    try {
      const { queue } = this;
      for (const handle of queue) {
        queue.delete(handle);
        handle.run();
      }
    } finally {
      if (this.hasTask !== true) {
        this.hasTask = false;
      }
    }
  }
}
export const defaultScheduler = new Scheduler();
