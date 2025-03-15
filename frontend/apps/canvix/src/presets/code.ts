import type { CanvixResponse } from '@/types';

export enum GroupCode {
  ALLGROUP = '',
  NOGROUP = '-1',
}

export const NO_GROUPED = { name: '未分组', id: -1, code: GroupCode.NOGROUP };
export const ALL_GROUPED = { name: '全部', id: -99, code: GroupCode.ALLGROUP };

const prefix = 'cVx';

/** code编码前缀 */
export enum CanvixCodePrefix {
  /** 大屏/项目 */
  project = `${prefix}A`,
  /** 模块 */
  block = `${prefix}B`,
  /** 设计资产 */
  design = `${prefix}C`,
  /** 字体资产 */
  font = `${prefix}D`,
  /** 自定义组件 */
  component = `${prefix}E`,
  /** 数据源 */
  data = `${prefix}F`,
  /** 空间 */
  workspace = `${prefix}P`,
}

type PrefixType = typeof CanvixCodePrefix;
const canvsaCodeInterpretation = Object.fromEntries(
  Object.entries(CanvixCodePrefix).map(([k, v]) => [v, k])
) as unknown as {
  -readonly [K in keyof PrefixType as PrefixType[K]]: K;
};

const prefixTag = Object.values(CanvixCodePrefix)
  .map((e) => e.at(-1))
  .join('');

const localPrefix = '_CvX';

export enum CanvasLocalPrefix {
  variable = `${localPrefix}V`,
  function = `${localPrefix}F`,
}

/** code编码字符长度 */
export const CanvixCodeLength = 14;

/**
 * @param {string} id the string to be tested
 * @return {keyof CanvixCodePrefix | null} `res` canvas code category or null
 */
export const lookupCanvixCode = (id: string) => {
  const res = id.match(new RegExp(`^${prefix}[${prefixTag}]`));
  if (!res) return null;
  return ((canvsaCodeInterpretation as any)[res[0]] as keyof PrefixType) ?? null;
};

export const isCanvixCode = (type: keyof typeof CanvixCodePrefix, code: string) => {
  if (!code) return false;
  const pfx = CanvixCodePrefix[type];
  return code.startsWith(pfx) && code.length === CanvixCodeLength;
};

export const RES_NULL: CanvixResponse<null> = {
  resultCode: 0,
  resultMessage: '',
  data: null
};

export const defaultApiRequest = () => Promise.resolve(RES_NULL);
export const imageAccept = 'image/*';
export const audioAccept = 'audio/wav,audio/mpeg';
export const videoAccept = 'video/mp4';
export const mediaAccept = `${imageAccept},${audioAccept},${videoAccept}`;
