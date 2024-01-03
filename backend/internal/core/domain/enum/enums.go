package enum

type Color int

const (
	Unspecified Color = iota
	Red
	Blue
	Green
)

//go:generate enumer -type=Color
