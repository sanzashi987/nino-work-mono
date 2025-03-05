/**
 * @description  RefreshTimer
 * @author wangsenyang
 * @date   2020-8-21
 */

type PromiseMethodType = (...args: any[]) => Promise<any>;

export type TimerSetterProps = {
  // id: string;
  promiseMethod: PromiseMethodType;
  times: number;
};
export default class RefreshTimer {
  private current: number | null = null;

  private promiseMethod: PromiseMethodType = async () => null;

  private timestamp = '';

  setTimer(cb: () => void, times: number): void {
    this.current = setTimeout(cb, times);
  }

  stop(): void {
    if (this.current) clearTimeout(this.current);
    this.current = null;
  }

  genTimestamp() {
    const timestamp = new Date().toISOString();
    this.timestamp = timestamp;
    return timestamp;
  }

  setTimerForTarget({
    promiseMethod,
    times
  }: TimerSetterProps): (args?: Record<string, any>) => Promise<any> {
    // eslint-disable-next-line @typescript-eslint/no-this-alias
    const self = this;
    this.genTimestamp();
    this.promiseMethod = promiseMethod;
    return async function actuator(
      args?: Record<string, any>,
      _timeStamp: string = self.timestamp
    ): Promise<any> {
      if (_timeStamp !== self.timestamp) return;
      const timestamp = self.genTimestamp();
      try {
        await self.promiseMethod(args);
      } catch (error) {
        Promise.reject(error);
      } finally {
        self.stop();
        self.setTimer(() => {
          actuator(args, timestamp);
        }, times);
      }
    };
  }

  destory() {
    this.stop();
    this.setTimer = () => null;
    this.timestamp = '';
    this.promiseMethod = async () => null;
  }
}
