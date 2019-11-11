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

// Setup initialize the log instance
func Setup(cfg *config.Config) *logrus.Logger {
	l := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)
	addr := cfg.LogStashConfig.Host + strconv.Itoa(cfg.LogStashConfig.Port)
	conn, err := net.DialTimeout("tcp", addr, LogStashConnectTimeout)
	if err != nil {
		fmt.Printf("connect logstash e:%v", err)
		os.Exit(error2.ExitCodeConnectError)
	}
	hook := NewHook(conn, DefaultFormatter(logrus.Fields{"type": "myappName"}))

	l.Hooks.Add(hook)
	return l
}
