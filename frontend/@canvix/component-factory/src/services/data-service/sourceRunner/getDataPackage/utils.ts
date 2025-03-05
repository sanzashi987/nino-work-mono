export const getErrorInfo = (error: DOMException & { msg?: string }) => ({
  name: error?.name,
  status: error?.code,
  message: error?.msg || error?.message
  // data: error?.response?.data
});
