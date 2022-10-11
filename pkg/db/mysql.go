package db

import (
	"errors"
	"fmt"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Charset  string `yaml:"charset"`
}

type DBPool struct {
	dbMap map[string]*gorm.DB
	Once  *sync.Once
}

var pool *DBPool

func init() {
	pool = &DBPool{
		dbMap: make(map[string]*gorm.DB),
	}
}

func Setup(conf map[string]*DbConfig, config *gorm.Config) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	var err error
	for dbname, v := range conf {
		// 初始化DB等
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True",
			v.Username,
			v.Password,
			v.Host,
			v.Port,
			dbname,
			v.Charset,
		)
		if pool.dbMap[dbname], err = gorm.Open(mysql.Open(dsn), config); err != nil {
			return err
		}
	}
	return nil
}

func Get(dbname string) (*gorm.DB, error) {
	db, ok := pool.dbMap[dbname]
	if !ok {
		return nil, errors.New(dbname + " not found")
	}
	return db, nil
}
