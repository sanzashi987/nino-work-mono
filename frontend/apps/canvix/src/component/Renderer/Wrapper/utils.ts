import sandbox from '@canvas/script-sandbox';
import { getRuntimeConfig } from '@app/utils/component';
import { createMemo } from '@canvas/utilities';
import type {
  UnifiedRenderOption,
  UnifiedRenderUtilInsideWrapper,
} from '@canvas/component-factory';
import { ComWrapperInstance } from './com-wrapper';
import type { RuntimeInterface } from '../context';

export function createUtils(
  scope: ComWrapperInstance,
  getAssetsUrl: RuntimeInterface['getAssetsUrl'],
) {
  const { primitiveUtils: screenUtils, config } = scope.props;
  const _getAssetsUrl = (fileName: string) => {
    const { name, version, user, isDebugger } = config.com;
    return getAssetsUrl({
      name,
      version,
      isDebugger,
      path: fileName,
      user: scope.props.projectCode,
      cType: user ? 1 : 0,
    });
  };

  /** merge the componentId  as `parentId` to the `userProps` in render function */
  const render: UnifiedRenderUtilInsideWrapper = (opt) => {
    if (!scope.props.childrenAllowed[opt.id]) {
      console.log('cannot render the children that does not belong to the component');
      return null;
    }

    const nextOpt: UnifiedRenderOption = {
      ...opt,
      // auto generate chain from `key`
      nextChain: `${scope.props.chain}.${opt.key}`,
      userProps: {
        ...(opt.userProps ?? {}),
        // the children always receive its parent's id
        parentId: config.id,
      },
    };
    return screenUtils.general.render(nextOpt);
  };

  const memoBasic = createMemo((basic: Record<string, any>, com: any, type: any) => {
    return getRuntimeConfig({
      input: { basic, com, type },
      runtimeKeys: ['basic'],
      config: {
        getAssetsUrl: _getAssetsUrl,
      },
    }).basic;
  });

  const memoAttr = createMemo((attr: Record<string, any>, com: any, type: any) => {
    return getRuntimeConfig({
      input: { attr, com, type },
      runtimeKeys: ['attr'],
      config: {
        getAssetsUrl: _getAssetsUrl,
      },
    }).attr;
  });

  const _getRuntimeConfig = (input: Record<string, any>) => {
    const { com, type } = input;
    return {
      attr: memoAttr(input.attr, com, type),
      basic: memoBasic(input.basic, com, type),
    };
  };

  return {
    getAssetsUrl: _getAssetsUrl,
    ...screenUtils.general,
    render,
    $emit: scope.emit,
    runInSandbox: sandbox.runInSandbox,
    getRuntimeConfig: _getRuntimeConfig,
  };
}
