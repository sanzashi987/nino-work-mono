/**
 * different from the `targetPlatform` in `process.env`, as the env variable
 * indicates the webpage is about to running in which platform
 *
 * The `TargetPlatform` here defines what kind of features the native platform
 * side is suppoorted and can be configured over the webpage editor
 *
 * Although the app is built by `flutter`, the flutter side may not always implement
 * a feature on both `IOS` and `Android` platform.
 */
export type TargetPlatform = 'IOS' | 'H5' | 'Android' | 'miniProgram' | 'SPA';

export type TargetPlatformSpecifier = {
  [Key in TargetPlatform]?: boolean;
};
