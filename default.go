package log

// Pacakge log is embedded (not imported) from:
// https://github.com/jessepeterson/go-log

// DefaultLogger is the default Logger used by the package functions.
var DefaultLogger Logger = NopLogger

// Info logs using the info level to the default logger.
func Info(i ...interface{}) {
	DefaultLogger.Info(i...)
}

// Debug logs using the debug level to the default logger.
func Debug(i ...interface{}) {
	DefaultLogger.Debug(i...)
}

// With returns a new nested Logger from the default logger.
// Usually for adding logging additional context.
func With(i ...interface{}) Logger {
	return DefaultLogger.With(i...)
}
