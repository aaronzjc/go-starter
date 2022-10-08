package config

import (
	"errors"
	"go-starter/pkg/db"

	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	Debug    bool         `yaml:"debug"`
	LogFile  string       `yaml:"logfile"`
	Port     string       `yaml:"port"`
	Database *db.DbConfig `yaml:"database"`
}

var vip *viper.Viper

var config Config

func init() {
	vip = viper.New()
}

func LoadConfig(path string) (*Config, error) {
	//加载配置
	vip.SetConfigFile(path)
	vip.SetConfigType("yml")
	if err := vip.ReadInConfig(); err != nil {
		logrus.Error("viper read config err", err)
		return nil, errors.New("read config file error")
	}
	vip.Unmarshal(&config)

	// 监听配置变更
	vip.OnConfigChange(func(in fsnotify.Event) {
		// do something
	})
	vip.WatchConfig()

	return &config, nil
}

func Get() *Config {
	return &config
}
