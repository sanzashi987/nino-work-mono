import tinycolor from 'tinycolor2';
import gradient from '@canvas/gradient-parser';

type VerifyResult = {
  /**
   * 基础格式校验，适用于输入框，返回false或者合法颜色值
   */
  basic: false | string;
  /**
   * 严格模式校验，适用于颜色选择器弹窗，返回false或者合法颜色值
   */
  strict: false | string;
};

const getColorsByStops = (stops: Array<[string, string]>) => {
  let res = '';
  stops.forEach((item) => {
    if (!tinycolor(item[0]).isValid()) throw new Error('格式不支持');
    // eslint-disable-next-line @typescript-eslint/naming-convention
    const _color = tinycolor(item[0]);
    const color = _color.getFormat() !== 'rgb' ? _color.toRgbString() : item[0];
    const percent = Math.round(parseFloat(item[1]));
    res += `, ${color} ${percent}%`;
  });
  return res;
};

const getLinearAngle = (gradientDefinition: string, linearAngle: number) => {
  let res = linearAngle || 0;
  if (gradientDefinition.includes('N/A')) {
    // deg缺省
    res = 180;
  }
  if (gradientDefinition.includes('turn')) {
    // turn 转 deg
    res = Math.round(parseFloat(gradientDefinition) * 360);
  }
  return `${res}deg`;
};

const formatRadialGradientColor = (str: string) => {
  const { stops, gradientDefinition } = gradient.parse(str, true);
  if (gradientDefinition !== 'circle') throw new Error('径向渐变格式不支持');
  let res = 'radial-gradient(circle';
  res += getColorsByStops(stops);
  res += ')';
  return res;
};

const formatLinearGradientColor = (str: string) => {
  const { stops, linearAngle, gradientDefinition } = gradient.parse(str, true);
  let res = 'linear-gradient(';
  res += getLinearAngle(gradientDefinition, linearAngle);
  res += getColorsByStops(stops);
  res += ')';
  return res;
};

const formatGradientColor = (
  str: string,
  regexp: RegExp,
  cb: (str: string) => string
): VerifyResult => {
  const isValid = regexp.test(str);
  if (!isValid || !CSS.supports('background-image', str)) {
    return {
      basic: false,
      strict: false
    };
  }
  //  能转成标准格式时优先转成标准格式
  //  非严格模式下，无法转成标准格式时，返回原数据
  try {
    const temp = cb(str);
    return {
      basic: temp,
      strict: temp
    };
  } catch (err) {
    return {
      basic: str,
      strict: false
    };
  }
};

const verifyCssSupports = (str: string, regexp: RegExp): VerifyResult => {
  const isValid = regexp.test(str);
  return {
    basic: isValid && CSS.supports('background-image', str) ? str : false,
    strict: false
  };
};

/** ----------------------------------------------- */

/**
 * 是否是合法的纯色
 */
export const isSolidColor = (str: string): VerifyResult => {
  const res = tinycolor(str).isValid() ? tinycolor(str).toString() : false;
  return {
    basic: res,
    strict: res
  };
};

/**
 * 是否是合法的线性渐变颜色
 */
const isLinearGradientColor = (str: string): VerifyResult => formatGradientColor(str, /^linear-gradient[(].*[)]$/i, formatLinearGradientColor);

/**
 * 是否是合法的径向渐变颜色
 * @description 标准格式仅支持circle，颜色类型为rgba
 */
const isRadialGradientColor = (str: string): VerifyResult => formatGradientColor(str, /^radial-gradient[(].*[)]$/i, formatRadialGradientColor);

/**
 * 是否是合法的重复渐变色
 * @description 支持重复线性渐变与重复径向渐变
 */
const isRepeatingGradient = (str: string): VerifyResult => verifyCssSupports(str, /^repeating-.*-gradient[(].*[)]$/i);

/**
 * 是否是合法的锥形渐变
 */
const isConicGradient = (str: string): VerifyResult => verifyCssSupports(str, /^conic-gradient[(].*[)]$/i);

/**
 * 是否是合法的渐变颜色（线性/径向/重复/锥形）
 */
export const isGradientColor = (str: string): VerifyResult => {
  const list: VerifyResult[] = [
    isLinearGradientColor(str),
    isRadialGradientColor(str),
    isRepeatingGradient(str),
    isConicGradient(str)
  ];

  const res: VerifyResult = {
    basic: false,
    strict: false
  };

  list.forEach((item) => {
    res.basic = res.basic || item.basic;
    res.strict = res.strict || item.strict;
  });
  return res;
};

export { tinycolor };
