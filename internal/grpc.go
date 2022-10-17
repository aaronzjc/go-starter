package internal

import (
	"errors"
	"fmt"
	"go-starter/internal/config"
	"go-starter/internal/route"
	"go-starter/pkg/logger"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

func SetupGrpc(ctx *cli.Context) error {
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
	}
	return nil
}

func RunGrpc(ctx *cli.Context) error {
	conf := config.Get()

	ln, err := net.Listen("tcp", conf.GetAddr())
	if err != nil {
		logger.Fatal("listen failed", err)
	}
	rpcServer := grpc.NewServer()
	route.SetupRpc(rpcServer)

	go rpcServer.Serve(ln)
	logger.Info("grpc listen at ", conf.GetAddr())

	// 监听关闭信号
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGQUIT, os.Interrupt, syscall.SIGTERM)
	<-sig

	// 收到关闭信号，主动回收连接
	rpcServer.GracefulStop()
	logger.Info("[STOP] server shutdown ok")
	return nil
}
