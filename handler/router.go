package handler

import (
	"decode_test/handler/middleware"
	v1 "decode_test/handler/v1"
	"decode_test/model"
	"decode_test/pkg/app"
	"decode_test/pkg/cache"
	"decode_test/pkg/config"
	"decode_test/pkg/fileserver"
	"decode_test/pkg/generator"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Setup(logger *logrus.Logger, cache *cache.Cache, f *fileserver.FileServer, cfg *config.Config,
	gen *generator.IDGenerator, dbWrapper *model.DBWrapper) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	context := app.NewApplicationContext(logger, cache, f, cfg, gen, dbWrapper)

	engine.Use(middleware.CtxMiddleware(logger))
	v1API := engine.Group("/api/mediaspace/v1")

	imageApi := v1API.Group("/image")
	imageApi.POST("/upload", context.Handle(v1.UploadImage))
	imageApi.DELETE("/delete", context.Handle(v1.DeleteImage))
	imageApi.PUT("/update", context.Handle(v1.UpdateImage))
	imageApi.GET("/query", context.Handle(v1.QueryImage))

	videoApi := v1API.Group("/video")
	videoApi.POST("/upload", context.Handle(v1.UploadVideo))
	videoApi.DELETE("/delete", context.Handle(v1.DeleteVideo))
	videoApi.PUT("/update", context.Handle(v1.UpdateVideo))
	videoApi.GET("/query", context.Handle(v1.QueryVideo))

	folderApi := v1API.Group("/folder")
	folderApi.GET("/list", context.Handle(v1.ListFolderContent))

	return engine
}
