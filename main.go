package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	(&cli.App{
		Name:  "Healthchecker",
		Usage: "A tiny tool that checks whether a website is running or is down",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Aliases:  []string{"d"},
				Usage:    "Domain name to check.",
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port number to check.",
				Required: false,
			},
		},
		Action: HealthAction,
	}).Run(os.Args)
}
