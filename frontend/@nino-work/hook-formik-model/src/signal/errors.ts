function defaultThrowError(): never {
  throw new Error();
}

let throwInvalidWriteToSignalErrorFn = defaultThrowError;

export function throwInvalidWriteToSignalError() {
  throwInvalidWriteToSignalErrorFn();
}

export function setThrowInvalidWriteToSignalError(fn: () => never): void {
  throwInvalidWriteToSignalErrorFn = fn;
}
