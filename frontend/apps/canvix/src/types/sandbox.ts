export type RunInSandboxProps = {
  args: string[];
  // content: string;
  argsValue: any[];
  id: string;
  /** 是否在控制台输出报错信息，过滤器测试时不输出 */
  logVisible?: boolean;
};

export type ReturnMessageType = {
  id: string;
  res: any;
  error: boolean;
  workerId: string;
};

export type WorkerItem = {
  id: string;
  state: 'running' | 'idle';
  worker: Worker;
};

export type PromisePair = [(val: any) => void, (val: any) => void];

export type SandboxRunnerType = (props:RunInSandboxProps)=> Promise<any>;
