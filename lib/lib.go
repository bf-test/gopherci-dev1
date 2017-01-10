package lib

// Logger is the interface to support swappable loggers.
type Logger interface {
	Info(string)
}

func NewBreaking() {}
