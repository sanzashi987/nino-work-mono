import React, { useContext, useEffect, useImperativeHandle, useState } from 'react';
import type { GetVariableType, TempValueHandles } from '@/types';
import { TypedVariableMutator, TypedVariableMutatorCtor } from './mutators';

export const TempValueContext = React.createContext<Record<string, any>>({});

const TempValue: React.ForwardRefRenderFunction<
TempValueHandles,
{
  initValue?: Record<string, any>;
  children: React.ReactNode;
}
> = ({ children, initValue }, ref) => {
  const [tempValue, setValue] = useState<Record<string, any>>(initValue ?? {});

  useImperativeHandle(ref, () => ({
    setValue,
    getValue: () => tempValue
  }));

  return <TempValueContext.Provider value={tempValue}>{children}</TempValueContext.Provider>;
};

export default React.forwardRef(TempValue);

type Updater = (x: any) => void;

export type VariableContextValue = {
  join(variableId: string, updater: Updater): void;
  leave(variableId: string, updater: Updater): void;
  getVariable(variableId: string): any;
  setVariable(variableId: string, value: any): void;
};

type VariableContextType = React.Context<VariableContextValue>;

type VariableDepotProps<T> = {
  configs: GetVariableType<T>[];
  children?: React.ReactNode;
};

export abstract class VariableDepot<T extends string> extends React.Component<
VariableDepotProps<T>
> {
  variableMap: Record<string, GetVariableType<T>> = {};

  mutators: { [Key in T]: TypedVariableMutator<T> };

  subscription = new Map<string, Set<(x: any) => void>>();

  constructor(props: VariableDepotProps<T>) {
    super(props);
    const mutators = this.getUsedMutators();
    this.mutators = Object.fromEntries(mutators.map((E) => [E.type, new E()])) as any;
    this.makeConfigMap(props.configs);
  }

  abstract getUsedMutators(): TypedVariableMutatorCtor<T>[];

  abstract getContext(): VariableContextType;

  makeConfigMap(configs: GetVariableType<T>[]) {
    if (!configs) return;
    configs.forEach((config) => {
      if (this.mutators[config.source]) {
        this.variableMap[config.id] = config;
      }
    });
  }

  setVariable = (variableId: string, value: any) => {
    const conf = this.variableMap[variableId];
    if (!conf) return;
    const mutator = this.mutators[conf.source];
    const last = mutator.getValue(conf);
    if (last === value) return;
    mutator.setValue(conf, value);
    this.subscription.get(conf.id)?.forEach((cb) => {
      cb(value);
    });
  };

  getVariable = (variableId: string) => {
    const conf = this.variableMap[variableId];
    if (!conf) return undefined;
    const mutator = this.mutators[conf.source];
    return mutator.getValue(conf);
  };

  join = (variableId: string, updater: Updater) => {
    if (!this.subscription.get(variableId)) {
      this.subscription.set(variableId, new Set());
    }
    this.subscription.get(variableId)?.add(updater);
  };

  leave = (variableId: string, updater: Updater) => {
    this.subscription.get(variableId)?.delete(updater);
  };

  contextImpl = {
    setVariable: this.setVariable,
    getVariable: this.getVariable,
    join: this.join,
    leave: this.leave
  };

  componentDidUpdate(prevProps: Readonly<VariableDepotProps<T>>): void {
    if (prevProps.configs !== this.props.configs) {
      this.makeConfigMap(this.props.configs);
    }
  }

  render(): React.ReactNode {
    const Context = this.getContext();
    return <Context.Provider value={this.contextImpl}>{this.props.children}</Context.Provider>;
  }
}

export const createUseVariableSubscribe = (context: VariableContextType) => (variableId: string) => {
  const { join, leave, getVariable, setVariable } = useContext(context);
  const [variable, setInternalVariable] = useState(() => getVariable(variableId));

  useEffect(() => {
    const receiveUpdate = (nextValue: any) => {
      setInternalVariable(() => nextValue);
    };
    join(variableId, receiveUpdate);
    return () => leave(variableId, receiveUpdate);
  }, []);

  return [variable, setVariable] as const;
};
