package logger

import "fmt"

type Logger interface {
	Log(args ...interface{})
}

type StdLogger struct {
}

func (l *StdLogger) Log(args ...interface{}) {
	fmt.Printf(args[0].(string)+"\n", args[1:]...)
}
