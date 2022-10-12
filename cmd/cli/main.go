package main

import (
	"go-starter/internal"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const appName = "go-starter-cli"

func main() {
	app := *cli.NewApp()
	app.Name = appName
	app.Usage = "run " + appName + " manager"
	app.Description = "run " + appName + " cli"
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
