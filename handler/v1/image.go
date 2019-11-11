package v1

import (
	"decode_test/pkg/app"
	"decode_test/pkg/e"
	"decode_test/service"
)

type UploadImageRequest struct {
	ParentPath string `form:"parent_path" valid:""`
	//`form:"image_file" valid:""`
}

func UploadImage(c *app.ApplicationContext) {
	var req UploadImageRequest
	code, err := c.Bind(&req)
	if err != nil {
		c.WriteResponse(code, err)
		return
	}
	formFile, err := c.GinC().FormFile("image_file")
	if err != nil {
		c.WriteResponse(e.ParamInvalid, err)
		return
	}
	service.UploadImageService(formFile, req.ParentPath)
}

type UpdateImageRequest struct {
	MediaID int64 `form:"media_id" valid:""`
	//`form:"image_file" valid:""`
}

func UpdateImage(c *app.ApplicationContext) {

}

func DeleteImage(c *app.ApplicationContext) {

}

func QueryImage(c *app.ApplicationContext) {

}