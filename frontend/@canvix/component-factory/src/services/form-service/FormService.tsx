import { verifySuccessType, verifyFailType, formValueType } from './systemFunction';
import ProtoService from '../proto-service';
import { event, service } from '../proto-service/annotations';

@service('form', 'form')
class FormService extends ProtoService {
  @event('校验成功时', verifySuccessType)
    verifySuccess = 'verifySuccess';

  @event('校验失败时', verifyFailType)
    verifyFailure = 'verifyFailure';

  // 更新表单内容或者设置表单值时触发
  @event('当表单值变化时', formValueType)
    formValueChange = 'formValueChange';
}

export default FormService;
