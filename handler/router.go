package handler

import (
	v1 "decode_test/handler/v1"
	"decode_test/pkg/app"
	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	v1API := engine.Group("/api/mediaspace/v1")

	imageApi := v1API.Group("/image")
	imageApi.POST("/upload", app.CreateHandler(v1.UploadImage))
	return engine
}
