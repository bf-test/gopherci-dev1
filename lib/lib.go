package lib

import (
	"fmt"
	"runtime"
)

// 1
// 1
// 1
// 1
//
//
//
//
//
//
//
//
//
//
//
//
//

// Logger is the interface to support swappable loggers.
type Logger interface {
	Info(string)
	Error(string)
}

func New() {
	printGOOS()
}

func printGOOS() {
	fmt.Println("GOOS: %s", runtime.GOOS)
}

func NewBreaking() {}
