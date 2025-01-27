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
type FileInfo = {
  file_id: string,
  name: string,
  uri: string,
  update_time: number,
  create_time: number
};

type DirInfo = {
  id: number,
  name :string
};

type DirResponse = {
  file: FileInfo[]
  dir: DirInfo[]
};
type BucketInfo = {
  id: number
  code: string
  dir_contents: DirResponse
};
export const getBucketInfo = defineApi<GetBucketRequest, BucketInfo>({ url: `${prefix}/bucket/info/:id` });
