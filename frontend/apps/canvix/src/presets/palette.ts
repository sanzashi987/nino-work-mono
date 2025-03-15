import { PaletteConfigType } from '@/types';

export const defaultPalette: PaletteConfigType = {
  palette: ['#FE672A', '#FF9500', '#FFBF00', '#B6B5B7', '#D5D2D8', '#DEDEDE', '#4C4D83'],
  bgColor: 'rgba(29,30,33,0.5)',
  textColor: '#FFFFFF',
  axisColor: '#4C4E51',
  assistColor: 'rgba(125,125,129,1)'
};

export const SystemModePresets = {
  light: { id: 'light', icon: 'LightMode', name: '浅色模式' },
  dark: { id: 'dark', icon: 'DarkMode', name: '深色模式' }
};
