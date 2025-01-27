import { PaginationResponse, PagninationRequest } from '@nino-work/shared';
import defineApi from './impls';

const prefix = '/backend/v1';

type BucketData = {
  id:number,
  code: string,
  update_time: number
  create_time: number
};

export const listBucket = defineApi<PagninationRequest, PaginationResponse<BucketData>>({
  url: `${prefix}/bucket/list`,
  method: 'POST'
});

type GetBucketRequest = {
  id :number
};

type BucketInfo = {
  id: number
  code: string
};
export const getBucketInfo = defineApi<GetBucketRequest, BucketInfo>({ url: `${prefix}/bucket/info/:id` });
