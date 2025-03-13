import { ComItemType } from '../com-config';

/** Layers */
export type LayerItem = {
  id: string;
  children?: LayerItem[];
  type: ComItemType; // | RWDType;
};

export type LayerList = LayerItem[];
