package v1

import (
	"decode_test/pkg/app"
	"decode_test/pkg/e"
	"decode_test/service"
)

func UploadImage(c *app.ContextProxy) {
	var req UploadImageRequest
	code, err := c.Bind(&req)
	if err != nil {
		c.WriteResponse(code, err)
		return
	}
	formFile, err := c.Context.FormFile("image_file")
	if err != nil {
		c.WriteResponse(e.ParamInvalid, err)
		return
	}
	service.UploadImageService(formFile,req.BizType,req.ParentPath)
}

type UploadImageRequest struct {
	BizType    int    `form:"biz_type" valid:""`
	ParentPath string `form:"parent_path" valid:""`
	//`form:"image_file" valid:""`
}
