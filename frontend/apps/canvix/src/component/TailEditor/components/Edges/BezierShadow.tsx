import React, { FC } from 'react';
import { drawBezier, EdgeBasicProps } from 'tail-js';

const BezierShadow: FC<EdgeBasicProps> = (props) => {
  const d = drawBezier(props);
  return <path d={d} className="bezier-edge" strokeWidth={15} />;
};

export default BezierShadow;
