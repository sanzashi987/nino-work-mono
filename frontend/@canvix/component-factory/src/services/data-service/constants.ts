// eslint-disable-next-line
//@ts-ignore
export const canvasApiService = `${window.conf?.canvasBackendBaseUrl ?? ''}`;

export const RES_ERR_NOT_FOUND = { error: true, status: '404', message: 'Resource Not Found' };

export const FILTER_ERR_TYPE_DISMATCH = {
  error: true,
  status: 'Mapping Data Type Not Match',
  message: 'The result from filter does not match the type of Object in Array'
};

export const MAPPING_ERR = (e?: any) => ({
  error: true,
  status: 'Error occurs during processing mapping',
  message: e || 'This can be an internal error or the type error of the filtered data'
});
