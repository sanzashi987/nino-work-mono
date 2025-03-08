import React, { useContext, FC, useMemo, useReducer } from 'react';
import memoize from 'proxy-memoize';
import produce from 'immer';
import type {
  ConnectInput,
  ConnectOuptut,
  ConnectOuptutProps,
  FiltersType,
  LayerItem,
} from '@app/types';
import { PanelMetaContext, RootMetaContext, ScreenConfigContext } from '@app/context';
import type { DataConfigTypeRuntime } from '@canvas/component-factory';
import { walkAlongTree } from '@canvas/utilities';
import { parseLayerString } from '@app/utils/split-key';
import { mergeConfig, createDeltaIdList, MergeParams } from '@app/utils/component';
import { Responsive } from '@canvas/types';

type ConnectRuntimeInput = {
  filters: FiltersType;
  config: Responsive.ConfigTypeSupportedInController;
};

const computeLocalFilters = ({ config, filters }: ConnectRuntimeInput) =>
  Object.entries(config.data ?? {}).reduce<Record<string, string[]>>((lastOuter, [key, val]) => {
    const { filters: localFilters, useFilters } = val.auxiliaries;
    let res: string[] = [];
    if (useFilters) {
      res = localFilters.reduce<string[]>((lastInner, e) => {
        if (e.enable && filters[e.id]) lastInner.push(filters[e.id].content);
        return lastInner;
      }, []);
    }
    lastOuter[key] = res;
    return lastOuter;
  }, {});

function memoLocalFilterCreator() {
  return memoize(computeLocalFilters);
}

function useMergedDataConfig(
  config: Responsive.ConfigTypeSupportedInController,
  filters: FiltersType,
) {
  const memoFilterSelector = useMemo(memoLocalFilterCreator, []);
  const computedFilters = memoFilterSelector({ config, filters });
  return useMemo(() => {
    return produce(config.data as DataConfigTypeRuntime, (draft) => {
      Object.keys(computedFilters).forEach((key) => {
        draft![key].filters = computedFilters[key];
      });
    });
  }, [config.data, computedFilters]);
}

export function useMergeConfig(params: MergeParams) {
  const { id, core, delta, theme, breakpoint, breakpoints } = params;

  const defaultProperty = useMemo(() => {
    return core?.[id] ?? {};
  }, [core, id]);

  const idList = useMemo(() => {
    return createDeltaIdList({ id, theme, breakpoint, breakpoints });
  }, [breakpoint, id, theme, breakpoints]);

  return useMemo(() => {
    return mergeConfig({ defaultProperty, delta, idList });
  }, [defaultProperty, delta, idList]);
}

export function connect(InputCom: ConnectInput): ConnectOuptut {
  const Fc: FC<ConnectOuptutProps> = ({ id, userProps, connect, primitiveUtils, chain }) => {
    const [, forceUpdate] = useReducer((state) => state++, 0);
    const { theme, breakpoint, breakpoints, projectId: dashboardId } = useContext(RootMetaContext)!;
    const { projectCode } = useContext(ScreenConfigContext)!;
    const { components, core, delta, filters, layers, info } = useContext(PanelMetaContext)!;
    const config = components[id] as Responsive.ConfigTypeSupportedInController;

    const memoData = useMergedDataConfig(config, filters);

    const memoConfig = useMergeConfig({
      id,
      core,
      delta,
      breakpoint,
      theme,
      breakpoints,
    });
    const { configChildren, childrenAllowed } = useMemo(() => {
      // console.log(chain, layers, id);
      // try {
      const res = walkAlongTree(parseLayerString(chain), {
        id: '',
        type: 'com',
        children: layers,
      } as LayerItem).children;

      const configChildren = res?.map((e) => ({ id: e.id, type: e.type }));

      const childrenAllowed: Record<string, true> = Object.fromEntries(
        configChildren?.map((e) => [e.id, true]) ?? [],
      );
      return { configChildren, childrenAllowed };
      // } catch (e) {
      //   console.log('fail to find current position with the given layers and chain');
      //   return undefined;
      // }
    }, [chain, layers]);

    const configRuntime = useMemo(() => {
      return produce(config as Responsive.ConfigTypeSupportedInControllerRuntime, (draft) => {
        draft.data = memoData;
        draft.children = configChildren;
        if (draft.type !== 'subcom') {
          draft.basic = memoConfig['basic']!;
        }
        draft.attr = memoConfig['attr']!;
        draft.hide = memoConfig['hide'];
      });
    }, [config, memoData, memoConfig, configChildren]);

    return useMemo(
      () => (
        <InputCom
          childrenAllowed={childrenAllowed}
          forceUpdate={forceUpdate}
          panelId={info.id}
          chain={chain}
          config={configRuntime}
          connect={connect}
          primitiveUtils={primitiveUtils}
          userProps={userProps}
          dashboardId={dashboardId}
          projectCode={projectCode}
        />
      ),
      [configRuntime, chain, userProps, info.id], //eslint-disable-line
    );
  };
  Fc.displayName = 'ComponentConnector';
  return Fc;
}
