import { isEmpty } from '@canvix/utils';
import ProtoService from '../proto-service';
import { handler, service } from '../proto-service/annotations';
import { AnnotationEndpointType } from '../proto-service/types';
// import { ExportAction, ProtoService, ServiceInit } from '../proto-service';
// import type { ServiceDefaultProps, GetIdentifierType } from '../proto-service';

// const BasicValueType = {
//   name: '{',
//   type: 'object',
//   description: '基础数据',
//   children: {
//     x: { name: 'x', type: 'number', optional: true, description: '横坐标' },
//     y: { name: 'y', type: 'number', optional: true, description: '纵坐标' },
//     w: { name: 'w', type: 'number', optional: true, description: '宽' },
//     h: { name: 'h', type: 'number', optional: true, description: '长' },
//     opacity: { name: 'opacity', type: 'number', optional: true, description: '透明度' },
//     deg: { name: 'deg', type: 'number', optional: true, description: '角度' },
//   },
// } as const;

// const mergeBasic = (
//   target: Record<string, any>,
//   source: Record<string, any>,
// ): Record<string, any> => {
//   return produce(target, (draft: Record<string, any>) => {
//     merge(draft, source);
//   });
// };

const BasicValueType: AnnotationEndpointType = {
  description: '组件容器的html属性',
  fields: {
    type: 'object',
    description: '组件容器的html属性',
    default: '$basic',
    name: 'basic'
  }
};

const SPRING_CONFIG_KEYS: Record<string, boolean> = {
  from: true,
  to: true,
  loop: true,
  delay: true,
  immediate: true,
  reset: true,
  reverse: true,
  pause: true,
  cancel: true,
  ref: true,
  config: true,
  event: true
};

@service('basic', 'basic')
class BasicService extends ProtoService {
  @handler('设置基础属性', BasicValueType)
    setBasic = (basicValue: Record<string, any>) => {
    // detect new keys
      if (
      // isEmpty判断用于取消主题预览
        isEmpty(basicValue)
      // only the no-exist, and not spring reserved key inserted will trigger react-setState
      // otherwise, only use spring start to apply the animation
      || Object.keys(basicValue).some((key) => !SPRING_CONFIG_KEYS[key] && !this.props.config[key])
      ) {
        this.props.setState((prev: any) => ({
          ...prev,
          config: {
            ...prev.config,
            basic: { ...prev.config.basic, ...basicValue }
          }
        }));
      } else {
        this.props.transitionRef.current.start({ ...basicValue });
      }
    };
}

export default BasicService;
