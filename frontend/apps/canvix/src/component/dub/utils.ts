/**
 *
 * @param name component name in english characters
 * @param version the version for the component in packge version formae
 *                i.e. `1.0.0`
 */
export function composeCacheId(name: string, version: string) {
  return `${name}-${version}`;
}
