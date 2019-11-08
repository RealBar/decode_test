package app

import (
	"decode_test/pkg/e"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ContextProxy struct {
	Context *gin.Context
}

type ProxyHandler func(proxy *ContextProxy)

func (c *ContextProxy) Bind(obj interface{}) (int, error) {
	if err := c.Context.Bind(obj); err != nil {
		return e.ParamInvalid, err
	}
	valid := validation.Validation{}
	b, err := valid.RecursiveValid(obj)
	if err != nil {
		return e.InternalError, err
	}
	if !b {
		return e.ParamInvalid, errors.New("invalid params")
	}
}

func (c *ContextProxy) WriteResponse(errCode int, dataOrErr interface{}) {
	httpCode := http.StatusOK
	switch {
	case errCode == e.OK:
		httpCode = http.StatusOK
	case errCode > e.InternalErrorDelimiter:
		httpCode = http.StatusInternalServerError
	case errCode < e.BadRequestDelimiter:
		httpCode = http.StatusBadRequest
	default:
		logrus.WithField("errCode",errCode).Error("errCode not recognized")
	}
	err, ok := dataOrErr.(error)
	if ok {
		logrus.WithError(err).Error(e.GetErrorMsg(errCode))
	} else if httpCode != http.StatusOK {
		logrus.WithField("httpCode", httpCode).Error(e.GetErrorMsg(errCode))
	}
	c.Context.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetErrorMsg(errCode),
		Data: dataOrErr,
	})
}

func CreateHandler(handler ProxyHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		proxy := &ContextProxy{Context: c}
		handler(proxy)
	}
}
