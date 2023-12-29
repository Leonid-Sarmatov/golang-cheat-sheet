package main

import (
	"github.com/urfave/cli/v2"
	"fmt"
	"os"
	"log"
)

func main() {
	app := &cli.App {
		Name: "Healthchecker",
		Usage: "A tiny tool for checks whether a website is running or is down",
		Flags: []cli.Flags {
			&cli.StringFlag {
				Name: "domain",
				Aleases: []string{"d"},
				Usage: "Domain name to chek.",
				Required: true,
			},
			&cli.StringFlag {
				Name: "port",
				Aleases: []string{"p"},
				Usage: "Port number to check.",
				Required: false,
			},
		},
		Actions: func(c *cli.Context) error {
			port := c.String("port")
			if port == "" {
				port = "80"
			}
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Falal(err)
	}
}