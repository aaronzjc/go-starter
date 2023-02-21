package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setupConfig() string {
	f, _ := os.CreateTemp("", "config")
	f.WriteString(`name: api
env: dev
log:
  level: debug
  file: "/var/log/go-starter/app.log"
http:
  tls: false
  url: your.domain.here
  port: 7980
rpc:
  port: 7981`)
	return f.Name()
}

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)

	conf, err := LoadConfig("invalid path")
	assert.NotNil(err)
	assert.Nil(conf)

	f := setupConfig()
	defer os.Remove(f)
	conf, err = LoadConfig(f)
	assert.Nil(err)
	assert.Equal(conf.Http.Port, 7980)
}
