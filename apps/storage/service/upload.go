package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
	dao "github.com/sanzashi987/nino-work/apps/storage/db/dao"
	model "github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/proto/storage"
)

type UploadServiceRpc struct{}

var UploadServiceRpcImpl = &UploadServiceRpc{}

const chunkSize = 1024 * 1024 / 2

func GetUploadServiceRpc() storage.StorageServiceHandler {
	return UploadServiceRpcImpl
}

func (serv UploadServiceRpc) UploadFile(ctx context.Context, stream storage.StorageService_UploadFileStream) (err error) {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	bucketDao := dao.NewBucketDao(ctx)
	var bucket *model.Bucket
	if req.BucketId > 0 {
		bucket, err = bucketDao.GetBucket(uint(req.BucketId))
	} else {
		return fmt.Errorf("bucket information required")
	}

	if err != nil {
		return fmt.Errorf("bucket not found: %v", err)
	}

	res := storage.FileDetailResponse{}
	uid, err := uuid.NewRandom()
	if err != nil {
		return
	}
	uuidStr := uid.String()
	tempFile, err := os.CreateTemp("", uuidStr)
	writer := bufio.NewWriter(tempFile)
	if err != nil {
		return
	}
	defer tempFile.Close()

	var size int64 = 0

	for {
		var req *storage.FileUploadRequest
		req, err = stream.Recv()

		if err != io.EOF {
			return
		}

		writer.Write(req.Data)
		if buffered := writer.Buffered(); buffered > chunkSize {
			size += int64(buffered)
			err = writer.Flush()
			if err != nil {
				fmt.Println("Flush Buffer Error:", err)
				return
			}
		}

		if err == io.EOF {
			size += int64(writer.Buffered())
			e := writer.Flush()
			if e != nil {
				fmt.Println("Flush Buffer Error:", e)
				return
			}
			break
		}

	}

	tempFilePath := tempFile.Name()
	mimeType, err := mimetype.DetectFile(tempFilePath)

	// dt := time.Now().Format("2006/01/02")
	mimeTypeSTr, ext := mimeType.String(), mimeType.Extension()
	path := fmt.Sprintf("./buckets/%s/%s.%s", bucket.Code, uuidStr, ext)
	os.Rename(tempFilePath, path)

	if err := dao.NewFileDao(ctx).CreateFile(
		uint(bucket.Id),
		req.Filename,
		mimeTypeSTr,
		path,
		uuidStr,
		ext,
		size,
	); err != nil {
		return err
	}

	res.Size = size
	res.Id, res.Path, res.MimeType, res.Extension = uuidStr, path, mimeTypeSTr, ext
	return stream.SendMsg(&res)
}

/** http */
type UploadServiceWeb struct{}

var UploadServiceWebImpl = &UploadServiceWeb{}

func (serv UploadServiceWeb) UploadFile() {}

func (serv UploadServiceRpc) GetFileDetail(ctx context.Context, in *storage.FileQueryRequest, out *storage.FileDetailResponse) error {
	fileId := in.Id
	if record, err := dao.NewFileDao(ctx).QueryFile(fileId); err != nil {
		return err
	} else {
		out.Extension, out.Id, out.Path, out.Size = record.Extension, record.FileId, record.URI, record.Size
	}

	return nil
}
