package log

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