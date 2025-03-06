import BezierEdge, { BezierBasicEdge } from './BezierEdge';
import BezierShadow from './BezierShadow';

export { BezierShadow, BezierBasicEdge };

const BezierEdgePackage = {
  default: BezierEdge,
  shadow: BezierShadow
};

export const EdgeTemplate = { endpoint: BezierEdgePackage // compatible to the past data structure
};

export default BezierEdgePackage;

export { BezierConnectingEdge } from './BezierEdge';
