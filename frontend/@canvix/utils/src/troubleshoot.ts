/* eslint-disable import/prefer-default-export */
/* eslint-disable prefer-rest-params */
/**
 * USE inside a class
 * will automatically call the `troubleshooter` method inside the class instance
 * the decorator only decorate prototype methods,
 * does not works on property with arrow function
 * @param target current constructor where the decorator is annotated
 * @param name name of the key of the decorated function
 * @param descriptor
 * as the descriptor value is the legacy function, the `this` is expected to be the instance
 * however the scope binding is recommended
 */
export function troubleshoot(target: any, name: string, descriptor: PropertyDescriptor): void {
  const method = descriptor.value;
  descriptor.value = function () {
    // try to invoke `troubleshooter` inside the instance
    (this as any).troubleshooter?.apply(this, Array.from(arguments));
    return method.apply(this, arguments);
  };
}
