package db

import (
	"go-starter/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

var (
	conf = config.Config{
		Database: map[string]config.DbConfig{
			"demo": {
				Host:     "127.0.0.1",
				Port:     3306,
				Username: "root",
				Password: "123456",
				Charset:  "utf8",
			},
		},
	}
)

func TestDb(t *testing.T) {
	require := require.New(t)

	err := Setup(&conf, &gorm.Config{})
	require.Nil(err)

	demo, err := Get("demo")
	require.Nil(err)
	require.NotEmpty(demo)

	db, _ := demo.DB()
	require.Nil(db.Ping())
}
