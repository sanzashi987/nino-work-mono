package service

import "context"

/** http */
type BucketServiceWeb struct{}

var BucketServiceWebImpl = &BucketServiceWeb{}

func (serv BucketServiceWeb) ListBuckets(ctx context.Context) {}

func (serv BucketServiceWeb) ListBucketDir(ctx context.Context) {}
