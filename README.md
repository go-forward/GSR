# GSR
Moving Go forward through collaboration and standards.

## GSR-3 Logger

```Go
type LoggerInterface interface {
	Debugf(format string, v ...interface{})
	Infof(format string, v ...interface{})
	Noticef(format string, v ...interface{})
	Warningf(format string, v ...interface{})
	Errorf(format string, v ...interface{})
	Criticalf(format string, v ...interface{})
	Alertf(format string, v ...interface{})
	Emergencyf(format string, v ...interface{})
	Panicf(format string, v ...interface{})
}

type LoggerAwareInterface interface {
	SetLogger(logger LoggerInterface)
}

const (
	LoggerLevelDebug     = iota
	LoggerLevelInfo
	LoggerLevelNotice
	LoggerLevelWarning
	LoggerLevelError
	LoggerLevelCritical
	LoggerLevelAlert
	LoggerLevelEmergency
	LoggerLevelPanic
)

```


### GSR-11 Container

```Go
type ContainerInterface interface {
	Get(id interface{}) interface{}
	Has(id interface{}) (bool, error)
}

var (
	ContainerError                 = fmt.Errorf("container error")
	ContainerErrorInstanceNotFound = fmt.Errorf("instance was not found")
)
```