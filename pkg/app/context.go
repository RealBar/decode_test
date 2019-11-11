package app

import (
	"decode_test/handler"
	"decode_test/model"
	"decode_test/pkg/cache"
	"decode_test/pkg/config"
	"decode_test/pkg/e"
	"decode_test/pkg/fileserver"
	"decode_test/pkg/generator"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
)

type ApplicationContext struct {
	ginC       *gin.Context
	cache      *cache.Cache
	cfg        *config.Config
	fileServer *fileserver.FileServer
	idGen      *generator.IDGenerator
	db         *model.DBWrapper
}

type ProxyHandler func(proxy *ApplicationContext)

func NewApplicationContext(cache *cache.Cache, f *fileserver.FileServer, cfg *config.Config, gen *generator.IDGenerator,
	dbWrapper *model.DBWrapper) *ApplicationContext {
	return &ApplicationContext{
		cache:      cache,
		cfg:        cfg,
		fileServer: f,
		idGen:      gen,
		db:         dbWrapper,
	}
}

func (c *ApplicationContext) Bind(obj interface{}) (int, error) {
	if err := c.ginC.Bind(obj); err != nil {
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
	return e.OK, nil
}

func (c *ApplicationContext) WriteResponse(errCode int, dataOrErr interface{}) {
	httpCode := http.StatusOK
	switch {
	case errCode == e.OK:
		httpCode = http.StatusOK
	case errCode > e.InternalErrorDelimiter:
		httpCode = http.StatusInternalServerError
	case errCode < e.BadRequestDelimiter:
		httpCode = http.StatusBadRequest
	default:
		c.Logger().WithField("errCode", errCode).Error("errCode not recognized")
	}
	err, ok := dataOrErr.(error)
	if ok {
		c.Logger().WithError(err).Error(e.GetErrorMsg(errCode))
	} else if httpCode != http.StatusOK {
		c.Logger().WithField("httpCode", httpCode).Error(e.GetErrorMsg(errCode))
	}
	c.ginC.JSON(httpCode, handler.Response{
		Code: errCode,
		Msg:  e.GetErrorMsg(errCode),
		Data: dataOrErr,
	})
}

func (c *ApplicationContext) Handle(handler ProxyHandler) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		c.ginC = ctx
		handler(c)
	}
}

func (c *ApplicationContext) GinC() *gin.Context {
	return c.ginC
}

func (c *ApplicationContext) Cache() *cache.Cache {
	return c.cache
}

func (c *ApplicationContext) Cfg() *config.Config {
	return c.cfg
}

func (c *ApplicationContext) FileServer() *fileserver.FileServer {
	return c.fileServer
}

func (c *ApplicationContext) Gen() *generator.IDGenerator {
	return c.idGen
}

func (c *ApplicationContext) DB() *model.DBWrapper {
	return c.db
}

func (c *ApplicationContext) Logger() *logrus.Logger {
	value, _ := c.ginC.Get(CtxKeyLogger)
	return value.(*logrus.Logger)
}
