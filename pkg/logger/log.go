package logger

import (
	"decode_test/pkg/config"
	error2 "decode_test/pkg/e"
	"fmt"
	"github.com/sirupsen/logrus"
	"net"
	"os"
	"strconv"
	"time"
)

const (
	LogStashConnectTimeout = time.Second * 60
)

var _logger *logrus.Logger

// Setup initialize the log instance
func Setup() {
	l := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	addr := config.Cfg.LogStashConfig.Host + strconv.Itoa(config.Cfg.LogStashConfig.Port)
	conn, err := net.DialTimeout("tcp", addr, LogStashConnectTimeout)
	if err != nil {
		fmt.Printf("connect logstash e:%v", err)
		os.Exit(error2.ExitCodeConnectError)
	}
	hook := NewHook(conn, DefaultFormatter(logrus.Fields{"type": "myappName"}))

	l.Hooks.Add(hook)
	_logger = l
}

func GetLogger() *logrus.Logger {
	return _logger
}