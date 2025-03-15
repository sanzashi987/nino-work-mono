import { BasicPanel } from '@/component/Controller';
import { LocalVariableDepot } from '@/component/VariableDepot';
import { PanelMinimalProps, PanelState, ResponsivePanelUtils, PanelLogicalUtilsType } from '@/types';

export abstract class PanelMinimal extends BasicPanel<
PanelMinimalProps,
PanelState,
ResponsivePanelUtils,
PanelLogicalUtilsType
> {
  declare depotRef: React.RefObject<LocalVariableDepot>;

  declare render: () => React.ReactNode;
}
