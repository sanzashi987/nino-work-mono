export type CanvixResponse<T = any> = {
  resultCode: number;
  data: T;
  resultMessage: string;
};

export type CanvixResponsePromise<T = any> = Promise<CanvixResponse<T>>;

export type Pagination = {
  pageIndex: number;
  pageSize: number;
  pageTotal: number;
  recordTotal: number;
};

type SandboxProps = {
  args: string[];
  // content: string;
  argsValue: any[];
  id: string;
  /** 是否在控制台输出报错信息，过滤器测试时不输出 */
  logVisible?: boolean;
};

export type SandboxRunnerType = (props: SandboxProps) => Promise<any | true>;
