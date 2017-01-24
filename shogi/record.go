package shogi

import "time"

type Player struct {
	FirstName string
	LastName  string
}

type PlayerPair struct {
	Black Player
	White Player
}

type Record struct {
	StartedAt  time.Time
	Tournament string
	Location   string
	Players    PlayerPair
	Meta       map[string]string
	Moves      []Move
	Game       *Game
}

type RecordLoader interface {
	Load(filePath string) (*Record, error)
}

type RecordFormat int

// RecordType list
const (
	RecordFormatNull RecordFormat = iota
	KI2
	KIF
	CSA
)
