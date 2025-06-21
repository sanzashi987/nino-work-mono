import { ModelMeta, PaginationResponse, PageSize } from '@nino-work/shared';
import defineRequester from './impls';

type BucketData = {
  id: number;
  code: string;
  update_time: number;
  create_time: number;
};

export const listBucket = defineRequester<PageSize, PaginationResponse<BucketData>>({
  url: 'bucket/list',
  method: 'POST',
});

type GetBucketRequest = {
  bucket_id: number | string;
};
type FileInfo = {
  file_id: string;
  name: string;
  uri: string;
  size: number;
  update_time: number;
  create_time: number;
};

export type DirInfo = {
  id: number;
  name: string;
};

export type DirResponse = {
  files: FileInfo[];
  dirs: DirInfo[];
};
export type BucketInfo = {
  id: number;
  code: string;
  dir_contents: DirResponse;
  root_path_id: number;
};
export const getBucketInfo = defineRequester<GetBucketRequest, BucketInfo>({ url: 'bucket/info' });

type ListBucketDirRequest = {
  bucket_id: number | string;
  path_id: number;
};

export const listBucketDir = defineRequester<ListBucketDirRequest, DirResponse>({
  url: 'bucket/dir/list',
});

export type BucketMeta = Omit<ModelMeta, 'name'>;

export const createBucket = defineRequester<{ code: string }, { id: number }>({
  url: 'bucket/create',
  method: 'POST',
});

export const createDir = defineRequester<{ bucket_id: number; parent_id: number; name: string }, void>({
  url: 'bucket/dir/create',
  method: 'POST',
});

export type UploadFileRequest = {
  bucket_id: number;
  path_id: number;
  file: File[];
};

export const uploadFiles = defineRequester<UploadFileRequest, { file_ids: string[] }>({
  url: 'asset/upload',
  method: 'POSTFORM',
});

export const deleteFile = defineRequester<{ bucket_id: number; file_id: string }, void>({
  url: 'asset/delete',
  method: 'POST',
});
