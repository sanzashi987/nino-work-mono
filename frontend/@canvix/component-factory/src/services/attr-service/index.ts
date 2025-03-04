import { produce } from 'immer';
import { merge } from '@canvix/utils';
import ProtoService from '../proto-service';
import { action, service } from '../proto-service/annotations';
import { AnnotationEndpointType } from '../proto-service/types';

const AttrValueType: AnnotationEndpointType = {
  description: '输入数据结构一致的部分数据将重写对应的属性 (需保证数据的层级结构)',
  fields: {
    name: 'attr',
    type: 'object',
    default: '$attr'
  }
};

@service('attr', 'attr')
class AttrService extends ProtoService {
  @action('设置属性', AttrValueType)
    setAttr = (attrValue: Record<string, any>) => {
      this.props.setState((prev) => produce(prev, (draft: any) => {
        draft.config.attr = produce(this.props.config, (draftInner) => {
          merge(draftInner, attrValue);
        });
      }));
    };
}

export default AttrService;
