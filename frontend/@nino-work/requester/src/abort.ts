export type AbortConfig = {
  timeout?: number;
};

class NinoFetchAbortController extends AbortController {
  public timer;

  constructor(props?: AbortConfig) {
    super();

    /** default 30s abort */
    const timeout = props?.timeout ?? 30000;
    if (timeout > 0) {
      this.timer = setTimeout(() => {
        this.abort('request timeout');
      }, timeout);
    }
  }

  override abort(reason?: any) {
    clearTimeout(this.timer);
    super.abort(reason);
  }
}

export default NinoFetchAbortController;
