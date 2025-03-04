import { ENABLE_REFLECT_PANEL, PANEL_LOCAL_ID } from '../consts';
import type {
  PanelServiceInstance,
  PostMethodType,
  TargetDescription,
  InteractionPayloadType,
  SourceDesciption,
  InteractionOutputPayloadType,
  InteractionInputPayloadType,
  PushMethodType,
  EdgeBasic,
  ConnectableContextType,
  ServiceHostInstance
} from '../types';
import type { ConnectorCore } from '../channel';

function genIdentifier({ sourceNode, source }: SourceDesciption) {
  return `${sourceNode}.${source}`;
}

function composeEdge({
  target,
  targetNode,
  source,
  sourceNode,
  ...other
}: EdgeBasic): [string, TargetDescription] {
  return [genIdentifier({ source, sourceNode }), { target, targetNode, ...other }];
}

function getSubscribers(
  publisher: SourceDesciption,
  subMap: Map<string, TargetDescription[]>
): TargetDescription[] {
  const identifier = genIdentifier(publisher);
  return subMap.has(identifier) ? subMap.get(identifier)! : [];
}

class InteractionService implements PanelServiceInstance<InteractionPayloadType> {
  static $name = 'interaction';

  static $supportedEvents = /^(\w+\.)+\w+$/;

  static $responsive = true;

  // static $enhancer = InteractionEnhancer;
  static defaultEdgeDescriber: [string, TargetDescription[]][] = ENABLE_REFLECT_PANEL;

  private sourceTargetMap = new Map<string, TargetDescription[]>();

  private interactionConfig!: ConnectableContextType['interaction'];

  private dashboardComponents!: ConnectableContextType['components'];

  constructor(
    public scope: ServiceHostInstance<{
      panelMeta: ConnectableContextType;
    }>,
    private post: PostMethodType<InteractionPayloadType>,
    protected push: PushMethodType<InteractionInputPayloadType>
  ) {
    this.updateConfig();
  }

  private validateNode(id: string) {
    const triggerNode = this.interactionConfig.nodes[id];
    if (id === PANEL_LOCAL_ID) return true;
    if (!triggerNode || triggerNode.disable) return false;
    if (triggerNode.type === 'logical') return true;
    const com = this.dashboardComponents[id];
    return !!com;
  }

  // make payload from `InteractionOutputPayloadType` to `InteractionInputPayloadType`
  onSchedule(payload: InteractionOutputPayloadType) {
    if (!this.validateNode(payload.sourceNode)) return;
    getSubscribers(payload, this.sourceTargetMap)
      .filter(({ targetNode }) => {
        const node = this.interactionConfig.nodes[targetNode];
        return (node && !node.disable) || targetNode === PANEL_LOCAL_ID;
      })
      .forEach((subscriber) => {
        const outPayload = { ...subscriber, value: payload.value };
        this.push(subscriber.targetNode, {
          type: InteractionService.$name,
          data: outPayload
        });
      });
  }

  emit = (id: string, eventName: string, value: any) => {
    const payload = {
      type: InteractionService.$name,
      data: {
        value,
        source: eventName,
        sourceNode: id
      }
    };
    this.post(payload);
  };

  handle(this: ConnectorCore, { target, value }: InteractionInputPayloadType) {
    try {
      const route = target.split('.');
      // the services from subscribers
      let fn: any = this.getServices();
      for (const property of route) {
        fn = fn[property];
        if (!fn) return;
      }
      if (typeof fn === 'function') {
        fn(value);
      }
    } catch (e) {
      console.log('error parse action', e);
    }
  }

  updateConfig() {
    const needInstall = this.loadData();
    needInstall && this.install();
  }

  private loadData() {
    const { components, interaction } = this.scope.props.panelMeta; // as ConnectableContextType;
    if (this.dashboardComponents !== components) {
      this.dashboardComponents = components;
    }
    if (this.interactionConfig !== interaction) {
      this.interactionConfig = interaction;
      return true;
    }
    return false;
  }

  private install() {
    const newMap = new Map<string, TargetDescription[]>(InteractionService.defaultEdgeDescriber);
    Object.values(this.interactionConfig.edges).forEach((edge) => {
      if (edge.disable) return;
      const [identifier, target] = composeEdge(edge);
      newMap.has(identifier)
        ? newMap.get(identifier)?.push(target)
        : newMap.set(identifier, [target]);
    });
    this.sourceTargetMap = newMap;
  }
}

export { InteractionService };
