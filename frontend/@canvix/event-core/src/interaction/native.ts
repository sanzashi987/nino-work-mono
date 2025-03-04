import { InteractionService as CommonInteraction } from './common';
import {
  ConnectableContextType,
  InteractionInputPayloadType,
  InteractionPayloadType,
  NativeConnectableContextType,
  PostMethodType,
  PushMethodType,
  ServiceHostInstance
} from '../types';

class InteractionService extends CommonInteraction {
  bridgeId;

  instanceId;

  channel;

  constructor(
    scope: ServiceHostInstance<{
      panelMeta: ConnectableContextType;
      globalMeta: NativeConnectableContextType;
    }>,
    post: PostMethodType<InteractionPayloadType>,
    push: PushMethodType<InteractionInputPayloadType>
  ) {
    super(scope, post, push);
    const { getProcessEnv } = scope.props.globalMeta;
    const { channel, instanceId, bridgeId, platform } = getProcessEnv();
    this.bridgeId = bridgeId;
    this.instanceId = instanceId;
    this.channel = channel;

    this.push = (id, payload) => {
      const { targetDescriber, target, targetNode, value } = payload.data;
      // flutter env specific logic
      if (targetDescriber?.useChannel[platform] === true) return this.sendMessageToNative(targetNode, target, value);
      push(id, payload);
    };

    window.addEventListener(bridgeId, this.onNativeMessage);
  }

  onNativeMessage = (e: CustomEvent) => {
    let res;
    try {
      res = JSON.parse(e.detail);
    } catch (e) {
      return;
    }
    const { type, value } = res ?? {};
    if (type !== 'message') return;
    const { node, event, params } = value;
    this.emit(node, event, params);
  };

  sendMessageToNative = (nodeAlias: string, handlerId: string, params: Record<string, any>) => {
    this.channel.postMessage(
      JSON.stringify({
        type: 'message',
        value: { instanceId: this.instanceId, node: nodeAlias, handler: handlerId, params }
      })
    );
  };

  onDestory() {
    window.removeEventListener(this.bridgeId, this.onNativeMessage);
  }
}

export { InteractionService };
