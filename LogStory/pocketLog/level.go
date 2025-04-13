package pocketLog

// Level represents a log level
type Level byte

const (
	//LevelDebug is the debug log level
	LevelDebug Level = iota

	//LevelInfo is the info log level
	LevelInfo

	//LevelError is the error log level
	LevelError
)
