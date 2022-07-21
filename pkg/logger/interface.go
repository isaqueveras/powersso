package logger

// Logger methods interface
type Logger interface {
	InitLogger()
	Debug(args ...interface{})
	Debugf(template string, args ...interface{})
	Info(args ...interface{})
	Infof(template string, args ...interface{})
	Warn(args ...interface{})
	Warnf(template string, args ...interface{})
	Error(args ...interface{})
	Errorf(template string, args ...interface{})
	DPanic(args ...interface{})
	DPanicf(template string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(template string, args ...interface{})
}

func (l *logger) Debug(args ...interface{}) {
	l.sugaredLogger.Debug(args...)
}

func (l *logger) Debugf(template string, args ...interface{}) {
	l.sugaredLogger.Debugf(template, args...)
}

func (l *logger) Info(args ...interface{}) {
	l.sugaredLogger.Info(args...)
}

func (l *logger) Infof(template string, args ...interface{}) {
	l.sugaredLogger.Infof(template, args...)
}

func (l *logger) Warn(args ...interface{}) {
	l.sugaredLogger.Warn(args...)
}

func (l *logger) Warnf(template string, args ...interface{}) {
	l.sugaredLogger.Warnf(template, args...)
}

func (l *logger) Error(args ...interface{}) {
	l.sugaredLogger.Error(args...)
}

func (l *logger) Errorf(template string, args ...interface{}) {
	l.sugaredLogger.Errorf(template, args...)
}

func (l *logger) DPanic(args ...interface{}) {
	l.sugaredLogger.DPanic(args...)
}

func (l *logger) DPanicf(template string, args ...interface{}) {
	l.sugaredLogger.DPanicf(template, args...)
}

func (l *logger) Panic(args ...interface{}) {
	l.sugaredLogger.Panic(args...)
}

func (l *logger) Panicf(template string, args ...interface{}) {
	l.sugaredLogger.Panicf(template, args...)
}

func (l *logger) Fatal(args ...interface{}) {
	l.sugaredLogger.Fatal(args...)
}

func (l *logger) Fatalf(template string, args ...interface{}) {
	l.sugaredLogger.Fatalf(template, args...)
}
