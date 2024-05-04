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
		if err == io.EOF {
			break
		}
		if err != io.EOF {
			return
		}

		writer.Write(req.Data)
		if buffered := writer.Buffered(); buffered > 4096 {
			size += int64(buffered)
			err = writer.Flush()
			if err != nil {
				fmt.Println("Flush Buffer Error:", err)
				return
			}
		}
	}
	size += int64(writer.Buffered())
	writer.Flush()

	tempFilePath := tempFile.Name()
	mimeType, err := mimetype.DetectFile(tempFilePath)

	dt := time.Now().Format("2006/01/02")
	path := fmt.Sprintf("./uploads/%s/%s.%s", dt, uuidStr, mimeType.Extension())
	os.Rename(tempFilePath, path)

	if err := model.NewUploadDao(ctx).CreateFile(mimeType.String(), path, uuidStr); err != nil {
		return err
	}

	res.Size = size
	res.Id, res.Path, res.MimeType, res.Extension = uuidStr, path, mimeType.String(), mimeType.Extension()
	return stream.SendMsg(&res)
}

func (serv UploadServiceRpc) GetFileDetail(ctx context.Context, in *upload.FileQueryRequest, out *upload.FileDetailResponse) (err error) {

}
