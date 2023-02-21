package config

import (
	"errors"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type HttpConfig struct {
	Tls  bool   `yaml:"tls"`
	Url  string `yaml:"url"`
	Port int    `yaml:"port"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type RpcConfig struct {
	Port int `yaml:"port"`
}

type Config struct {
	Name     string              `yaml:"name"`
	Env      string              `yaml:"env"`
	Log      LogConfig           `yaml:"log"`
	Http     HttpConfig          `yaml:"http"`
	Rpc      RpcConfig           `yaml:"rpc"`
	Redis    RedisConfig         `yaml:"redis"`
	Database map[string]DbConfig `yaml:"database"`
}

var (
	vip    *viper.Viper
	config *Config
)

func init() {
	vip = viper.New()
	config = &Config{}
}

func LoadConfig(path string) (*Config, error) {
	//加载配置
	vip.SetConfigFile(path)
	vip.SetConfigType("yml")
	if err := vip.ReadInConfig(); err != nil {
		return nil, errors.New("read config file error")
	}
	vip.Unmarshal(&config)

	// 监听配置变更
	vip.OnConfigChange(func(in fsnotify.Event) {
		// do something
	})
	vip.WatchConfig()

	return config, nil
}

func Get() *Config {
	return config
}
