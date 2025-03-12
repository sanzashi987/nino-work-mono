import { PaletteConfigType } from './palette';

export type ThemeType = {
  id: string;
  name: string;
  icon: string;
  default?: boolean;
  palette: null | PaletteConfigType;
};

export type ThemeMetaType = {
  followSystem: boolean;
  configs: ThemeType[];
};

type BreakpointValue =
  | {
    id: string;
    lower: number;
  }
  | { id: 'default'; lower: 0 };

export type BreakpointMetaType = BreakpointValue[];

export type Dimension = {
  width: number;
  height: number;
};

export type Coordinates = {
  x: number;
  y: number;
};
