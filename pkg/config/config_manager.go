package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"os"
	"strings"
	"time"
)

const (
	CONFIG_ENV_KEY = "ENV"

	ENV_DEV  Env = "DEV"
	ENV_TEST Env = "TEST"
	ENV_LIVE Env = "LIVE"
)

var (
	Cfg        *Config
	CurrentEnv = ENV_DEV
)

type Env string

type Config struct {
	RedisConfig      RedisCfg      `toml:"redis"`
	DataBaseConfig   DataBaseCfg   `toml:"database"`
	LogStashConfig   LogStashCfg   `toml:"logstash"`
	FileServerConfig FileServerCfg `toml:"fileserver"`
	MongoDBConfig    MongoDBCfg    `toml:"mongodb"`
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
	Username  string   `toml:"username"`
	Password  string   `toml:"password"`
	HostPorts []string `toml:"host_ports"`
	DBName    string   `toml:"dbName"`
}
type DataBaseCfg struct {
	BasicCfg
	Username string `toml:"username"`
	Password string `toml:"password"`
	DBName   string `toml:"dbName"`
}

type FileServerCfg struct {
	DownLoadUrl string  `toml:"download_url"`
	UpLoadUrl   string  `toml:"upload_url"`
	Pool        PoolCfg `toml:"pool"`
}

type PoolCfg struct {
	MaxIdle int `toml:"maxIdle"`
	// in milliseconds
	IdleTimeOut time.Duration `toml:"idleTimeOut"`
	InitSize    int           `toml:"initSize"`
	MaxSize     int           `toml:"maxSize"`
}

type LogStashCfg struct {
	BasicCfg
}

func Setup() {
	env, b := os.LookupEnv(CONFIG_ENV_KEY)
	if !b {
		fmt.Printf("env key \"%v\" not found", CONFIG_ENV_KEY)
		os.Exit(-1)
	}
	CurrentEnv = Env(strings.ToUpper(env))
	configFilePath := "conf/" + strings.ToLower(env) + ".toml"
	Cfg = new(Config)
	_, e := toml.DecodeFile(configFilePath, Cfg)
	if e != nil {
		fmt.Printf("init config file e:%v", e)
		os.Exit(-1)
	}
}
