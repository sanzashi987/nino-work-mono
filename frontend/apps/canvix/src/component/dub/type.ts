export enum LoadStatus {
  NULL,
  PENDING,
  RESOLVED,
}
export type CacheComponentType = {
  exports: object;
  // eslint-disable-next-line @typescript-eslint/ban-types
  factory?: Function;
  fired: boolean;
  loaded: LoadStatus;
};
