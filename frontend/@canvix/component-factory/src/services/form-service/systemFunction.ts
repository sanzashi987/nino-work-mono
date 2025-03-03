import { AnnotationEndpointType, FieldsType } from '../proto-service/types';

const systemFunction: Record<string, { label: string; content: string }> = {
  required: {
    label: '必填',
    content: `const type = Object.prototype.toString.call(value).slice(8, - 1).toLocaleLowerCase();\r\nif (type === "array") return value && value.length > 0;\r\nif (type === "number") return value || value === 0;\r\nreturn value;`,
  },
  number: {
    label: '数字',
    content: 'return /^[0-9]*$/.test(value)',
  },
};

export const verifyType: AnnotationEndpointType = {
  fields: {
    type: 'null',
    name: 'verify',
  },
};

export const verifySuccessType: AnnotationEndpointType = {
  fields: {
    type: 'object',
    name: 'success',
    children: {
      type: {
        type: 'string',
        name: 'type',
        description: '状态',
        default: 'success',
      },
      value: {
        type: 'any',
        name: 'value',
        description: '表单值',
      },
    },
  },
};
export const verifyFailType: AnnotationEndpointType = {
  fields: {
    type: 'object',
    name: 'fail',

    children: {
      type: {
        type: 'string',
        name: 'type',
        description: '状态',
        default: 'fail',
      },
      value: {
        type: 'any',
        name: 'value',
        description: '校验提示信息',
      },
    },
  },
};

export const formValueType: AnnotationEndpointType = {
  fields: {
    type: 'object',
    name: 'value',
    children: {
      value: {
        type: 'any',
        name: 'value',
        description: '数据结构与控件保持一致',
      },
    },
  },
};

export default systemFunction;
