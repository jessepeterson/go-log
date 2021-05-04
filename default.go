package log

// Pacakge log is embedded (not imported) from:
// https://github.com/jessepeterson/go-log

// DefaultLogger is the default Logger used by the package funcs
var DefaultLogger Logger = NopLogger

// Info logs using the info level
func Info(i ...interface{}) {
	DefaultLogger.Info(i...)
}

// Debug logs using the debug level
func Debug(i ...interface{}) {
	DefaultLogger.Debug(i...)
}

// With nests the logger interface.
// Usually for adding logging context to a sub-logger.
func With(i ...interface{}) Logger {
	return DefaultLogger.With(i...)
}
