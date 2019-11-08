package service

import (
	"decode_test/pkg/e"
	"decode_test/pkg/utils"
	"mime/multipart"
)

func UploadImageService(file *multipart.FileHeader, bizType int, parentPath string) UploadImageResult {
	fileName := utils.TruncateFilename(file.Filename)
	sizeInKb := file.Size
	if sizeInKb > utils.MaxImageSize {
		return UploadImageResult{msg:"file size exceed",code:e.ParamInvalid}
	}

}

type UploadImageResult struct {
	err  error
	msg  string
	code int
	path string
}
