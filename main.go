package main

import (
	"fmt"
	"os"

	"github.com/s-shin/spelunker/command"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "spelunker"
	app.Usage = "spelunker"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "",
			Subcommands: []cli.Command{
				{
					Name:    "start",
					Aliases: []string{"s"},
					Usage:   "",
					Action: func(c *cli.Context) error {
						fmt.Println("server")
						return nil
					},
				},
			},
		},
	}

	app.Action = command.ActionInteractive

	app.Run(os.Args)
}
