package main

import (
	"log"
	"os"

	"go-starter/internal"

	"github.com/urfave/cli"
)

const appName = "go-starter-grpc"

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
	app.Before = internal.SetupGrpc
	app.Action = internal.RunGrpc

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
