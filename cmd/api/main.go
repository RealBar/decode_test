package main

import (
	"decode_test/pkg/config"
	"decode_test/pkg/logger"
)

func main() {
	config.Setup()
	logger.Setup()

}
