package shogi

import "github.com/pkg/errors"

// Record is "kifu" in japanese.
type Record struct {
	StartingSetup *State
	Actions       []Action
	// TODO: meta data
}

func (r *Record) GetStateAfterMoves(nth int) (*State, error) {
	s := r.StartingSetup.Clone()
	ms := r.Moves()
	if nth < 0 || len(ms) < nth {
		return nil, errors.Errorf("nth is out-of-range: %d", nth)
	}
	for i, m := range ms {
		if i >= nth {
			break
		}
		if err := s.ApplyMove(m); err != nil {
			return nil, err
		}
	}
	return s, nil
}

func (r *Record) GetLatestState() (*State, error) {
	s := r.StartingSetup.Clone()
	ms := r.Moves()
	for _, m := range ms {
		if err := s.ApplyMove(m); err != nil {
			return nil, err
		}
	}
	return s, nil
}

// Moves returns move actions in Actions.
func (r *Record) Moves() []*Move {
	ms := make([]*Move, 0, len(r.Actions))
	for _, a := range r.Actions {
		if m, ok := a.(*Move); ok {
			ms = append(ms, m)
		}
	}
	return ms
}

func (r *Record) String() string {
	return "TODO"
}

// type Recorder interface {
// 	Record() *Record
// 	State() State
// 	Nth() int
// 	Next() error
// 	Prev() error
// 	First()
// 	Last()
// 	Seek(nth int) error
// }
//
// type Player struct {
// 	FirstName string
// 	LastName  string
// }
//
// type PlayerPair struct {
// 	Black Player
// 	White Player
// }
//
// type Record struct {
// 	// StartedAt  time.Time
// 	// Tournament string
// 	// Location   string
// 	// Players    PlayerPair
// 	// Meta       map[string]string
// 	Moves []Move
// 	Game  *Game
// }
//
// type RecordLoader interface {
// 	Load(filePath string) (*Record, error)
// }
//
// type RecordFormat int
//
// // RecordType list
// const (
// 	RecordFormatNull RecordFormat = iota
// 	KI2
// 	KIF
// 	CSA
// )
