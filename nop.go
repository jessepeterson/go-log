package log

// nopLogger does nothing
type nopLogger struct{}

// Info does nothing
func (*nopLogger) Info(_ ...interface{}) {}

// Debug does nothing
func (*nopLogger) Debug(_ ...interface{}) {}

// With returns (the same) logger
func (logger *nopLogger) With(_ ...interface{}) Logger {
	return logger
}

// NopLogger is a Logger that does nothing
var NopLogger = &nopLogger{}
