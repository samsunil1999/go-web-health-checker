package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "healthchecker",
		Usage: "It is a tool to check whether the website is running or not",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "domain",
				Usage:    "Domain name to check",
				Aliases:  []string{"d"},
				Required: true,
			},
			&cli.StringFlag{
				Name:     "port",
				Usage:    "Port number to check",
				Aliases:  []string{"p"},
				Required: false,
			},
		},
		Action: func(c *cli.Context) error {
			port := c.String("port")
			if c.String("port") == "" {
				port = "80"
			}

			// call the Check() with domian & port
			status := Check(c.String("domain"), port)
			fmt.Println(status)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
