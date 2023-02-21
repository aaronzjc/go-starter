package main

import (
	"log"
	"os"

	"go-starter/internal"

	"github.com/urfave/cli"
)

var (
	appName = "go-starter-grpc"
	usage   = "run grpc server"
	desc    = "grpc demo"
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

	app.Before = internal.SetupGrpc
	app.Action = internal.RunGrpc

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
