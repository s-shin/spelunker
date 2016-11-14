package shogi

// Side of players.
type Side int

// Side constants.
const (
	SideNull Side = iota
	Black
	White
)

type sideStrings struct {
	CSA string
}

var sideStringsMap = map[Side]sideStrings{
	SideNull: sideStrings{
		CSA: " ",
	},
	Black: sideStrings{
		CSA: "+",
	},
	White: sideStrings{
		CSA: "-",
	},
}

func (side Side) String() string {
	return sideStringsMap[side].CSA
}

// Reverse the side.
func (side Side) Reverse() Side {
	switch side {
	case Black:
		return White
	case White:
		return Black
	}
	return SideNull
}
