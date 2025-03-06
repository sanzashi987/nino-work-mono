import {
  ServiceComponent, typeToService, parseConfig, EndpointsType, EndpointType
} from '@canvix/component-factory';
import { BLOCK_ACTION_ID, BLOCK_EVENT_ID, InteractionNodeType } from '@canvix/event-core';
import { BasicCanvixFieldType, ComponentPackageType } from '@canvix/shared';
import { ConfigType } from 'dayjs';

/** ****************** pure functions *****************  */
export type EndpointResType = {
  source: EndpointType[];
  target: EndpointType[];
  childList: string[];
};

function addPrefix(input: EndpointsType) {
  return Object.entries(input ?? {}).map(([k, v]) => ({
    ...v,
    id: `instance.${k}`
  }));
}

function getConfiguredEndpoints(packageJSON: BasicCanvixFieldType) {
  const { events, handlers } = packageJSON!;
  const source = addPrefix(events ?? {});
  const target = addPrefix(handlers ?? {});
  return { source, target };
}

function addPrefixArr(input: ReturnType<typeof addPrefix>, prefix = 'instance') {
  return input.map((e) => ({ ...e, id: `${prefix}.${e.id}` })) as any;
}

function addToRes(
  res: EndpointResType,
  source: EndpointType[],
  target: EndpointType[],
  child?: string[]
) {
  res.target = res.target.concat(target);
  res.source = res.source.concat(source);
  if (child) res.childList = res.childList.concat(child);
}

function addToResFromServices(
  endpoints: EndpointResType,
  services: ServiceComponent[],
  packageJSON: any
) {
  services.forEach((ctor) => {
    const {
      getComponentActions,
      getComponentEvents,
      events = new Map(),
      handlers = new Map(),
      serviceName
    } = ctor;
    addToRes(
      endpoints,
      getComponentEvents(events, serviceName, packageJSON),
      getComponentActions(handlers, serviceName, packageJSON)
    );
  });
}

function addToResFromPackageJson(res: EndpointResType, packageJSON: any) {
  const { source, target } = getConfiguredEndpoints(packageJSON);
  addToRes(res, source, target);
}

/** ****************** factories ****************** */

type StandardLogicalModule = {
  default: {
    getComponentEvents?: (config: any) => EndpointType[];
    getComponentActions?: (config: any) => EndpointType[];
    getComponentChildList?: (config: any) => string[];
  };
};

type ModuleLoader = (id: string, name: string, version: string) => Promise<StandardLogicalModule>;

export type PackageLoaderParams = {
  name: string;
  version: string;
  user?: string | null;
  latest?: boolean;
  prefix?: string;
  isDebugger?: boolean;
};

type PackageLoader = (
  params: PackageLoaderParams,
) => Promise<ComponentPackageType | null>;

type DeprecatedDefiner = (id: string) => readonly [boolean, boolean];

type ComponentFinder = (id: string) => ConfigType | null;

type PanelNodeControlTypeGetter = (
  panelId: string,
) => readonly [Record<string, InteractionNodeType>, boolean];

type GetLogicalEndpointsProps = {
  id: string;
  attr?: Record<string, any>;
  data?: any;
  com: InteractionNodeType['com'];
};

const createGetEndpoints = (
  moduleLoader: ModuleLoader,
  packageLoader: PackageLoader,
  componentFinder: ComponentFinder,
  panelNodeControlTypeGetter: PanelNodeControlTypeGetter,
  deprecatedDefiner: DeprecatedDefiner
) => {
  const getLogicalEndpoints = async ({ id, com, attr, data }: GetLogicalEndpointsProps) => {
    const { version, name, user } = com!;
    const res: EndpointResType = { source: [], target: [], childList: [] };
    const nodeModule = await moduleLoader(id, name, version);
    const packageJSON = await packageLoader({ name, version, user });
    const { getComponentActions, getComponentEvents, getComponentChildList } = nodeModule.default;
    // Return skeleton layout when fetching error packageJSON
    if (!packageJSON) return [res, false] as const;
    addToResFromPackageJson(res, packageJSON);
    // const { source, target } = getConfiguredEndpoints(packageJSON!);
    // addToRes(res, source, target);
    if (data) {
      addToResFromServices(res, [typeToService.data], packageJSON);
    }

    addToRes(
      res,
      addPrefixArr(getComponentEvents?.(attr) ?? []),
      addPrefixArr(getComponentActions?.(attr) ?? []),
      getComponentChildList?.(attr) ?? []
    );
    return [res, !!packageJSON?.verticalLayout] as const;
  };

  const getBlockEndpoints = async ({ id, com, attr }: GetLogicalEndpointsProps) => {
    const { version, name } = com!;
    const nodeModule = await moduleLoader(id, name, version);
    const { getComponentEvents, getComponentActions } = nodeModule.default;
    // let prefix: string | false = false,
    let cb: StandardLogicalModule['default']['getComponentActions'] = () => [];
    if (id === BLOCK_ACTION_ID) {
      // prefix = BLOCK_ACTION_PREFIX;
      cb = getComponentEvents;
    } else if (id === BLOCK_EVENT_ID) {
      // prefix = BLOCK_EVENT_PREFIX;
      cb = getComponentActions;
    }
    // if (!prefix) throw new Error('Not a block interface node');
    return addPrefixArr(cb?.(attr) ?? []) as EndpointType[];
  };

  const getPanelEndpoints = async (
    res: EndpointResType,
    nodes: Record<string, InteractionNodeType>,
    dataControlled: boolean
  ) => {
    const actionNode = nodes[BLOCK_ACTION_ID];
    const eventNode = nodes[BLOCK_EVENT_ID];
    const loop = [
      {
        name: 'source',
        node: eventNode || null
      },
      {
        name: 'target',
        node: !dataControlled && actionNode ? actionNode : null
      }
    ] as const;

    // eslint-disable-next-line no-restricted-syntax
    for (const { name, node } of loop) {
      if (!node) {
        // eslint-disable-next-line no-continue
        continue;
      }
      // eslint-disable-next-line no-await-in-loop
      const endpoint = await getBlockEndpoints({
        id: node.id,
        com: node.com!,
        attr: node.attr ?? {}
      });
      res[name] = [...res[name], ...endpoint];
    }
  };

  const getComponentEndpoints = async (endpoints: EndpointResType, id: string, type: string) => {
    const comConfig = componentFinder(id);
    if (!comConfig) return;
    const com = (comConfig as any).com!;
    const packageJSON: any = (await packageLoader({
      name: com.name,
      version: com.version,
      user: com.user
    })) || {};

    const usedCtor = parseConfig(typeToService, comConfig);
    // const usedCtor: any[] = [];
    if (getComponentEndpoints.BasicType.has(type)) {
      usedCtor.push(typeToService.basic);
    }
    if (getComponentEndpoints.AttrType.has(type)) {
      usedCtor.push(typeToService.attr);
    }
    addToResFromServices(endpoints, usedCtor, packageJSON);
    addToResFromPackageJson(endpoints, packageJSON);
  };
  getComponentEndpoints.BasicType = new Set(['com', 'group', 'refPanel', 'container']);
  getComponentEndpoints.AttrType = new Set(['com', 'subcom', 'refPanel']);

  const getNormalNodeEndpoints = async (id: string, type: string) => {
    const endpoints: EndpointResType = { source: [], target: [], childList: [] };
    let deprecated = false;
    let isDelete = false;
    [isDelete, deprecated] = deprecatedDefiner(id);
    if (!isDelete) {
      if (getNormalNodeEndpoints.panelType.includes(type)) {
        const [nodes, dataControlled] = panelNodeControlTypeGetter(id);
        await getPanelEndpoints(endpoints, nodes, dataControlled);
      } else {
        await getComponentEndpoints(endpoints, id, type);
      }
    }
    return { deprecated, endpoints, isDelete };
  };
  getNormalNodeEndpoints.panelType = ['panel', 'subpanel'];

  return {
    getLogicalEndpoints,
    // getBlockEndpoints,
    // getPanelEndpoints,
    // getComponentEndpoints,
    getNormalNodeEndpoints
  };
};

export { createGetEndpoints };
