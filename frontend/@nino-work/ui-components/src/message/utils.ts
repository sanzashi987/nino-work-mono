/* eslint-disable import/prefer-default-export */
import { MessageContent, MessageConfig, BasicConfigWithType } from './type';
import type { MessageKey } from './Message';

const isObject = (input: any) => Object.prototype.toString.call(input) === '[object Object]';
const isError = (input: any) => Object.prototype.toString.call(input) === '[object Error]';
const isCancel = (input: any) => false;

/**
 * 格式化消息内容
 * @description 返回false时，不显示消息内容
 * @param config 消息配置项
 * @param type 消息类型
 * @param duration 时长
 * @returns
 */
export const formatConfig = (
  config: MessageContent | MessageConfig,
  type: MessageKey,
  duration?: number
): BasicConfigWithType | false => {
  // config不存在或者CanceledError时，返回false
  if (!config || isCancel(config)) return false;

  // Error，content字段取resultMessage或者message字段
  if (isError(config)) {
    return {
      type,
      duration,
      content: (config as any)?.resultMessage || (config as any)?.message,
    };
  }

  // 第一个配置项为普通对象
  if (isObject(config)) {
    const { content, ...other } = config as MessageConfig;
    if (!content || isCancel(content)) return false;
    const nextContent = isError(content)
      ? (content as any)?.resultMessage || (content as any)?.message
      : content;
    return {
      ...other,
      content: nextContent as React.ReactNode,
      type,
    };
  }

  // config为ReactNode
  return {
    type,
    duration,
    content: config as React.ReactNode,
  };
};
