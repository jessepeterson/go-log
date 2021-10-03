package log

import (
	"testing"
)

type testLogger struct {
	infoKeyvals  map[string]interface{}
	debugKeyvals map[string]interface{}
}

func newTestLogger() *testLogger {
	return &testLogger{
		infoKeyvals:  make(map[string]interface{}),
		debugKeyvals: make(map[string]interface{}),
	}
}

func (t *testLogger) Info(keyvals ...interface{}) {
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			panic("expected key to be a string")
		}
		t.infoKeyvals[key] = keyvals[i+1]
	}
}

func (t *testLogger) Debug(keyvals ...interface{}) {
	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			panic("expected key to be a string")
		}
		t.infoKeyvals[key] = keyvals[i+1]
	}
}

func (t *testLogger) With(keyvals ...interface{}) Logger {
	logger := newTestLogger()
	logger.Info(keyvals...)
	logger.Debug(keyvals...)
	return logger
}

func TestMultiLogging(t *testing.T) {
	logger1 := newTestLogger()
	logger2 := newTestLogger()
	multi := NewMultiLogger(logger1, logger2)
	multi.Info("msg", "test")
	_, ok := logger1.infoKeyvals["msg"]
	if !ok {
		t.Errorf("msg key not present in first logger")
	}
	_, ok = logger2.infoKeyvals["msg"]
	if !ok {
		t.Errorf("msg key not present in second logger")
	}
}
