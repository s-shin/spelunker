package command

import (
	"fmt"
	"strings"

	"github.com/s-shin/spelunker/shogi"
	"github.com/urfave/cli"
	"gopkg.in/readline.v1"
)

type interactiveCommandContext struct {
	Game *shogi.Game
}

func newInteractiveCommandContext() *interactiveCommandContext {
	return &interactiveCommandContext{
		Game: shogi.NewGameWithStartingPositions(),
	}
}

func (c *interactiveCommandContext) CommandNew(args []string) {
	c.Game = shogi.NewGameWithStartingPositions()
	fmt.Println("New game was setup.")
}

func (c *interactiveCommandContext) CommandShow(args []string) {
	numArgs := len(args)
	if len(args) >= 2 {
		fmt.Println("Usage: show [type]")
		return
	}
	t := "game"
	if numArgs == 1 {
		t = args[0]
	}
	switch t {
	case "game":
		fmt.Println(c.Game.AppliedGame())
	case "moves":
		fmt.Println(c.Game.Moves)
	}
}

func (c *interactiveCommandContext) CommandMove(args []string) {
	numArgs := len(args)
	if numArgs != 2 && numArgs != 3 {
		fmt.Println("Usage: move <from> <to> [<promote>]")
		return
	}
	from := shogi.MakePositionFromString(args[0])
	if from == shogi.PositionNull {
		fmt.Printf("Invalid position string: %s\n", args[0])
		return
	}
	to := shogi.MakePositionFromString(args[1])
	if to == shogi.PositionNull {
		fmt.Printf("Invalid position string: %s\n", args[1])
		return
	}
	promote := false
	if numArgs == 3 && args[2] != "0" {
		promote = true
	}
	if g, err := c.Game.Move(from, to, promote); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(g)
	}
}

func (c *interactiveCommandContext) CommandDrop(args []string) {
	if len(args) != 2 || len(args[0]) != 2 || len(args[1]) != 2 {
		fmt.Println("Usage: drop <piece> <to>")
		return
	}
	piece := shogi.MakePieceFromString(args[0])
	if piece == shogi.PieceNull {
		fmt.Printf("Invalid piece string: %s\n", args[0])
		return
	}
	to := shogi.MakePositionFromString(args[1])
	if to == shogi.PositionNull {
		fmt.Printf("Invalid position string: %s\n", args[1])
		return
	}
	fmt.Println(piece, to)
	fmt.Println("TODO")
}

func ActionInteractive(c *cli.Context) error {
	rl, err := readline.New("> ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	cc := newInteractiveCommandContext()

	for {
		line, err := rl.Readline()
		if err != nil {
			break
		}
		fields := strings.Fields(line)
		command := fields[0]
		args := fields[1:]
		switch command {
		case "help":
			fmt.Printf(`Available commands:
- new
- show [type=game|record]
- move <from> <to> [promote=1|0]
- drop <from> <to>
`)
		case "new":
			cc.CommandNew(args)
		case "show":
			cc.CommandShow(args)
		case "move":
			cc.CommandMove(args)
		case "drop":
			cc.CommandDrop(args)
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
	return nil
}
