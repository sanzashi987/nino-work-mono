import { AnnotationEndpointType } from '../proto-service/types';

const duration = {
  name: 'duration',
  type: 'number',
  default: 0,
  description: '动画持续时间'
} as const;

export const hide: AnnotationEndpointType = {
  fields: {
    name: 'hide',
    type: 'object',
    description: '只接受示例数据类型, 或者不传递数据',
    children: {
      type: {
        name: 'type',
        type: 'string',
        default: 'fadeOut',
        description: '出场动画类'
      },
      duration
    }
  }
};

export const unmount: AnnotationEndpointType = {
  ...hide,
  description: '卸载会禁用除"显示","隐藏"以外的所有动作事件'
};

export const display: AnnotationEndpointType = {
  fields: {
    name: 'display',
    type: 'object',
    description: '只接受示例数据类型, 或者不传递数据',
    children: {
      type: {
        name: 'type',
        type: 'string',
        default: 'fadeIn',
        description: '入场动画类'
      },
      duration
    }
  }
};

export const toggleVisible: AnnotationEndpointType = {
  fields: {
    description: '只接受示例数据类型, 或者不传递数据',
    name: 'toggleVisible',
    type: 'object',
    children: {
      display: display.fields!,
      hide: hide.fields!
    }
  }
};

export const init: AnnotationEndpointType = {
  fields: {
    name: '实例',
    description: "如逻辑节点的'设置实例'动作\"",
    type: 'pair',
    pairType: 'instance'
  },
  describer: { pairType: 'instance' }
};
