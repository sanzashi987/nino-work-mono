/* eslint-disable no-nested-ternary */
import React, { FC } from 'react';
import {
  EdgeProps,
  drawBezier,
  EdgeBasicProps,
  ConnectingEdgeProps
} from 'tail-js';
import styles from './BezierEdge.module.scss';

const { 'bezier-edge': BezierClass, 'bezier-connecting-edge': BezierConnectingClass } = styles;

const BezierEdge: FC<EdgeProps> = ({ selected, hovered, edge, ...xy }) => {
  const { disable } = edge;
  const stroke = disable
    ? 'var(--canvix-interaction-disable-color)'
    : selected
      ? 'var(--canvix-interaction-active-color)'
      : 'var(--canvix-interaction-edge-color)';
  const width = selected || hovered ? 3 : 1;
  const path = drawBezier(xy);

  const markerEnd = selected
    ? 'canvas-active-arrow'
    : disable
      ? 'canvas-disable-arrow'
      : hovered
        ? 'canvas-hover-arrow'
        : 'canvas-basic-arrow';

  return (
    <path
      className={BezierClass}
      // style={{ transition: 'stroke-width 0.2s linear' }}
      d={path}
      markerEnd={`url(#${markerEnd})`}
      // markerStart={markerStart}
      stroke={stroke}
      strokeWidth={width}
    />
  );
};

export const BezierBasicEdge: FC<EdgeBasicProps> = (props) => (
  <path
    className={BezierClass}
    // style={{ transition: 'stroke-width 0.2s linear' }}
    d={drawBezier(props)}
    markerEnd="url(#canvas-basic-arrow)"
    stroke="var(--canvix-interaction-edge-color)"
    strokeWidth="1"
    fill="transparent"
  />
);

export const BezierConnectingEdge: FC<ConnectingEdgeProps> = ({ pairedStatus, ...props }) => {
  const appendClass = pairedStatus ?? '';
  return (
    <path
      className={`${BezierConnectingClass} ${appendClass}`}
      // style={{ transition: 'stroke-width 0.2s linear' }}
      d={drawBezier(props)}
      markerEnd="url(#canvas-basic-arrow)"
      stroke="var(--canvix-interaction-edge-color)"
      strokeWidth={appendClass ? 3 : 1}
      fill="transparent"
    />
  );
};

export default BezierEdge;
