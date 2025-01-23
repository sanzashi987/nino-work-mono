import { OptionsObject, SnackbarKey } from 'notistack';
import type { ClickAwayListenerProps } from '@mui/material';
import type { CanceledError } from '@canvas/utilities';
import type { MessageKey } from './Message';

export type MessageContent = React.ReactNode | CanceledError<any> | Error;
export interface BasicConfig extends OptionsObject {
  content: React.ReactNode;
  /**
   * 消息显示时长，单位ms，默认为2000ms，传null时不自动消失
   */
  duration?: number | null;
  key?: number | string;
  /**
   * 点击目标元素（消息内容）外时抛出的事件
   */
  onClickAway?: ClickAwayListenerProps['onClickAway'];
}

interface Method {
  /**
   * @description content: 显示内容，类型为ReactNode | Error | CanceledError；入参也可以为配置项对象（需包含content字段）
   * @description content 为CanceledError时，不展示消息内容,如果想展示CanceledError内容，请通过ReactNode方式传参
   * @description content 为Error时，message内容为error.resultMessage || error.message
   * @description duration：消息显示时长，单位ms，默认为2000ms，传null时不自动消失
   * @returns SnackbarKey：消息唯一标识，可用于销毁当前消息
   */
  (content: MessageContent, duration?: number | null): SnackbarKey;
  /**
   * @description config：配置项，对象(需包含content字段,content代表显示内容，类型为ReactNode | Error | CanceledError）
   * @returns SnackbarKey：消息唯一标识，可用于销毁当前消息
   */
  (config: MessageConfig): SnackbarKey;
}
type MessageMethod = Record<MessageKey, Method>;

export type MessageConfig = BasicConfig & {
  content: MessageContent;
};

export interface BasicConfigWithType extends BasicConfig {
  type: MessageKey;
}

export interface MessageType extends MessageMethod {
  /**
   * 展示消息时，额外执行的函数，默认为()=>null
   */
  handleMessage: (message: BasicConfigWithType) => void;
  /**
   * 销毁指定消息,key不传时销毁所有
   */
  destroy: (key?: SnackbarKey) => void;
}
