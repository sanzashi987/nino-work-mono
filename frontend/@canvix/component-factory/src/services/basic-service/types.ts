type Default = {
  opacity: number;
};

export type Absolute = {
  x: number;
  y: number;
  w: number;
  h: number;
} & Default;

export type Flex = {
  display: 'flex';
  flexDirection: 'column' | 'row' | 'row-reverse' | 'column-reverse';
  flexWrap: 'nowrap' | 'wrap' | 'wrap-reverse';
  alignItems: 'normal' | 'center' | 'flex-start' | 'flex-end';
  justifyContent:
  | 'normal'
  | 'flex-start'
  | 'center'
  | 'space-between'
  | 'space-around'
  | 'space-evenly'
  | 'flex-end';
};

export type Box = {
  position: 'static';
};

export type AbsoluteBasic = {
  position: 'absolute';
  top: string;
  left: string;
  right: string;
  bottom: string;
};

export type TransformBasic = {
  x: number;
  y: number;
  z: number;
  rotate: number;
  rotateX: number;
  rotateY: number;
  rotateZ: number;
  scale: number;
  scaleX: number;
  scaleY: number;
  scaleZ: number;
};

export type BoxBasic = {
  height: string;
  width: string;
  marginTop: string;
  marginLeft: string;
  marginRight: string;
  marginBottom: string;
  paddingTop: string;
  paddingLeft: string;
  paddingRight: string;
  paddingBottom: string;
  maxHeight: string;
  minHeight: string;
  maxWidth: string;
  minWidth: string;
  backgroundColor: string;
  backgroundImage: string;
  flexGrow: number;
  flexShrink: number;
};

// export type BorderBasic = {};
// export type TextBasic = {
//   color:
// }

export type BoxShadowBasic = {
  offsetX: string;
  offsetY: string;
  blur: string;
  spread: string;
  color: string;
};

export type FilterBasic = object;

export type RWDCombination = Flex | Box;
