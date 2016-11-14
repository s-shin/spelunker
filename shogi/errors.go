package shogi

import "fmt"

// MoveError is used when an error occurs in moving.
type MoveError string

// NewMoveError is the constructor of MoveError.
func NewMoveError(m Move) MoveError {
	return MoveError(fmt.Sprintf("Failed to move: %v", m))
}

func (m MoveError) Error() string {
	return string(m)
}

//---

// InvalidStateError is used when some inconsistency occurs in game.
type InvalidStateError string

// NewInvalidStateError is the constructor of NewInvalidStateError.
func NewInvalidStateError(desc string) InvalidStateError {
	return InvalidStateError("Invalid state: " + desc)
}

func (m InvalidStateError) Error() string {
	return string(m)
}
