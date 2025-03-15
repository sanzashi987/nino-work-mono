import React from 'react';
import { noop } from '@/utils';
import { VariableContextValue, VariableDepot, createUseVariableSubscribe } from './TempValue';
import {
  SharedVariableMutator,
  StaticVariableMutator,
  TypedVariableMutatorCtor,
  UrlVariableMutator
} from './mutators';
import { GlobalVariableSource, LocalVariableSource } from '@/types';

const defaultContext: VariableContextValue = {
  join: noop,
  leave: noop,
  getVariable: noop,
  setVariable: noop
};

const GlobalVariableContext = React.createContext(defaultContext);
export class GlobalVariableDepot extends VariableDepot<GlobalVariableSource> {
  getContext() {
    return GlobalVariableContext;
  }

  getUsedMutators(): TypedVariableMutatorCtor<'static' | 'sqlite' | 'shared' | 'url'>[] {
    return [StaticVariableMutator, UrlVariableMutator, SharedVariableMutator];
  }
}
export const useGlobalVariable = createUseVariableSubscribe(GlobalVariableContext);

const LocalVariableContext = React.createContext(defaultContext);
export class LocalVariableDepot extends VariableDepot<LocalVariableSource> {
  getUsedMutators(): TypedVariableMutatorCtor<'static'>[] {
    return [StaticVariableMutator];
  }

  getContext(): React.Context<VariableContextValue> {
    return LocalVariableContext;
  }
}
export const useLocalVariable = createUseVariableSubscribe(LocalVariableContext);

export const useVariable = (variableId: string, global = false) =>
  // eslint-disable-next-line
   (global ? useGlobalVariable(variableId) : useLocalVariable(variableId));
