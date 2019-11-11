package config

import (
	"decode_test/pkg/app"
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
	"time"
)

type Config struct {
	RedisConfig      RedisCfg      `toml:"redis"`
	DataBaseConfig   DataBaseCfg   `toml:"database"`
	LogStashConfig   LogStashCfg   `toml:"logstash"`
	FileServerConfig FileServerCfg `toml:"fileserver"`
	MongoDBConfig    MongoDBCfg    `toml:"mongodb"`
	CurrentEnv       app.Env
	ListenPort       int
	MaxHeaderBytes   int
	ReadTimeout      time.Duration
	WriteTimeout     time.Duration
}

type BasicCfg struct {
	Host string `toml:"host"`
	Port int    `toml:"port"`
}
type RedisCfg struct {
	BasicCfg
	Password string  `toml:"password"`
	Pool     PoolCfg `toml:"pool"`
}
type MongoDBCfg struct {
	Username         string   `toml:"username"`
	Password         string   `toml:"password"`
	Addresses        []string `toml:"addresses"`
	UseSecondaryRead bool     `toml:"use_secondary_read"`
	DBName           string   `toml:"dbName"`
	Pool             PoolCfg  `toml:"pool"`
	ReplicaSet       string   `toml:"replicaSet"`
}
type DataBaseCfg struct {
	BasicCfg
	Username string  `toml:"username"`
	Password string  `toml:"password"`
	DBName   string  `toml:"dbName"`
	Pool     PoolCfg `toml:"pool"`
}

type FileServerCfg struct {
	Endpoint        string `toml:"endpoint"`
	AccessKey       string `toml:"accessKey"`
	AccessKeySecret string `toml:"accessKeySecret"`
}

type PoolCfg struct {
	// in milliseconds
	IdleTimeOut time.Duration `toml:"idleTimeOut"`
	InitSize    int           `toml:"initSize"`
	MaxSize     int           `toml:"maxSize"`
}

type LogStashCfg struct {
	BasicCfg
}

func Setup() *Config {
	cfg := new(Config)
	cfg.ListenPort = app.DefaultListenPort
	cfg.CurrentEnv = app.DefaultEnv
	cfg.MaxHeaderBytes = app.DefaultMaxHeaderBytes
	cfg.ReadTimeout = app.DefaultReadTimeout
	cfg.WriteTimeout = app.DefaultReadTimeout
	envStr, b := os.LookupEnv(app.ConfigEnvKey)
	if b {
		cfg.CurrentEnv = app.Env(strings.ToUpper(envStr))
	} else {
		fmt.Printf("envStr key \"%v\" not found, using default env:%s", app.ConfigEnvKey, app.DefaultEnv)
	}
	configFilePath := "conf/" + strings.ToLower(envStr) + ".toml"
	_, e := toml.DecodeFile(configFilePath, cfg)
	if e != nil {
		fmt.Printf("init config file e:%v", e)
		panic(e)
	}

	return cfg
}
