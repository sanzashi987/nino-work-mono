import { WorkerItem, RunInSandboxProps, ReturnMessageType, SandboxRunnerType } from '@/types';
import { uuid } from '@/utils';

type PromisePair = [(val: any) => void, (val: any) => void];

const timeout = 10 * 1000;
const maxWorkers = window.navigator.hardwareConcurrency || 4;

class ScriptSandbox {
  pending: Record<string, PromisePair> = {};

  taskQueue: Array<RunInSandboxProps> = [];

  workers: Record<string, WorkerItem> = {};

  // Leave the mounting action to specific apps
  // constructor() {
  //   (Component.prototype as any).runInSandbox = this.runInSandbox;
  // }

  getIdleWorker = (): WorkerItem | undefined => {
    const workerList = Object.values(this.workers);
    if (workerList.length < 1) return undefined;
    return workerList.find((item) => item.state === 'idle');
  };

  enQueue = (task: RunInSandboxProps) => {
    this.taskQueue.push(task);
  };

  deQueueAndRun = (worker: WorkerItem) => {
    if (this.taskQueue.length < 1) return;
    const task = this.taskQueue.shift();
    this._runInSandbox(task!, worker);
  };

  createWorker = () :WorkerItem | null => {
    if (Object.keys(this.workers).length >= maxWorkers) return null;
    const newWorker = new Worker('/sandbox.worker.js');
    const id = `worker_${uuid(5)}`;
    const worker: WorkerItem = {
      id,
      worker: newWorker,
      state: 'idle'
    };
    this.workers[id] = worker;
    return worker;
  };

  afterWorkerTimeout = (workerId: string) => {
    // 超时终止worker，默认10s
    this.workers[workerId].worker.terminate();
    // 销毁当前worker
    delete this.workers[workerId];
    if (this.taskQueue.length < 1) return;
    // 如果任务队列非空，新增worker并执行
    const newWorker = this.createWorker();
    if (newWorker) {
      this.deQueueAndRun(newWorker);
    }
  };

  afetrWorkerReturn = (workerId: string) => {
    // 将当前worker置为空闲状态
    this.workers[workerId].state = 'idle';
    // 如果任务队列非空，则执行
    this.deQueueAndRun(this.workers[workerId]);
  };

  onReturn = (e: any, timer: number) => {
    clearTimeout(timer);
    const { id, res, error, workerId } = e.data as ReturnMessageType;
    if (this.pending[id]) {
      const [resolve, reject] = this.pending[id];
      if (error === false) {
        resolve?.(res);
      } else {
        reject(error);
      }
    }
    delete this.pending[id];
    this.afetrWorkerReturn(workerId);
  };

  _runInSandbox = (props: RunInSandboxProps, worker: WorkerItem) => {
    const { id, worker: sandbox } = worker;
    this.workers[id].state = 'running';
    sandbox.onerror = function (err) {
      // eslint-disable-next-line @typescript-eslint/no-throw-literal
      throw err;
    };
    const timer = setTimeout(() => {
      delete this.pending[props.id];
      this.afterWorkerTimeout(worker.id);
      // rej(new Error('运行超时'));
      console.log(
        `%c *** ${id} runs failed in sandbox ***`,
        'background:#fc7b2120;color:#fc7b21',
        `timeout of ${timeout}ms exceeded`
      );
    }, timeout);
    sandbox.onmessage = (e: any) => this.onReturn(e, timer);
    sandbox?.postMessage({ ...props, workerId: id });
  };

  runInSandbox: SandboxRunnerType = (props) => new Promise((res, rej) => {
    try {
      const idleWorker = this.getIdleWorker();
      if (idleWorker) {
        // 存在空闲worker，立即执行
        this._runInSandbox(props, idleWorker);
      } else if (Object.keys(this.workers).length < maxWorkers) {
        // worker未达上限，创建新worker并执行
        const newWorker = this.createWorker();

        if (newWorker) {
          this._runInSandbox(props, newWorker);
        }
      } else {
        // 无空闲worker并且worker上限，进入队列等待
        this.enQueue(props);
      }
    } catch (e) {
      console.log('fail to initialize the sandbox', e);
    }
    this.pending[props.id] = [res, rej];
  });
}

export default new ScriptSandbox();
