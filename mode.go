package common

type Mode int

const (
	DISABLE Mode = 1 >> 1
	DEV     Mode = 1 << 0
	REL     Mode = 1 << 1
	ALI     Mode = 1 << 2
)

var Modes = map[string]Mode{
	"disable": DISABLE,
	"dev":     DEV,
	"rel":     REL,
	"ali":     ALI,
}

func (m *Mode) String() string {
	for s, mode := range Modes {
		if mode == *m {
			return s
		}
	}
	return "unknown"
}
