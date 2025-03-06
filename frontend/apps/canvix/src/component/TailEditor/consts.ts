const duration = {
  name: 'duration',
  type: 'number',
  default: 0,
  description: '动画持续时间',
} as const;

const hide = {
  name: 'hide',
  type: 'object',
  description: '只接受示例数据类型, 或者不传递数据',
  children: {
    type: {
      name: 'type',
      type: 'string',
      default: 'fadeOut',
      description: '出场动画类',
    },
    duration,
  },
} as const;

const display = {
  name: 'display',
  type: 'object',
  description: '只接受示例数据类型, 或者不传递数据',
  children: {
    type: {
      name: 'type',
      type: 'string',
      default: 'fadeIn',
      description: '入场动画类',
    },
    duration,
  },
} as const;

export const comReservedTargets = [
  {
    name: '显示',
    id: 'instance.display',
    fields: display,
  },
  {
    name: '隐藏',
    id: 'instance.hide',
    fields: hide,
  },
  {
    name: '切换显隐',
    id: 'instance.toggleVisible',
    fields: {
      description: '只接受示例数据类型, 或者不传递数据',
      name: 'toggleVisible',
      type: 'object',
      children: {
        display,
        hide,
      },
    },
  },
];
