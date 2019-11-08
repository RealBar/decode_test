package app

import (
	error2 "decode_test/pkg/error"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ContextProxy struct {
	context *gin.Context
}

func (c *ContextProxy) Bind(obj interface{}) {
	if err := c.context.Bind(obj); err != nil {
		c.WriteResponse(http.StatusBadRequest, error2.PARAM, nil)
	}

}

func (c *ContextProxy) WriteResponse(httpCode, errCode int, data interface{}) {
	if errCode > 9000 {
		logrus
	}
	c.context.JSON(httpCode, Response{
		Code: errCode,
		Msg:  error2.GetErrorMsg(errCode),
		Data: data,
	})
}
