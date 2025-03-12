import React, { createContext, FC, useEffect } from 'react';
import keyboardJS from 'keyboardjs';
import { noop } from '@/utils';
import { mergeCallback, ShortcutItem } from './utils';

declare module 'keyboardjs' {
  export const KeyCombo: {
    parseComboStr:(str:string) =>string[][];
  };

  interface Ctor {
    new():typeof keyboardJS
  }
  export const Keyboard: Ctor;
}

type ShortcutContextType = {
  getHoldingStatusByKey(key: string): boolean;
  bindShortcut(config: ShortcutItem): () => void;
  bindHolding(key: string): () => void;
};

const defaultShortCut = {
  getHoldingStatusByKey: () => false,
  bindShortcut: () => noop,
  bindHolding: () => noop
};

export const ShortcutContext = createContext<ShortcutContextType>(defaultShortCut);

type ShortcutsProps = {
  scoped: string;
  children?: React.ReactNode;
};

const keyboardStatic = new keyboardJS.Keyboard();

const createBinderFromCore = (shortcutCore: typeof keyboardStatic) => {
  const keyHolding: Record<string, boolean> = {};
  const toBeUnbind: Record<string, typeof noop> = {};
  const getHoldingStatusByKey = (key: string) => !!keyHolding[key];

  const bindShortcut = ({ id, scoped, keyCombo, pressed, released, disabled }: ShortcutItem) => {
    const pressedExactApplied = mergeCallback(keyCombo, disabled, pressed);
    const ctx = scoped ?? 'global';
    shortcutCore.withContext(ctx, () => {
      shortcutCore.bind(keyCombo, pressedExactApplied, released);
    });
    const unbind = () => {
      delete toBeUnbind[id];
      shortcutCore.withContext(ctx, () => {
        shortcutCore.unbind(keyCombo, pressedExactApplied, released);
      });
    };

    toBeUnbind[id]?.();
    toBeUnbind[id] = unbind;

    return toBeUnbind[id];
  };
  const bindHolding = (key: string) => bindShortcut({
    id: `_reserved_holding_id_${key}`,
    keyCombo: key,
    pressed: () => {
      keyHolding[key] = true;
    },
    released: () => {
      keyHolding[key] = false;
    }
  });

  const destructor = () => {
    Object.values(toBeUnbind).forEach((fun) => fun());
  };

  return {
    bindShortcut,
    bindHolding,
    getHoldingStatusByKey,
    destructor
  };
};

const createShortcutProvider = (holding: string[], shortcutList: ShortcutItem[]) => {
  const dynamicBinders = createBinderFromCore(keyboardJS);
  const staticBinders = createBinderFromCore(keyboardStatic);

  const ctx: ShortcutContextType = {
    getHoldingStatusByKey: dynamicBinders.getHoldingStatusByKey,
    bindShortcut: (config) => {
      if (config.static) return noop;
      return dynamicBinders.bindShortcut(config);
    },
    bindHolding: dynamicBinders.bindHolding
  };

  const Shortcuts: FC<ShortcutsProps> = ({ scoped, children }) => {
    useEffect(() => {
      keyboardJS.setContext(scoped);
    }, [scoped]);

    useEffect(() => {
      const { bindHolding } = ctx;
      holding.forEach((key) => {
        bindHolding(key);
      });

      // only register the static features
      shortcutList.filter((config) => config.static).forEach(staticBinders.bindShortcut);

      return () => {
        dynamicBinders.destructor();
        staticBinders.destructor();
      };
    }, []);

    return <ShortcutContext.Provider value={ctx}>{children}</ShortcutContext.Provider>;
  };

  return { contextInterface: ctx, Shortcuts };
};

export default createShortcutProvider;
export { keyboardJS, keyboardStatic };
export type { ShortcutItem } from './utils';
