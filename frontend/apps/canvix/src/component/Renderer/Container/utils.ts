import { CSSProperties } from 'react';

const flexKey = ['flexDirection', 'flexWrap', 'alignItems', 'justifyContent'];

// export function filterBasic<T extends object>(basic: T) {
//   return Object.keys(basic).reduce((l, e) => {
//     //@ts-ignore
//     if (basic[e]) l[e] = basic[e];
//     return l;
//   }, {}) as NonNullable<T>;
// }

// export function mergeCssStyle(...args: CSSProperties[]) {
//   return args.reduce<CSSProperties>((last, css) => {
//     const { transform, transition, ...other } = css;
//     if (transform) {
//       last.transform = last.transform ?? '' + ' ' + transform;
//     }
//     if (transition) {
//       last.transition = last.transition ?? '' + ' ' + transition;
//     }
//     return { ...last, ...other };
//   }, {});
// }

// export const DIR_MAP = {
//   column: 'bottom',
//   'column-reverse': 'top',
//   row: 'right',
//   'row-reverse': 'left',
// } as const;

// export const EMPTY_OFFSET = {
//   left: false,
//   top: false,
//   bottom: false,
//   right: false,
// };

// export function getSortedChildrenMiddleLine(element: HTMLCollection) {
//   return Array.from(element).reduce<{ v: number[]; h: number[] }>(
//     (last, child) => {
//       const { top, left, width, height } = child.getBoundingClientRect();
//       if (child.className !== 'ghost') {
//         last.v.push(top + height / 2);
//         last.h.push(left + width / 2);
//       }
//       return last;
//     },
//     { v: [], h: [] },
//   );
// }
