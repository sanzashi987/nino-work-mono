import produce from 'immer';
import {
  ComInfo, FileType, GenerateType, LayerItem, MergeParams, PanelMetaRuntime,
  PanelMetaType
} from '@/types';
import { composeDeltaKey, composePackageKey } from './keys';
import { getRealPaths } from './paths';
import storageHub from './storage';
import { isDefaultTheme } from './palette';

type BaseInput = {
  attr?: Record<string, any>;
  basic?: Record<string, any>;
  com?: ComInfo;
  type: FileType;
};

type RuntimeConfigParams<T extends BaseInput = BaseInput> = {
  /** 组件config对象 */
  input: T;
  /** 需要转为运行时配置属性值 */
  runtimeKeys: Array<'basic' | 'attr'>;
  /** 用于转换处理的配置项 */
  config: {
    getAssetsUrl?: (fileName: string) => string;
  };
};

function isStartWithKey(path: string, keys: string[]) {
  return keys.reduce((prev, current) => prev && path.startsWith(current), true);
}

function parseSingleValue(params: {
  input: Record<string, any>;
  path: string;
  config: RuntimeConfigParams['config'];
  parse: GenerateType['parseValue'];
}) {
  const { input, parse, path, config } = params;
  const keyChain = path.split('.');
  const lastKey = keyChain.at(-1)!;
  const ref = keyChain.slice(0, -1).reduce<any>((a, b) => a?.[b], input);
  if (ref) {
    ref[lastKey] = parse!({ value: ref[lastKey], config });
  }
}

/**
 * 组件配置config转为运行时config
 * @param params
 * @returns
 */
export const getRuntimeConfig = <T extends BaseInput>(params: RuntimeConfigParams<T>): T => {
  const { input, config, runtimeKeys } = params;
  if (!input) return input;
  const { type, com } = input;
  // com字段不存在时，key为null
  const key = composePackageKey({ type, com });
  if (!key) return input;
  const packagePaths = storageHub.getPackagePaths();
  // console.log("packagePaths",packagePaths,input);
  const parsePath = packagePaths?.[key]?.parse;
  if (!parsePath) return input;

  // TODO 优化，减少脱离平台运行时的打包体积，不引入ui-components
  const { parseValueMap } = CWidgets.PropertyTransfrom;

  // 按路径长度降序（保证处理某个属性值时，其子值已全部处理）
  const sortList = Object.keys(parsePath).sort((a, b) => {
    const aLength = a.split('.').length;
    const bLength = b.split('.').length;
    return bLength - aLength;
  });

  const runtimeInput = runtimeKeys.reduce((prev, current) => {
    prev[current] = input[current];
    return prev;
  }, {} as Record<string, any>);

  const runtimeOutput = produce(runtimeInput, (draft) => {
    sortList.forEach((sortKey) => {
      const parse = parseValueMap[parsePath[sortKey]];
      if (parse && isStartWithKey(sortKey, runtimeKeys)) {
        // 包含$index时，特殊处理
        if (/\$index/.test(sortKey)) {
          const paths = getRealPaths(input, sortKey);
          paths.forEach((item) => {
            parseSingleValue({ input: draft, parse, path: item.path, config });
          });
        } else {
          parseSingleValue({ input: draft, path: sortKey, parse, config });
        }
      }
    });
  });

  return produce(input, (draft) => {
    runtimeKeys.forEach((k) => {
      draft[k] = runtimeOutput[k];
    });
  });
};

export function isDefaultBreakpoint(id: string) {
  return id === 'default';
}

export const createDeltaIdList = (params: Omit<MergeParams, 'delta' | 'core'>) => {
  const { id, theme, breakpoint, breakpoints = [] } = params;
  const breakpointIndex = breakpoints.findIndex((e) => e.id === breakpoint);
  const matchedBreakpoints = breakpoints
    .slice(0, breakpointIndex + 1)
    .map((e) => (isDefaultBreakpoint(e.id) ? '*' : e.id));
  const res: Array<keyof PanelMetaRuntime['delta']> = [];
  matchedBreakpoints.map((bk) => {
    // 非默认主题下，以默认主题为基础进行合并
    const themeKeys = isDefaultTheme(theme) ? ['*'] : ['*', theme];
    themeKeys.forEach((themeKey) => {
      const deltaKey = composeDeltaKey({
        globalBreakpoint: bk,
        localBreakpoint: '*',
        theme: themeKey,
        comId: id
      });
      res.push(deltaKey);
    });
  });
  return res;
};

const mergeSingleConfig = (object: Record<string, any>, source: Record<string, any>) => {
  Object.keys(source).forEach((key) => {
    const keyChain = key.split('.');
    const lastKey = keyChain[keyChain.length - 1];
    const ref = keyChain.slice(0, -1).reduce<any>((a, b) => a?.[b], object);
    if (ref) {
      ref[lastKey] = source[key];
    }
  });
  return object;
};

export const mergeConfig = (params: {
  defaultProperty: Record<string, any>;
  delta: PanelMetaRuntime['delta'];
  idList: Array<keyof PanelMetaRuntime['delta']>;
}) => {
  const { defaultProperty, delta, idList } = params;
  return produce(defaultProperty, (draft) => {
    idList.forEach((id) => {
      mergeSingleConfig(draft, delta[id] ?? {});
    });
  });
};

export const mergeCoreAndDelta = (params: MergeParams) => {
  const { delta, core = {}, id, theme, breakpoint, breakpoints } = params;
  const idList = createDeltaIdList({ id, theme, breakpoint, breakpoints });
  return mergeConfig({ defaultProperty: core, delta, idList });
};

/**
 * 递归获取组件的children
 * @param params
 */
const getLayerChildren = (params: {
  res: {
    children: LayerItem[];
  };
  layers: LayerItem[];
  comId: string;
}) => {
  const { layers, comId } = params;
  for (let i = 0; i < layers.length; i++) {
    if (layers[i].id === comId) {
      params.res.children = [...(layers[i].children || [])];
      break;
    }
    if (layers[i].children) {
      getLayerChildren({ res: params.res, layers: layers[i].children!, comId });
    }
  }
};

/**
 * 获取目标组件的children，并且包含com中的字段信息(仅第一层)
 * @param params
 */
export const getConfigChildren = (params: {
  panels: PanelMetaType;
  panelId: string;
  comId: string;
}) => {
  const { panelId, panels, comId } = params;
  const { layers } = panels[panelId];
  const layer: {
    children: LayerItem[];
  } = { children: [] };
  getLayerChildren({
    res: layer,
    layers,
    comId
  });
  return layer.children.map((item) => ({
    ...item,
    panelId,
    config: {
      ...panels[panelId].components[item.id]?.com,
      type: item.type
    }
  }));
};
