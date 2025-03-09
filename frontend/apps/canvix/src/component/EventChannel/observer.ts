import { PanelServiceInstance, PostMethodType, PushMethodType } from './types';
import { Connector } from './channel';

export const FetchInstance = 'observer.fetchInstance';
export const TransferInstance = 'observer.transferInstance';

abstract class ObserverService implements PanelServiceInstance {
  static $name = 'observer';

  static $supportedEvents = /^!(.*)$/;

  static $responsive = false;

  handle: (payload: any) => void;

  constructor(scope: any, post: PostMethodType<any>, private push: PushMethodType<any>) {
    const eventHub = this.getEventHub();
    eventHub.on(FetchInstance, this.fetchComponentScope);
    this.handle = function (this: Connector) {
      // eslint-disable-next-line @typescript-eslint/naming-convention, @typescript-eslint/no-this-alias
      const _this = this;
      eventHub.emit(TransferInstance, _this);
    };
  }

  abstract getEventHub(): {
    // eslint-disable-next-line @typescript-eslint/ban-types
    on(event: string, callback: Function): void;
    emit(event: string, payload: any): void;
  };

  fetchComponentScope = (id: string) => {
    this.push(id, {
      type: ObserverService.$name,
      data: id
    });
  };

  onSchedule() {}

  emit() {}

  updateConfig() {}
}

export default ObserverService;
