package main

type SupportedGamesMode int64

const (
	Unknown SupportedGamesMode = iota
	Standard
	Royale
	Wrapped
)

var GAME_MODE SupportedGamesMode

func (s SupportedGamesMode) String() string {
	switch s {
	case Standard:
		return "standard"
	case Royale:
		return "royale"
	case Wrapped:
		return "wrapped"
	}
	return "unknown"
}

func ReadMode(s string) SupportedGamesMode {
	switch s {
	case "standard":
		return Standard
	case "royale":
		return Royale
	case "wrapped":
		return Wrapped
	}
	return Unknown
}
