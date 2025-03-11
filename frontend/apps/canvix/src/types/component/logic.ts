export type Operator = '$=' | '$!=' | '$>' | '$<' | '$>=' | '$<=' | '$in' | '$nin';
export type LogicalType = '$and' | '$or';
type ConditionItem = [string, Operator, any];
export type ShowInPanel = {
  conditions: ConditionItem[];
  logicalType?: LogicalType;
};
export type OptionType<T> = {
  label: string;
  value: T;
  src?: string;
  disabled?: boolean;
} & Record<string, any>;
