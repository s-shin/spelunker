package shogi

// Action in a record.
type Action interface {
	// ActionType function is the signature of Action.
	ActionType() string
}
