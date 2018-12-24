package gsr

type Logger interface {
	Infof(format string, v ...interface{})
	Fatalf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}

const (
	LoggerLevelInfo  = iota
	LoggerLevelFatal
	LoggerLevelPanic
)
