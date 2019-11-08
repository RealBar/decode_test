package handler

import (
	"github.com/gin-gonic/gin"
)

func initRouter() {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	v1API := engine.Group("/api/mediaspace/v1")

	imageApi := v1API.Group("/image")
}
