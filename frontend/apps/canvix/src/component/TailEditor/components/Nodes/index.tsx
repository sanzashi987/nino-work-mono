import { InteractionNodeTypeRuntime } from '@canvix/event-core';
import { BlankLoadable } from '@canvix/utils';
import type { Node } from 'tail-js';

export const LazyNormal = BlankLoadable(() => import('./NormalNode'));
export const LazyLogical = BlankLoadable(() => import('./LogicalNode'));
export const LazyPanel = BlankLoadable(() => import('./PanelNode'));

export const NodeTemplatePicker = ({ type }: Node<InteractionNodeTypeRuntime>) => [type, 'default'] as [string, string];
// return [['logical', 'refPanel'].includes(type) ? type : 'normal', 'default'] as [string, string];

// export { LazyNormal as NormalNode, LazyLogical as LogicalNode };

export { default as LogicalNode } from './LogicalNode';
export { default as NormalNode } from './NormalNode';
