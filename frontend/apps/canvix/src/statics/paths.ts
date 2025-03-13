import { CObjectValueType } from '@/types';

type ResItem = { paths: Array<string | number>; index: number };

/**
 * 递归生成真实路径
 * @param params
 * @returns
 */
const recursiveGeneratePath = (params: {
  /** 组件当前config */
  config: any;
  /** 返回结果 */
  res: ResItem[];
  /** 当前路径及index */
  current: ResItem;
  /** 剩余路径 */
  keyChain: string[];
}) => {
  const { config, res, current, keyChain } = params;
  let newConfig = config;
  for (let i = 0; i < keyChain.length; i++) {
    // 仅CTabs等对象数组或CObjectTabs等CObjectValueType类型组件存在$index
    if (keyChain[i] === '$index') {
      const isArray = Object.prototype.toString.call(newConfig) === '[object Array]';
      const newKeyChain = keyChain.slice(i + 1);
      const list: Array<{
        index: number;
        key: string | number;
      }> = isArray
      // 兼容旧版本CTabs数据格式Array<Record<string, any>>
        ? (newConfig as Array<Record<string, any>>).map((e, idx) => ({ key: idx, index: idx }))
      // CObjectValueType类型数据
        : Object.entries(newConfig as CObjectValueType).map((e) => ({
          key: e[0],
          index: e[1].order
        }));
      // eslint-disable-next-line @typescript-eslint/no-loop-func
      list.forEach((item) => {
        const { key, index } = item;
        const tempConfig = newConfig[key];
        const newCurrent: ResItem = {
          paths: [...current.paths, key],
          index
        };
        if (i === keyChain.length - 1) {
          res.push(newCurrent);
        } else {
          recursiveGeneratePath({
            res,
            config: tempConfig,
            current: newCurrent,
            keyChain: newKeyChain
          });
        }
      });
      break;
    } else {
      newConfig = newConfig[keyChain[i]];
      current.paths.push(keyChain[i]);
      if (i === keyChain.length - 1) {
        res.push(current);
      }
    }
  }
  return res;
};

/**
 * 替换占位符，获取真实更新路径
 * @param config 组件config
 * @param path 配置路径，如"a.b.c.$index.d"
 * @returns
 */
export const getRealPaths = (
  config: Record<string, any>,
  path: string
): Array<{ path: string; index: number }> => {
  // "a.b.c.$index.d" 转为数组路径
  const keyChain = path.split('.');
  const res: ResItem[] = [];
  const current: ResItem = {
    paths: [],
    index: 0
  };
  recursiveGeneratePath({ res, current, keyChain, config });
  return res.map((item) => ({
    index: item.index,
    path: item.paths.join('.')
  }));
};
