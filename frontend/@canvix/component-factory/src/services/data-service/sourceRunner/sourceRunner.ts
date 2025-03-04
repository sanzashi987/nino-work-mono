/* eslint-disable no-await-in-loop */
import sandbox from '@canvix/script-sandbox';
import { uuid, isEqual } from '@canvix/utils';
import type { IdentifierSource, SourceConfigRuntime, SourceRunnerProps } from '@canvix/shared';
import { SourceType } from '@canvix/shared';
import * as getDataPack from './getDataPackage';
import RefreshTimer from '../requester/refreshTimer';
import LinkList, { LinkNode } from '../utils/linkList';
import { genCancelToken } from '../requestService';
import { RES_ERR_NOT_FOUND, FILTER_ERR_TYPE_DISMATCH, MAPPING_ERR } from '../constants';
import type { GetIdentifierType } from '../../proto-service/types';

function sourceType2Method(sourceType: SourceType) {
  return getDataPack[sourceType];
}

function ensureObject(a: any) {
  if (!!a && typeof a === 'object' && !(a instanceof Array)) {
    return a;
  }
  return {};
}

type FetchDataType = (p?: Record<string, any>) => Promise<void>;
type ShortCircuitType = {
  shortCircuit: boolean;
  dataMapped?: any;
  errorMessage?: any;
};
class BaseSourceRunner {
  timer: RefreshTimer = new RefreshTimer();

  cancelToken: ReturnType<typeof genCancelToken> | null = null;

  dataRaw: Record<string, any>[] = [];

  dataFiltered: any;

  dataMapped: Record<string, any>[] = [];

  sourceOverrideCache: Record<string, any> = {};

  filterExternalCache: Record<string, any> = {};

  constructor(public props: SourceRunnerProps, public getIdentifier: GetIdentifierType) {}

  async runProcessData() {
    const processor = this.processData();
    let result: IteratorResult<ShortCircuitType> = { value: { shortCircuit: false }, done: false };
    do {
      result = await processor.next();
    } while (!result.done && !result.value.shortCircuit);
    if (!result.done) {
      this.handleShortCircuit(result.value);
    }
  }

  handleShortCircuit(shortCircuit: ShortCircuitType) {
    const { dataMapped, errorMessage } = shortCircuit;
    this.dataMapped = dataMapped;
    this.handleError(errorMessage);
  }

  async* processData() {
    yield this.filterData();
    yield this.mapData();
    this.notifyController();
    // yield { shortCircuit: false };
  }

  notifyController(): void {
    const data = this.dataMapped;
    const { sourceName } = this.props;
    this.props.setData(sourceName, data);
    this.props.$emit(this.dataUpdatedEvent, data);
  }

  setDataRaw = (payload: Record<string, any>[]) => {
    if (this.cancelToken) {
      this.cancelToken.cancel();
      // this.cancelToken();
      this.cancelToken = null;
    }
    if (payload === undefined) return; //
    this.dataRaw = payload;
    // this.processData();
    this.runProcessData();
  };

  handleError(err: any = {}): void {
    this.props.$emit(this.dataFailureEvent, {
      sourceName: this.props.sourceName,
      error: JSON.parse(JSON.stringify(err))
    });
    this.props.$emit('throwError', err);
  }

  pendingRequest = async (source: any, identifier: IdentifierSource) => {
    const { type } = this.props.sourceConfig;
    if (this.cancelToken) {
      this.cancelToken.cancel('Cancel pending request and retry');
    }
    this.cancelToken = genCancelToken();
    const method = sourceType2Method(type);
    try {
      const res = await method(source, identifier, { cancelToken: this.cancelToken.token });
      this.cancelToken = null;
      return res;
    } catch (e) {
      return { needUpdate: false, output: [], error: e };
    }
  };

  fetchData: FetchDataType = async (sourceValue) => {
    const { source, type } = this.props.sourceConfig;
    if (type === SourceType.Passive) return;
    const { sourceName } = this.props;
    const identifier = { ...this.getIdentifier(), sourceName };
    this.sourceOverrideCache = { ...this.sourceOverrideCache, ...(sourceValue ?? {}) };
    const sourceOverridden = { ...source, ...this.sourceOverrideCache };
    const originalData = source ? await this.pendingRequest(sourceOverridden, identifier) : null;
    if (!originalData) {
      this.handleError(RES_ERR_NOT_FOUND);
      return;
    }
    const { needUpdate, output, error } = originalData;
    this.dataRaw = output as Record<string, any>[];
    if (!needUpdate) {
      // TODO output as the hint , error indicates the error type predefined
      this.dataMapped = error as any;
      this.handleError(error);
      return;
    }
    // this.processData();
    this.runProcessData();
  };

  invokeFilter = (filterValue?: Record<string, any>) => {
    this.filterExternalCache = { ...this.filterExternalCache, ...ensureObject(filterValue) };
    this.runProcessData();
  };

  async filterData(): Promise<ShortCircuitType> {
    const { filters = [] } = this.props.sourceConfig;
    let dataFiltered = JSON.parse(JSON.stringify(this.dataRaw)); // TODO replace the deepclone
    const { comId } = this.getIdentifier();
    for (const filterStr of filters) {
      try {
        dataFiltered = await sandbox.runInSandbox({
          args: ['data', 'externalValue', filterStr],
          argsValue: [dataFiltered, this.filterExternalCache],
          id: `${comId}@component_filter_${uuid()}`
        });
      } catch (e) {
        this.dataFiltered = dataFiltered;
        return { shortCircuit: false };
      }
    }
    this.dataFiltered = dataFiltered;
    return { shortCircuit: false };
  }

  mapData(): ShortCircuitType {
    try {
      const { mappingList = {} } = this.props.sourceConfig;
      const mapRules = Object.entries(mappingList);
      if (mapRules.length === 0) {
        this.dataMapped = this.dataFiltered;
        return { shortCircuit: false };
      }
      const res = this.getDataFiltered(this.dataFiltered);
      if (res.shortCircuit) return res;

      const gen = res.dataMapped.transverse();
      const arr: Record<string, any>[] = [];
      for (const v of gen) {
        const objItem: Record<string, any> = { ...v.node };

        for (const [mapKey, mapValue] of mapRules) {
          if (mapValue && Reflect.has(v.node, mapValue as string)) {
            v.node[mapValue as string] !== undefined
              && (objItem[mapKey] = v.node[mapValue as string]);
          } else {
            objItem[mapKey] !== undefined && (objItem[mapKey] = v.node[mapKey]);
          }
        }
        Object.keys(objItem).length && arr.push(objItem);
      }
      this.dataMapped = arr;
      return { shortCircuit: false };
    } catch (e) {
      console.log(e);
      return { shortCircuit: true, dataMapped: [], errorMessage: MAPPING_ERR(e) };
    }
  }

  getDataFiltered(dataArr: Record<string, any>[]): ShortCircuitType {
    if (!dataArr?.forEach) {
      // do not process mapping
      return { shortCircuit: true, dataMapped: [], errorMessage: FILTER_ERR_TYPE_DISMATCH };
    }
    const prepare = new LinkList<Record<string, any>>();
    dataArr.forEach((item) => {
      const linkNode = new LinkNode<Record<string, any>>(item, item.id);
      prepare.append(linkNode);
    });
    return { shortCircuit: false, dataMapped: prepare, errorMessage: FILTER_ERR_TYPE_DISMATCH };
  }

  switchMode = (): void => {
    const { controlledMode = null, autoUpdate = null } = this.props.sourceConfig;
    this.timer.stop();
    if (controlledMode) return;
    if (autoUpdate) {
      this.timer.setTimerForTarget({
        promiseMethod: this.fetchData,
        times: autoUpdate * 1000
      })();
    } else {
      this.fetchData();
    }
  };

  updateConfig(nextConfig: SourceConfigRuntime): void {
    const { source, filters, controlledMode, autoUpdate, mappingList } = this.props.sourceConfig;
    const { source: s, filters: f, controlledMode: c, autoUpdate: a, mappingList: m } = nextConfig;
    this.props.sourceConfig = nextConfig;
    // TODO to be disccussed , whether the identifier should be updated here
    if (controlledMode !== c || a !== autoUpdate) {
      this.switchMode();
    } else if (source !== s) {
      this.fetchData();
    } else if (!isEqual(filters, f)) {
      // this.processData();
      this.runProcessData();
    } else if (!isEqual(mappingList, m)) {
      const shortCircuit = this.mapData();
      if (shortCircuit.shortCircuit) {
        this.handleShortCircuit(shortCircuit);
        return;
      }
      this.notifyController();
    }
  }

  get dataUpdatedEvent(): string {
    return `dataUpdated.${this.props.sourceName}`;
  }

  get dataFailureEvent(): string {
    return `dataFailure.${this.props.sourceName}`;
  }

  get dataResponse() {
    return { origin: this.dataRaw, value: this.dataMapped };
  }

  destroy = () => {
    this.timer.destory();
  };
}

class SourceRunner extends BaseSourceRunner {
  constructor(public props: SourceRunnerProps, public getIdentifier: GetIdentifierType) {
    super(props, getIdentifier);
    this.switchMode();
  }
}

export default SourceRunner;
export { BaseSourceRunner };
