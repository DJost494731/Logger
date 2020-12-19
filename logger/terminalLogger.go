package logger

import "fmt"

type TerminalLogger struct { }

func (l TerminalLogger) Info(message string) {
	l.Log("Info", message)
}

func (l TerminalLogger) Warning(message string) {
	l.Log("Warning", message)
}

func (l TerminalLogger) Error(message string) {
	l.Log("Error", message)
}

func (l TerminalLogger) Log(logType string, message string) {
	fmt.Println(logType + ": " + message)
}
