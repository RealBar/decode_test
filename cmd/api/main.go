package main

import (
	"decode_test/handler"
	"decode_test/model"
	"decode_test/pkg/cache"
	"decode_test/pkg/config"
	"decode_test/pkg/fileserver"
	"decode_test/pkg/generator"
	"decode_test/pkg/logger"
	"net/http"
	"strconv"
)

func main() {
	cfg := config.Setup()
	l := logger.Setup(cfg)
	c := cache.Setup(cfg, l)
	g := generator.Setup(l, c)
	f := fileserver.Setup(cfg, l)
	d := model.Setup(cfg, l)
	engine := handler.Setup(l, c, f, cfg, g, d)

	server := http.Server{
		Addr:           ":" + strconv.Itoa(cfg.ListenPort),
		Handler:        engine,
		MaxHeaderBytes: cfg.MaxHeaderBytes,
		ReadTimeout:    cfg.ReadTimeout,
		WriteTimeout:   cfg.WriteTimeout,
	}

	err := server.ListenAndServe()
	if err != nil {
		l.WithError(err).Error("server shutdown with error")
	} else {
		l.Info("server shut down normally")
	}
}
