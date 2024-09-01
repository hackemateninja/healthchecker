package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func HealthAction(c *cli.Context) error {
	port := c.String("port")
	domain := c.String("domain")

	if len(port) <= 0 {
		port = "80"
	}
	status := Check(domain, port)
	fmt.Println(status)
	return nil
}
