import produce from 'immer';
import { ComInfo, FileType, GenerateType } from '@/types';
import { composePackageKey } from './keys';
import { getRealPaths } from './paths';
import storageHub from './storage';

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
