import produce from 'immer';
import { hide, display, toggleVisible, init, unmount } from './consts';
import ProtoService from '../proto-service';
import { event, handler, service } from '../proto-service/annotations';

type TransitionInputType = {
  type: string;
  duration: number;
};
@service('instance', 'com')
class InstanceService extends ProtoService {
  @event('实例', init)
  init = 'init';

  @handler('显示', display)
  display = (config?: TransitionInputType) => {
    this.props.setState((prev: any) => {
      if (!prev.config.hide) return null;
      // this.props.transitionRef?.current?.['display']?.(config);
      this.props.instanceRef.current?.['display']?.();
      return produce(prev, (draft: any) => {
        // delete draft.config.hide;
        draft.config.hide = 0;
      });
    });
  };

  @handler('隐藏', hide)
  hide = (config?: TransitionInputType) => {
    this.props.setState((prev: any) => {
      if (prev.config.hide === 1) return null;
      // this.props.transitionRef?.current?.['hide']?.(config);
      this.props.instanceRef.current?.['hide']?.();
      return produce(prev, (draft: any) => {
        draft.config.hide = 1;
      });
    });
  };

  @handler('卸载', unmount)
  unmount = (config?: TransitionInputType) => {
    this.props.setState((prev: any) => {
      if (prev.config.hide === 2) return null;
      // this.props.transitionRef?.current?.['hide']?.(config);
      this.props.instanceRef.current?.['unmount']?.();
      return produce(prev, (draft: any) => {
        draft.config.hide = 2;
      });
    });
  };

  // @handler('切换显隐', toggleVisible)
  // toggleVisible = (config?: { display: TransitionInputType; hide: TransitionInputType }) => {
  //   this.props.setState((prev: any) =>
  //     prev.config.hide ? this.display(config?.display) : this.hide(config?.hide),
  //   );
  // };

  componentDidMount() {
    this.props.selfRef.current = new Proxy(this, {
      get: (target: any, propKey: any) => {
        if (target[propKey]) return target[propKey];
        return target.props.instanceRef.current?.[propKey];
      },
    });
  }
}

export default InstanceService;
