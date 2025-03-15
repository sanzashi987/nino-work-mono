import { isCanvixCode } from '@/presets/code';
import { GenerateType } from '@/types';

const generate: GenerateType<any> = {
  getPaths(params) {
    const { config, createObject, key, result } = params;

    createObject(result, key, {
      configs: [{
        type: 'parse',
        value: config.type
      }],
      component: config.type
    });
  },

  parseValue(params) {
    const { value, config } = params;
    // 图片id时拼接上完整url，渐变色时不拼接
    return isCanvixCode('design', value)
      ? `url(${config.getAssetsUrl?.(value)})`
      : value;
  }

};
export default generate;
