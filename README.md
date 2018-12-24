# GSR
Moving Go forward through collaboration and standards.

## GSR-3 Logger

```Go
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
```