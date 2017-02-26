package main

import (
	"fmt"
	"os"

	"github.com/chzyer/readline"
	"github.com/s-shin/spelunker/shogi/script"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "spelunker"
	app.Usage = "spelunker"
	app.Version = "0.1.0"

	app.Commands = []cli.Command{
		{
			Name:    "interactive",
			Aliases: []string{"i"},
			Usage:   "Run as interactive mode.",
			Action: func(c *cli.Context) error {
				rl, err := readline.New("> ")
				if err != nil {
					return err
				}
				defer rl.Close()

				runner := script.NewRunner()
				for {
					line, err := rl.Readline()
					if err != nil { // io.EOF
						break
					}
					rs, err := runner.RunLines(line)
					if err != nil {
						fmt.Println(err.Error())
						continue
					}
					for _, r := range rs {
						if r != "" {
							fmt.Println(r)
						}
					}
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
