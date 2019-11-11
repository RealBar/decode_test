package middleware

import (
	"decode_test/handler"
	"decode_test/pkg/app"
	"decode_test/pkg/e"
	"decode_test/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

func CtxMiddleware(logger *logrus.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		ownerIDStr := c.GetHeader(app.CtxKeyOwnerID)
		ownerID, err := strconv.ParseInt(ownerIDStr, 10, 64)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, handler.Response{Code: e.ParamInvalid, Msg: "invalid request"})
			return
		}
		c.Set(app.CtxKeyOwnerID, ownerID)
		uuid := utils.GenerateUUID()
		c.Set(app.CtxKeyRequestID, uuid)
		c.Writer.Header().Set(app.CtxKeyResponseID, uuid)
		l := logger.WithField(app.CtxKeyRequestID, uuid)
		c.Set(app.CtxKeyLogger, l)
	}
}
