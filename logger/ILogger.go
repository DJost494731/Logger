package logger

import "fmt"

var logger ILogger

type ILogger interface {
	Info(message string)
	Warning(message string)
	Error(message string)

	Log(logType string, message string)
}

//Logger gets an instance of the logger.
func Logger() ILogger {
	if logger == nil {
		initializeDefaultLogger()
	}
	return logger
}

//UseLogger sets the logger type.
func UseLogger(loggerToUse ILogger) {
	logger = loggerToUse
	loggerType := fmt.Sprintf("%T", loggerToUse)
	logger.Info("Logger has been changed successfully! Now using '" + loggerType + "'.")
}

func initializeDefaultLogger() {
	logger = TerminalLogger{}
	logger.Info("No Logger has been specified; using TerminalLogger by default.")
}

type MockLogger struct {
	logMessage *string
}

func (m MockLogger) Error(msg string){}
func (m MockLogger) Info(msg string){}
func (m MockLogger) Warning(msg string){}
func (m MockLogger) Log(logtype string, msg string) { *m.logMessage = msg }