package service

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"

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
				fmt.Println("刷新缓冲区时出错:", err)
				return
			}
		}
	}
	writer.Flush()

	mimetype.Detect()

	return
}
