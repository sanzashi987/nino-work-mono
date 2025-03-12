import type { ComInfo } from '../com-config';
import type { FileType } from '../services';
import type { DeltaKey } from './meta';

/** 0 stands for system built-in palettes, 1 stands for user created palette */
type PaletteFlag = 0 | 1;

export type PaletteConfigType = {
  /**
   * 背景颜色，主要用于面板背景色
   */
  bgColor: string;
  /**
   * 文本颜色
   */
  textColor: string;
  /**
   * 坐标轴颜色
   */
  axisColor: string;
  /**
   * 辅助信息颜色
   */
  assistColor: string;
  /**
   * 调色板颜色，主要用于系列颜色
   */
  palette: string[];
};

export type PaletteItemType = {
  name: string;
  palette: PaletteConfigType;
  id: /* string |  */ number;
  flag: PaletteFlag;
};

export type PaletteApplyType = {
  [panelId: string]: {
    panel: Record<string, any>;
    components: {
      [comId: string]: Record<string, any>;
    };
  };
};

export type PaletteDeltaKey = {
  [panelId: string]: {
    panel: DeltaKey;
    components: {
      [comId: string]: DeltaKey;
    };
  };
};

export type PalettePreviewType = PaletteApplyType;

export type PaletteApplyOutput = {
  preview: PalettePreviewType;
  apply: PaletteApplyType;
};

export type UpdatePaletteComs = {
  [panelId: string]: {
    panel: {
      /** 面板basic */
      basic?: Record<string, any>;
    };
    components: {
      [comId: string]: {
        /** 组件attr */
        attr: Record<string, any>;
        /** 组件basic */
        basic?: Record<string, any>;
        type: FileType;
        com: ComInfo;
      };
    };
  };
};

type PaletteItemFromResponse = Record<string, any> & {
  theme: string;
  userIdentify: string;
  deleted: number;
  /**
   * format `YYYY-MM-DD hh:mm:ss`
   */
  createTime: string;
  updateTime: string;
} & Omit<PaletteItemType, 'palette'>;

export type PaletteListFromResponse = PaletteItemFromResponse[];

export type PaletteItemList = PaletteItemType[];

export type SinglePaletteConfig = Record<string, string>;
export type PaletteConfigs = Record<string, SinglePaletteConfig>;

export type UnsafeRequestCreatePalettePayload = RequestCreatePalettePayload & {
  flag: PaletteFlag;
};

export type RequestCreatePalettePayload = {
  name: string;
  theme: string;
  // flag: PaletteFlag;
};

export type RequestUpdatePalettePayload = {
  id: number;
} & Partial<RequestCreatePalettePayload>;
