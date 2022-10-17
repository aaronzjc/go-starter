package config

import (
	"errors"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type LogConfig struct {
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type Config struct {
	Name     string              `yaml:"name"`
	Env      string              `yaml:"env"`
	Log      LogConfig           `yaml:"logFile"`
	Host     string              `yaml:"host"`
	Port     int                 `yaml:"port"`
	Database map[string]DbConfig `yaml:"database"`
}

func (c *Config) GetAddr() string {
	return fmt.Sprintf("%s:%d", config.Host, config.Port)
}

var (
	vip    *viper.Viper
	config Config
)

func init() {
	vip = viper.New()
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

	return &config, nil
}

func Get() *Config {
	return &config
}
