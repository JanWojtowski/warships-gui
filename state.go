package gui

type State string

const (
	Empty State = "Empty"
	Hit   State = "Hit"
	Miss  State = "Miss"
	Ship  State = "Ship"
	Sunk  State = "Sunk"
)
