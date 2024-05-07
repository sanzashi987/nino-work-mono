package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"time"

	model "github.com/cza14h/nino-work/apps/upload/db"
	"github.com/cza14h/nino-work/proto/upload"
	"github.com/gabriel-vasile/mimetype"
	"github.com/google/uuid"
)

type UploadServiceRpc struct{}

var UploadServiceRpcImpl = &UploadServiceRpc{}

const chunkSize = 1024 * 1024 / 2

func GetUploadServiceRpc() upload.FileUploadServiceHandler {
	return UploadServiceRpcImpl
}

func (serv UploadServiceRpc) UploadFile(ctx context.Context, stream upload.FileUploadService_UploadFileStream) (err error) {
	res := upload.FileDetailResponse{}
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
		var req *upload.FileUploadRequest
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

	dt := time.Now().Format("2006/01/02")
	mimeTypeSTr, ext := mimeType.String(), mimeType.Extension()
	path := fmt.Sprintf("./uploads/%s/%s.%s", dt, uuidStr, ext)
	os.Rename(tempFilePath, path)

	if err := model.NewUploadDao(ctx).CreateFile(mimeTypeSTr, path, uuidStr, ext, size); err != nil {
		return err
	}

	res.Size = size
	res.Id, res.Path, res.MimeType, res.Extension = uuidStr, path, mimeTypeSTr, ext
	return stream.SendMsg(&res)
}

func (serv UploadServiceRpc) GetFileDetail(ctx context.Context, in *upload.FileQueryRequest, out *upload.FileDetailResponse) error {
	fileId := in.Id
	if record, err := model.NewUploadDao(ctx).QueryFile(fileId); err != nil {
		return err
	} else {
		out.Extension, out.Id, out.Path, out.Size = record.Extension, record.FileId, record.URI, record.Size
	}

	return nil
}
