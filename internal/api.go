package internal

import (
	"context"
	"errors"
	"fmt"
	"go-starter/internal/config"
	"go-starter/internal/route"
	"go-starter/pkg/db"
	"go-starter/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

func setup(conf *config.Config) error {
	var err error
	// 初始化日志组件
	err = logger.Setup(conf.Name, conf.LogFile)
	if err != nil {
		return err
	}

	// 初始化DB等
	err = db.Setup(conf.Database, &gorm.Config{})
	if err != nil {
		return err
	}

	// 设置调试模式
	if conf.Debug {
		logger.GetLog().SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	}
	return nil
}

func RunApi(ctx *cli.Context) error {
	if ctx.String("config") == "" {
		fmt.Println("invalid config option, use -h get full doc")
		return nil
	}

	// 初始化项目配置
	configFile := ctx.String("config")
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		return errors.New("config file not found")
	}
	conf, err := config.LoadConfig(configFile)
	if err != nil {
		return err
	}

	// 初始化一些基础组件
	err = setup(conf)
	if err != nil {
		return err
	}

	// 启动服务器
	app := gin.New()
	route.Setup(app)
	server := &http.Server{
		Addr:         conf.Addr,
		Handler:      app,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
	go server.ListenAndServe()
	logger.GetLog().Info("[START] server listen at ", conf.Addr)

	// 监听关闭信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM)
	<-sig

	// 收到关闭信号，主动回收连接
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctxTimeout); err != nil {
		logger.GetLog().Error(nil, "[STOP] server shutdown error", err)
		return err
	}
	logger.GetLog().Info("[STOP] server shutdown ok")
	return nil
}
