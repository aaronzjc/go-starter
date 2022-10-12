package internal

import (
	"errors"
	"fmt"
	"go-starter/internal/config"
	"go-starter/internal/task"
	"go-starter/pkg/logger"
	"os"

	"github.com/urfave/cli"
)

func SetupCli(ctx *cli.Context) error {
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
	err = logger.Setup(conf.Name, conf.LogFile)
	if err != nil {
		return err
	}

	return nil
}

func RegistCmds() []cli.Command {
	return []cli.Command{
		task.NewTestTask(),
	}
}
