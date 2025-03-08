/* eslint-disable @typescript-eslint/no-namespace */
/* eslint-disable no-restricted-syntax */
import { isMac } from '@canvix/utils';
import keyboard from 'keyboardjs';

type KeyEventCallback = (e?: keyboard.KeyEvent) => void;

export type ShortcutItem = {
  id: string;
  keyCombo: string | string[];
  pressed: KeyEventCallback;
  released?: KeyEventCallback;
  scoped?: string;
  /**
   * if given a function, it will be merged to the shortcut callback
   */
  disabled?: (() => boolean | null) | boolean | null;
  /**
   * the static item will only be registered to the shorcut durting the initialization
   * pass the statics to the `createShortcutProvider`, if the static item is passed to
   * the shortcuts core after init, it will be ignored.
   * */
  static?: boolean;
};

const modifiers:Record<string, string> = {
  ctrl: 'ctrl',
  option: 'alt',
  alt: 'alt',
  command: 'meta',
  shift: 'shift'
};

const standardModifiers = ['ctrl', 'alt', 'meta', 'shift'];

function getModifiers(str: string) {
  const res: Record<string, string> = {};
  keyboard.KeyCombo.parseComboStr(str)
    .reduce((memo: any, nextSubCombo: any) => memo.concat(nextSubCombo), [])
    .filter((e: string) => !!modifiers[e])
    .forEach((e: string) => {
      const projected = modifiers[e];
      res[projected] = projected;
    });
  return res;
}

function outsideBannedList(e: keyboard.KeyEvent, modifierList: string[][]) {
  for (const pos of modifierList) {
    const inside = pos.some((key) => !!(e as any)[`${key}Key`]);
    if (inside) return false;
  }
  return true;
}

function detectExact(modifierList: string[][], e?: keyboard.KeyEvent) {
  if (!e) return true;
  return outsideBannedList(e, modifierList);
}

function listModifiers(keyCombo: string | string[]) {
  const testArr = typeof keyCombo === 'string' ? [keyCombo] : keyCombo;
  return testArr.map((keyString) => {
    const privateModifiers = getModifiers(keyString);
    return standardModifiers.filter((modifier) => !privateModifiers[modifier]);
  });
}

export function mergeCallback(
  keyCombo: string | string[],
  disabled: ShortcutItem['disabled'],
  cb?: (e?: keyboard.KeyEvent) => void
) {
  const bannedList = listModifiers(keyCombo);

  if (typeof disabled === 'function') {
    return (e?: keyboard.KeyEvent) => {
      const disable = disabled();
      if (disable === true || disable === null) return;
      if (detectExact(bannedList, e)) {
        cb?.(e);
      }
    };
  }
  return (e?: keyboard.KeyEvent) => {
    if (detectExact(bannedList, e)) {
      cb?.(e);
    }
  };
}

type FeatureConfig = {
  id: string;
  shortcutWin?: string | string[];
  shortcutMac?: string | string[];
  callback(...p: any[]): any;
  scoped?: string;
};

const fieldKey = isMac ? 'shortcutMac' : 'shortcutWin';
export function reshapeFeatureToItem(input: FeatureConfig): ShortcutItem | null {
  const { id, scoped, callback, [fieldKey]: keyCombo } = input;
  if (!keyCombo) return null;
  return {
    id,
    scoped,
    pressed: callback,
    keyCombo
  };
}
