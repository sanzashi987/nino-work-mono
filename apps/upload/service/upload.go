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

func GetUploadServiceRpc() upload.FileUploadServiceHandler {
	return UploadServiceRpcImpl
}

func (serv UploadServiceRpc) UploadFile(ctx context.Context, stream upload.FileUploadService_UploadFileStream) (err error) {
	res := upload.FileUploadResponse{}
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
	for {
		var req *upload.FileUploadRequest
		req, err = stream.Recv()
		if err == io.EOF {
			break
		}
		if err != io.EOF {
			return
		}

		writer.Write(req.Data)
		if writer.Buffered() > 4096 {
			err = writer.Flush()
			if err != nil {
				fmt.Println("Flush Buffer Error:", err)
				return
			}
		}
	}
	writer.Flush()

	tempFilePath := tempFile.Name()
	mimeType, err := mimetype.DetectFile(tempFilePath)

	dt := time.Now().Format("2006/01/02")
	path := fmt.Sprintf("./uploads/%s/%s.%s", dt, uuidStr, mimeType.Extension())
	os.Rename(tempFilePath, path)

	if err := model.NewUploadDao(ctx).CreateFile(mimeType.String(), path, uuidStr); err != nil {
		return err
	}

	res.Id = uuidStr
	return stream.SendMsg(&res)
}
