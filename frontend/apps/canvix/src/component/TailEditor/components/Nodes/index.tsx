import { InteractionNodeTypeRuntime } from '@canvix/event-core';
import type { Node } from 'tail-js';

export function NodeTemplatePicker({ type }: Node<InteractionNodeTypeRuntime>) {
  return [type, 'default'] as [string, string];
}

export { default as LogicalNode } from './LogicalNode';
export { default as NormalNode } from './NormalNode';
