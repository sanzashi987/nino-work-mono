import React from 'react';
import type { DataResponseType } from '@canvix/shared';
import SourceRunnerWrapper from './sourceRunner';
import ProtoService from '../proto-service';
import { event, action, service } from '../proto-service/annotations';
import type {
  EventCollection,
  HandlerCollection,
  EndpointType,
  AnnotationEndpointType
} from '../proto-service/types';

const dataRawType: AnnotationEndpointType = {
  description: '大部分组件接收的数据为对象数组,不同组件需求的数据类型不相同',
  fields: {
    type: 'array',
    name: 'data',
    default: '$configData'
  }
};

const dataFailureType: AnnotationEndpointType = {
  description: '标准数据为对象数组',
  fields: {
    type: 'any',
    name: 'error'
    // default: '$',
  }
};
const IObjectType: AnnotationEndpointType = {
  description: '设置提供给过滤器的额外字段, 以非嵌套对象的形式提供',
  fields: {
    type: 'object',
    name: 'filterVariables',
    children: {
      任意字段: {
        type: 'any',
        name: '任意字段'
      }
    }
  }
};

const linkValueType: AnnotationEndpointType = {
  description: '输入格式严格一致的参数将重写其中的配置',
  fields: {
    type: 'object',
    name: 'source',
    default: '$sourceConfig'
  }
};

function getEndpointFromPackage(
  collection: HandlerCollection | EventCollection,
  prefix: string,
  packageJSON: Record<string, any>
) {
  const { apis } = packageJSON;
  const template = [...collection.entries()];
  const res: EndpointType[] = [];
  Object.entries(apis ?? {}).forEach((kv) => {
    const [sourceName, { name = '' }] = kv as [string, Record<string, any>];
    template.forEach(([actionName, desc]) => {
      if (!desc.isPublic) return;
      res.push({
        /**
         * The reason for `actionName` concated first is the
         * method invokation requires a specified entry and
         * returns the Proxy, rather than using the `sourceName`
         * which leading to proxy the whole component instance.
         */
        id: `${prefix}.${actionName}.${sourceName}`,
        ...desc,
        name: desc.name.replace('${}', name)
      });
    });
  });
  return res;
}

const handlers = ['fetchData', 'setDataRaw', 'invokeFilter'] as const;

@service('data', 'data')
class DataService extends ProtoService {
  static getComponentActions(
    actions: HandlerCollection,
    serviceName: string,
    packageJSON: Record<string, any>
  ) {
    return getEndpointFromPackage(actions, serviceName, packageJSON);
  }

  static getComponentEvents(
    events: EventCollection,
    serviceName: string,
    packageJSON: Record<string, any>
  ) {
    return getEndpointFromPackage(events, serviceName, packageJSON);
  }

  subSources;

  handlerForward;

  $scopedEmit;

  constructor(props: ProtoService['props']) {
    super(props);
    this.subSources = Object.fromEntries(
      Object.keys(props.config).map((key) => [key, React.createRef<any>()])
    );
    this.handlerForward = Object.fromEntries(
      handlers.map((key) => [
        key,
        Proxy.revocable(
          {},
          { get: (target, sourceName: string) => this.subSources[sourceName].current?.[key] }
        )
      ])
    );
    this.$scopedEmit = this.$emit.bind(this);
  }

  setData = (sourceName: string, data: Record<string, any>[]) => {
    this.props.setState((prev: any) => ({ data: { ...prev.data, [sourceName]: data } }));
  };

  @event('数据更新完成时(${})', dataRawType)
    dataUpdated = 'dataUpdated';

  @event('数据更新失败时(${})', dataFailureType)
    dataFailure = 'dataFailure';

  @action('设置数据(${})', dataRawType)
  get setDataRaw() {
    return this.handlerForward.setDataRaw?.proxy;
  }

  @action('请求数据(${})', linkValueType)
  get fetchData() {
    return this.handlerForward.fetchData?.proxy;
  }

  @action('调用过滤器(${})', IObjectType)
  get invokeFilter() {
    return this.handlerForward.invokeFilter?.proxy;
  }

  @action('设置所有数据', dataRawType, false)
    setDataAll = (data: any) => {
      Object.values(this.subSources).forEach((e) => {
        e.current?.setDataRaw(data);
      });
    };

  get dataResponse(): DataResponseType {
    return Object.fromEntries(
      Object.entries(this.subSources).map(([k, v]) => [k, v.current.getDataResponse?.()])
    ) as DataResponseType;
  }

  render(): React.ReactNode {
    return (
      <>
        {Object.entries(this.props.config).map(([sourceName, sourceConfig]) => (
          <SourceRunnerWrapper
            ref={this.subSources[sourceName]}
            sourceName={sourceName}
            sourceConfig={sourceConfig}
            setData={this.setData}
            $emit={this.$scopedEmit}
            getIdentifier={this.props.getIdentifier}
            key={sourceName}
          />
        ))}
      </>
    );
  }

  componentWillUnmount() {
    Object.keys(this.handlerForward).forEach((key) => {
      this.handlerForward[key].revoke();
    });
    this.handlerForward = {};
  }
}

export default DataService;
export { PassiveDataContext, BaseSourceRunner } from './sourceRunner';
export { initDataService } from './utils/annotations';
