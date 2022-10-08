package main

import (
	"os"

	"go-starter/internal"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const appName = "go-starter"

func main() {
	app := *cli.NewApp()
	app.Name = appName
	app.Usage = "run " + appName + " server"
	app.Description = "run " + appName + " server"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  "config,c",
			Usage: "(config) Load configuration from `FILE`",
		},
	}
	app.Action = internal.RunApi

	if err := app.Run(os.Args); err != nil {
		logrus.Fatal(err)
	}
}
