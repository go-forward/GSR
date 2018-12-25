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


### GSR-16 SimpleCache

```Go
type CacheInterface interface {
	Get(key, def interface{}) (interface{}, error)
	Set(key, val interface{}, ttl time.Duration) (bool, error)
	Clear() (bool, error)
	GetMultiple(keys, defs []interface{}) ([]interface{}, error)
	SetMultiple(vals map[interface{}]interface{}, ttl time.Duration) (bool, error)
	DeleteMultiple(keys []interface{}) (bool, error)
	Has(key interface{}) (bool, error)
}

var (
	SimpleCacheError                = fmt.Errorf("simplecache error")
	SimpleCacheErrorInvalidArgument = fmt.Errorf("simplecache invalid argument")
)
```


### GSR-15 HTTP Server Request Handlers

```Go

type RequestHandleFunc func(w http.ResponseWriter, r *http.Request)


type MiddlewareInterface interface {
	Process(next RequestHandleFunc) RequestHandleFunc
}

```