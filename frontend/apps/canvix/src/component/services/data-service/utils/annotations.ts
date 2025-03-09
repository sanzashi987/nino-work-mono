// eslint-disable-next-line import/prefer-default-export
export function initDataService(serviceName = 'data', configKey = 'data') {
  return function (ctor: any) {
    Object.defineProperties(ctor, {
      configKey: { value: configKey, writable: false },
      serviceName: { value: serviceName, writable: false }
    });
  };
}
