package log

var logger = &Logger{level: LevelError, enableFuncCallDepth: true, loggerFuncCallDepth: 3, colorful: true}

//LevelFatal   = 0
//LevelError   = 1
//LevelWarning = 2
//LevelInfo    = 3
//LevelDebug   = 4
func SetLevel(level int) {
	logger.level = level
}

func EnableFuncCallDepth(b bool) {
	logger.enableFuncCallDepth = b
}

func SetLogFuncCallDepth(d int) {
	logger.loggerFuncCallDepth = d
}

func EnableColorful(b bool) {
	logger.colorful = b
}

func Fatal(v ...interface{}) {
	logger.Fatal(v...)
}

func Error(v ...interface{}) {
	logger.Error(v...)
}

func Warning(v ...interface{}) {
	logger.Warning(v...)
}

func Info(v ...interface{}) {
	logger.Info(v...)
}

func Debug(v ...interface{}) {
	logger.Debug(v...)
}
