package service

import (
	"decode_test/pkg/app"
	"decode_test/pkg/e"
	"decode_test/pkg/utils"
	"github.com/disintegration/imaging"
	"image"
	"mime/multipart"
)

func UploadImageService(c *app.ApplicationContext, file *multipart.FileHeader, parentPath string) UploadImageResult {
	sizeInByte := file.Size
	if sizeInByte > app.MaxImageSize || sizeInByte == 0 {
		return UploadImageResult{msg: "file size exceeds or is 0", code: e.ParamInvalid}
	}
	f, err := file.Open()
	if err != nil {
		return UploadImageResult{msg: "open upload file error", err: err, code: e.InternalError}
	}
	defer f.Close()

	cfg, formatStr, err := image.DecodeConfig(f)
	if err != nil {
		return UploadImageResult{msg: "decode config failed, maybe image format not supported", err: err,
			code: e.ParamInvalid}
	}
	if cfg.Height > app.MaxImageHeight || cfg.Width > app.MaxImageWidth {
		return UploadImageResult{msg: "image width or height exceeds", err: err, code: e.ParamInvalid}
	}
	format := app.GetMediaFormat(formatStr)
	if format == app.Unknown {
		return UploadImageResult{msg: "image format \"" + formatStr + "\" not supported"}
	}
	img, _, err := image.Decode(f)
	if err != nil {
		return UploadImageResult{msg: "decode file failed, image file damaged or format not supported", err: err,
			code: e.ParamInvalid}
	}

	// generate thumbnail if necessary
	thumbnail := img
	w, h, need := app.ThumbnailSize(img.Bounds())
	if need {
		thumbnail = imaging.Resize(img, w, h, imaging.NearestNeighbor)
	}
	fileName := utils.TruncateFilename(file.Filename, app.MaxFileNameLen)
	md5, err := utils.GenerateMD5(f)
	if err != nil {
		return UploadImageResult{msg: "decode file failed, image file damaged or format not supported", err: err,
			code: e.ParamInvalid}
	}
	ownerID := c.GinC().GetInt64(app.CtxKeyOwnerID)
	err = c.DB().CreateMedia(c, fileName, app.ImageType, format, md5, sizeInByte, ownerID, "")
	if err != nil {
		return UploadImageResult{msg: "create media failed", err: err, code: e.InternalError}
	}

}

type UploadImageResult struct {
	err  error
	msg  string
	code int
	data interface{}
}
