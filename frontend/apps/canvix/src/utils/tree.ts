export function walkAlongTree<T extends { children?: T[] }>(path: number[], tree: T) {
  return path.reduce<T>((last, val) => last.children![val], tree);
}

export function initArrayWith(length: number, input: any) {
  return Array.from(Array(length), () => input);
}

export function sortPath(pathA: number[], pathB: number[]) {
  // if (pathA.length === 0 || pathB.length === 0) {
  //   return pathA.length - pathB.length;
  // }
  for (let i = 0; i < Math.min(pathA.length, pathB.length); i++) {
    if (pathA[i] === pathB[i]) {
      continue;
    }
    return pathA[i] - pathB[i];
  }
  return pathA.length - pathB.length;
}
