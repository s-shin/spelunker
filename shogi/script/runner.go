package script

import (
	"strconv"

	"github.com/pkg/errors"
	"github.com/s-shin/spelunker/parsec"
	"github.com/s-shin/spelunker/shogi"
)

// Modes
const (
	ModeUnknown = iota
	ModeEdit
	ModePlay
)

// Versions
const (
	Version1      = 1 + iota
	VersionLatest = Version1
)

// Command list
const (
	CommandInit    = "init"
	CommandTitle   = "title"
	CommandShow    = "show"
	CommandMode    = "mode"
	CommandEdit    = "edit"
	CommandMove    = "move"
	CommandSave    = "save"
	CommandHelp    = "help"
	CommandVersion = "version"
)

// Help texts
const (
	Help = `Usage: <command> [arguments]

Available commands:

    init    initialize game
	title   set game title
    show    show information
    edit	edit state
    move    move piece
    save    save game
    help    show help text

Run 'help <command>' for more information.`

	HelpInit    = `Usage: init [state:hirate] [record:<path>]`
	HelpTitle   = `Usage: title <text>`
	HelpShow    = `Usage: show <target>`
	HelpEdit    = `Usage: edit <src> <dst>`
	HelpMove    = `Usage: move <side><from><to><piece>`
	HelpSave    = `Usage: save ...`
	HelpHelp    = `Usage: help <command>`
	HelpVersion = `Usage: version <number>`
)

var helpTexts = map[string]string{
	CommandInit:    HelpInit,
	CommandTitle:   HelpTitle,
	CommandShow:    HelpShow,
	CommandEdit:    HelpEdit,
	CommandMove:    HelpMove,
	CommandSave:    HelpSave,
	CommandVersion: HelpVersion,
}

type Runner struct {
	commands          map[string]CommandFunc
	game              *shogi.Game
	mode              int
	version           int
	commandLineParser parsec.Parser
}

type CommandFunc func(args []*CommandArg) (string, error)

func NewRunner() *Runner {
	s := &Runner{}

	s.commands = map[string]CommandFunc{
		CommandInit:    s.CommandInit,
		CommandTitle:   s.CommandTitle,
		CommandShow:    s.CommandShow,
		CommandEdit:    s.CommandEdit,
		CommandMove:    s.CommandMove,
		CommandSave:    s.CommandSave,
		CommandHelp:    s.CommandHelp,
		CommandVersion: s.CommandVersion,
	}

	s.CommandVersion([]*CommandArg{{"", strconv.Itoa(VersionLatest)}})

	return s
}

// CommandHelp is the implementation of help command.
func (s *Runner) CommandHelp(args []*CommandArg) (string, error) {
	if len(args) == 0 {
		return Help, nil
	}
	if len(args) != 1 {
		return "", errors.New("one argument required")
	}
	if help, ok := helpTexts[args[0].Value]; ok {
		return help, nil
	}
	return "", errors.New("no help text for " + args[0].Value)
}

// CommandVersion is the implementation of version command.
func (s *Runner) CommandVersion(args []*CommandArg) (string, error) {
	v, err := strconv.ParseUint(args[0].Value, 10, 32)
	if err != nil {
		return "", errors.Wrapf(err, "arg1 should be number: %s", err)
	}
	switch v {
	case Version1:
	default:
		return "", errors.Errorf("unsupported version: %d", v)
	}
	s.version = int(v)
	return s.CommandInit(nil)
}

// CommandInit is the implementation of init command.
func (s *Runner) CommandInit(args []*CommandArg) (string, error) {
	s.game = shogi.NewGameWithBoard(shogi.NewHirateBoard())
	s.mode = ModeEdit
	return "OK", nil
}

// CommandTitle is the implementation of title command.
func (s *Runner) CommandTitle(args []*CommandArg) (string, error) {
	return "TODO", nil
}

// CommandShow is the implementation of show command.
func (s *Runner) CommandShow(args []*CommandArg) (string, error) {
	if len(args) != 1 {
		return "", errors.New("one argument required")
	}
	if s.mode == ModeUnknown {
		return "", errors.New("not initialized")
	}
	var output string
	switch args[0].Value {
	case "state":
		state, err := s.game.State()
		if err != nil {
			return "", err
		}
		output = state.String()
	case "record":
		output = s.game.Record().String()
	default:
		return "", errors.New("unknown target: " + args[0].Value)
	}
	return output, nil
}

// CommandEdit is the implementation of edit command.
func (s *Runner) CommandEdit(args []*CommandArg) (string, error) {
	return "TODO", nil
}

// CommandMove is the implementation of move command.
func (s *Runner) CommandMove(args []*CommandArg) (string, error) {
	if len(args) != 1 {
		return "", errors.Errorf("invalid length of arguments: %d", len(args))
	}
	s := args[0].Value
	if len(s) != 7 {
		return "", errors.Errorf("")
	}
	var m shogi.Move
	switch s[0] {
	case '+':
		m.Side = shogi.Black
	case '-':
		m.Side = shogi.White
	default:
		return "", errors.Errorf("")
	}
	m.From := shogi.MakePositionFromString(s[1:2])
	if m.From == shogi.PositionNull {
		return "", errors.Errorf("")
	}
	m.To := shogi.MakePositionFromString(s[3:4])
	if m.To == shogi.PositionNull {
		return "", errors.Errorf("")
	}
	m.Piece = shogi.MakePieceFromString(s[5:6])
	if m.Piece == shogi.PieceNull {
		return "", errors.Errorf("")
	}
	
	return "TODO", nil
}

// CommandSave is the implementation of save command.
func (s *Runner) CommandSave(args []*CommandArg) (string, error) {
	return "TODO", nil
}

func (s *Runner) Run(cmd *Command) (string, error) {
	if fn, ok := s.commands[cmd.Name]; ok {
		return fn(cmd.Args)
	}
	return "", errors.New("unknown command: " + cmd.Name)
}

func (s *Runner) RunLines(input string) ([]string, error) {
	cmds, err := ParseShogiScript(input)
	if err != nil {
		return []string{}, err
	}
	rs := make([]string, 0, len(cmds))
	for _, c := range cmds {
		r, err := s.Run(c)
		if err != nil {
			return rs, err
		}
		rs = append(rs, r)
	}
	return rs, nil
}
