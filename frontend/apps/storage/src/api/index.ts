import { ModelMeta, PaginationResponse, PagninationRequest } from '@nino-work/shared';
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
  bucket_id :number | string
};
type FileInfo = {
  file_id: string,
  name: string,
  uri: string,
  update_time: number,
  create_time: number
};

export type DirInfo = {
  id: number,
  name :string
};

export type DirResponse = {
  files: FileInfo[]
  dirs: DirInfo[]
};
export type BucketInfo = {
  id: number
  code: string
  dir_contents: DirResponse
  root_path_id: number
};
export const getBucketInfo = defineApi<GetBucketRequest, BucketInfo>({ url: `${prefix}/bucket/info` });

type ListBucketDirRequest = {
  bucket_id: number | string
  path_id: number
};

export const listBucketDir = defineApi<ListBucketDirRequest, DirResponse>({ url: `${prefix}/bucket/list` });

export type BucketMeta = Omit<ModelMeta, 'name'>;

export const createBucket = defineApi<{ code: string }, { id: number }>({
  url: `${prefix}/bucket/create`,
  method: 'POST'
});
