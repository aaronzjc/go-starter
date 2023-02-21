package main

import (
	"go-starter/internal"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var (
	appName = "go-starter-cron"
	usage   = "run cron"
	desc    = "cron demo"
	version = "0.1"
)

func main() {
	app := *cli.NewApp()
	app.Name = appName
	app.Usage = usage
	app.Description = desc
	app.Version = version

	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config,c",
			Usage: "(config) Load configuration from `FILE`",
		},
	}

	app.Before = internal.SetupCli
	app.Commands = internal.RegistCmds()

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
