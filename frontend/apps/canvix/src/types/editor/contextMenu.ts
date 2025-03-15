import type { ReactNode } from 'react';
import type { Emitter } from 'mitt';

export type MenuContextValueType = (
  e: React.MouseEvent,
  clickPayload?: any,
  opts?: MenuItemType[],
) => void;

export type MenuItemTypeLegacy = {
  uncheckedLabel: string;
  uncheckedIcon: ReactNode;
  checkedIcon?: ReactNode;
  checkedLabel?: string;
  disabled?: boolean | ((args: any) => boolean);
  checked?: boolean | ((args: any) => boolean);
  onClick: (activeIds?: string[], eventHub?: Emitter<any>) => any;
};

type ContextCallbackWithEvent<ReturnType = void> = (
  e: React.MouseEvent,
  invokedBy: string,
) => ReturnType;
type ContextCallback<ReturnType = void> = (invokedBy: string) => ReturnType;

export type MenuItemType = {
  name: string | ContextCallback<string>;
  icon?: ReactNode | ContextCallback<ReactNode>;
  disabled?: boolean | ContextCallback<boolean>;
  shortcutNode: ReactNode | null;
  callback?: ContextCallbackWithEvent;
};

export type ContextMenuProps = {
  children?: ReactNode;
};

export type ContextMenuStates = {
  display: boolean;
  position: { x: number; y: number };
  // operating: string;
};
