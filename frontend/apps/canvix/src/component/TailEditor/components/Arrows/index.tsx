import React, { FC } from 'react';
import { wrapAnchor, AnchorProps, Marker } from 'tail-js';
import styles from './index.module.scss';

const { 'arrow-inner': arrowclass } = styles;

const BasicArrow: FC<AnchorProps> = () => {
  return (
    <polyline
      className={arrowclass}
      stroke="var(--canvas-interaction-edge-color)"
      strokeWidth="2"
      points="-20,-7 -10,0 -20,7"
      fill="var(--canvas-interaction-edge-color)"
    />
  );
};

const HoverArrow: FC<AnchorProps> = () => {
  return (
    <polyline
      className={arrowclass}
      stroke="var(--canvas-interaction-edge-color)"
      strokeWidth="6"
      points="-20,-7 -10,0 -20,7"
      fill="var(--canvas-interaction-edge-color)"
    />
  );
};
const ActiveArrow: FC<AnchorProps> = () => {
  return (
    <polyline
      className={arrowclass}
      stroke="var(--canvas-interaction-active-color)"
      strokeWidth="6"
      points="-20,-7 -10,0 -20,7"
      fill="var(--canvas-interaction-active-color)"
    />
  );
};
const DisableArrow: FC<AnchorProps> = () => {
  return (
    <polyline
      className={arrowclass}
      stroke="var(--canvas-interaction-disable-color)"
      strokeWidth="6"
      points="-20,-7 -10,0 -20,7"
      fill="var(--canvas-interaction-disable-color)"
    />
  );
};

export const ArrowTemplates = {
  basic: wrapAnchor(BasicArrow),
  hover: wrapAnchor(HoverArrow),
  active: wrapAnchor(ActiveArrow),
  disable: wrapAnchor(DisableArrow),
};

export const Arrows: Marker[] = [
  { id: 'canvas-basic-arrow', type: 'basic' },
  { id: 'canvas-hover-arrow', type: 'hover' },
  { id: 'canvas-active-arrow', type: 'active' },
  { id: 'canvas-disable-arrow', type: 'disable' },
];
