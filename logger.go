package log

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// RFC5424 log message levels.
const (
	LevelFatal   = 0
	LevelError   = 1
	LevelWarning = 2
	LevelInfo    = 3
	LevelDebug   = 4
)

var level = [LevelDebug + 1]string{"FATAL", "ERROR", "WARNING", "INFO", "DEBUG"}

type Logger struct {
	level               int
	enableFuncCallDepth bool
	loggerFuncCallDepth int
	colorful            bool
}

func (l *Logger) Fatal(v ...interface{}) {
	format := formatLog(v...)
	if LevelFatal > l.level {
		return
	}
	l.writeMsg(LevelFatal, format, v...)
}

func (l *Logger) Error(v ...interface{}) {
	format := formatLog(v...)
	if LevelError > l.level {
		return
	}
	l.writeMsg(LevelError, format, v...)
}

func (l *Logger) Warning(v ...interface{}) {
	format := formatLog(v...)
	if LevelWarning > l.level {
		return
	}
	l.writeMsg(LevelWarning, format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	format := formatLog(v...)
	if LevelInfo > l.level {
		return
	}
	l.writeMsg(LevelInfo, format, v...)
}

func (l *Logger) Debug(v ...interface{}) {
	format := formatLog(v...)
	if LevelDebug > l.level {
		return
	}
	l.writeMsg(LevelDebug, format, v...)
}

func (l *Logger) writeMsg(logLevel int, msg string, v ...interface{}) error {
	if len(v) > 0 {
		msg = fmt.Sprintf(msg, v...)
	}
	track := ""
	if l.enableFuncCallDepth {
		for i := 0; i < l.loggerFuncCallDepth; i++ {
			_, file, line, ok := runtime.Caller(3 + i)
			if !ok {
				break
			}
			track = track + " [" + file + ":" + strconv.FormatInt(int64(line), 10) + "]"
		}
	}

	msgJson, _ := json.Marshal(map[string]string{
		"level": level[logLevel],
		"time":  time.Now().Format(time.RFC850),
		"msg":   msg,
		"track": track,
	})

	if l.colorful {
		msg = colors[logLevel](string(msgJson))
	}
	os.Stdout.Write(append(msgJson, '\n'))
	return nil
}

func formatLog(v ...interface{}) string {
	return strings.Repeat("%v ", len(v))
}
