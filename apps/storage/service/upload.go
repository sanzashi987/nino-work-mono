package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sanzashi987/nino-work/apps/storage/consts"
	dao "github.com/sanzashi987/nino-work/apps/storage/db/dao"
	model "github.com/sanzashi987/nino-work/apps/storage/db/model"
	"github.com/sanzashi987/nino-work/pkg/db"
	"github.com/sanzashi987/nino-work/proto/storage"
)

type UploadServiceRpc struct{}

var UploadServiceRpcImpl = &UploadServiceRpc{}

const chunkSize = 1024 * 1024 / 2

func GenUUID() (*string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	uuidStr := uid.String()
	return &uuidStr, nil
}

func GetUploadServiceRpc() storage.StorageServiceHandler {
	return UploadServiceRpcImpl
}

func (serv UploadServiceRpc) UploadFile(ctx context.Context, stream storage.StorageService_UploadFileStream) (err error) {
	req, err := stream.Recv()
	if err != nil {
		return err
	}

	tx := db.NewTx(ctx)

	var bucket *model.Bucket
	if req.BucketId > 0 {
		bucket, err = dao.GetBucket(tx, req.BucketId)
	} else {
		return fmt.Errorf("bucket information required")
	}

	if err != nil {
		return fmt.Errorf("bucket not found: %v", err)
	}

	res := storage.FileDetailResponse{}
	uuidStr, err := GenUUID()
	if err != nil {
		return
	}
	tempFile, err := os.CreateTemp("", *uuidStr)
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
	mimeTypeStr, ext := mimeType.String(), mimeType.Extension()
	path := fmt.Sprintf("./buckets/%s/%s.%s", bucket.Code, uuidStr, ext)
	os.Rename(tempFilePath, path)

	toInsert := model.Object{
		BucketID:  uint64(bucket.Id),
		FileId:    *uuidStr,
		URI:       path,
		Name:      req.Filename,
		MimeType:  mimeTypeStr,
		Extension: ext,
		Size:      size,
	}

	if err := tx.Create(&toInsert).Error; err != nil {
		return err
	}

	res.Size = size
	res.Id, res.Path, res.MimeType, res.Extension = *uuidStr, path, mimeTypeStr, ext
	return stream.SendMsg(&res)
}

/** http */
type UploadServiceWeb struct{}

var UploadServiceWebImpl = &UploadServiceWeb{}

func parseFile(ctx *gin.Context, payload *UploadFilePayload, bucketPath string, file *multipart.FileHeader) (*model.Object, error) {
	uuidStr, err := GenUUID()
	if err != nil {
		return nil, err
	}

	ext := filepath.Ext(file.Filename)
	fileName := filepath.Base(file.Filename)

	path := fmt.Sprintf("%s/%d/%s", bucketPath, payload.BucketID, *uuidStr)
	if ext != "" {
		path = fmt.Sprintf("%s%s", path, ext)
	}

	if err = ctx.SaveUploadedFile(file, path); err != nil {
		return nil, err
	}

	mimeType, err := mimetype.DetectFile(path)
	if err != nil {
		return nil, err
	}

	toInsert := model.Object{
		BucketID:  payload.BucketID,
		ParentId:  payload.PathId,
		FileId:    *uuidStr,
		URI:       path,
		Dir:       false,
		Name:      fileName,
		MimeType:  mimeType.String(),
		Extension: ext,
		Size:      file.Size,
	}

	return &toInsert, nil
}

type UploadFilePayload struct {
	BucketID uint64
	PathId   uint64
	Files    []*multipart.FileHeader
}

func (serv UploadServiceWeb) UploadFile(ctx *gin.Context, userId uint64, payload *UploadFilePayload) ([]string, error) {
	bucketPath := ctx.GetString(consts.BucketPath)

	tx := db.NewTx(ctx).Begin()

	toInsert := []*model.Object{}
	res := []string{}

	existFiles := []*model.Object{}

	if err := tx.Where("bucket_id = ? AND parent_id = ? AND dir = ?", payload.BucketID, payload.PathId, false).Find(&existFiles).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	existFileNames := make(map[string]any)
	for _, file := range existFiles {
		existFileNames[file.Name] = true
	}

	for _, file := range payload.Files {
		if _, exists := existFileNames[file.Filename]; exists {
			continue
		}
		obj, err := parseFile(ctx, payload, bucketPath, file)
		if err != nil {
			return nil, err
		}
		toInsert = append(toInsert, obj)
		res = append(res, obj.FileId)
	}

	if err := tx.Create(&toInsert).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	return res, nil
}

func (serv UploadServiceWeb) UploadLargeFile(ctx context.Context) {

}

func (serv UploadServiceRpc) GetFileDetail(ctx context.Context, in *storage.FileQueryRequest, out *storage.FileDetailResponse) error {
	fileId := in.Id
	record := model.Object{}

	if err := db.NewTx(ctx).Where("file_id = ?", fileId).Find(&record).Error; err != nil {
		return err
	} else {
		out.Extension, out.Id, out.Path, out.Size = record.Extension, record.FileId, record.URI, record.Size
	}

	return nil
}
