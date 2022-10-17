package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	assert := assert.New(t)

	_, err := LoadConfig("invalid path")
	assert.NotNil(err)
}
