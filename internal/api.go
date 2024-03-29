package internal

import (
	"context"
	"errors"
	"fmt"
	"go-starter/internal/api"
	"go-starter/internal/config"
	"go-starter/pkg/logger"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

func SetupApi(ctx *cli.Context) error {
	var err error
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
	// 初始化日志组件
	err = logger.Setup(conf.Name, conf.Log.File)
	if err != nil {
		return err
	}
	// 设置调试模式
	if conf.Env != "prod" {
		logger.SetLevel(conf.Log.Level)
		gin.SetMode(gin.DebugMode)
	}
	return nil
}

func RunApi(ctx *cli.Context) error {
	conf := config.Get()
	// 启动服务器
	app := gin.New()
	api.SetupRoute(app)
	addr := fmt.Sprintf(":%d", conf.Http.Port)
	server := &http.Server{
		Addr:         addr,
		Handler:      app,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 10,
	}
	go server.ListenAndServe()
	logger.Info("[START] server listen at ", addr)

	// 监听关闭信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM)
	<-sig

	// 收到关闭信号，主动回收连接
	ctxTimeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	if err := server.Shutdown(ctxTimeout); err != nil {
		logger.Error("[STOP] server shutdown error", err)
		return err
	}
	logger.Info("[STOP] server shutdown ok")
	return nil
}
