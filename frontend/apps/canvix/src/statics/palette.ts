import { SystemModePresets } from '@/presets/palette';
import { ThemeMetaType } from '@/types';

export const DEFAULT_THEME = SystemModePresets.light.id;

export function isDefaultTheme(themeId: string) {
  return themeId === DEFAULT_THEME;
}

export function isSystemTheme(themeId: string) {
  return themeId in SystemModePresets;
}

export function canFollowSystem(themes: ThemeMetaType) {
  const sysTheme = new Set(
    themes.configs.filter((theme) => isSystemTheme(theme.id)).map((theme) => theme.id)
  );

  return sysTheme.has(SystemModePresets.light.id) && sysTheme.has(SystemModePresets.dark.id);
}
