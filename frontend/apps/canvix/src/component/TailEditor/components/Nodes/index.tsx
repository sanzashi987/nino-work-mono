import { BlankLoadable } from '@canvas/utilities';
import type { InteractionNodeTypeRuntime } from '@canvas/event-core';
import type { Node } from 'tail-js';

// const NodeTemplates: NodeTemplatesType = {
//   logical: { default: LogicalTemplate },
//   normal: { default: NormalTemplate },
// };

export const LazyNormal = BlankLoadable(() => import('./NormalNode'));
export const LazyLogical = BlankLoadable(() => import('./LogicalNode'));
export const LazyPanel = BlankLoadable(() => import('./PanelNode'));

export const NodeTemplatePicker = ({ type }: Node<InteractionNodeTypeRuntime>) => {
  // return [['logical', 'refPanel'].includes(type) ? type : 'normal', 'default'] as [string, string];
  return [type, 'default'] as [string, string];
};

// export { LazyNormal as NormalNode, LazyLogical as LogicalNode };

export { default as LogicalNode } from './LogicalNode';
export { default as NormalNode } from './NormalNode';
export { default as RefNode } from './RefNode';
