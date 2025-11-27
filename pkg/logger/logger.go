package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

const (
	DEBUG int16 = iota
	INFO
	WARN
	ERROR
)

var CurrentLevel int16 = DEBUG

func logMessage(level int16, msg string, a ...interface{}) {
	if level < CurrentLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02T15:04:05")
	formattedMsg := fmt.Sprintf(msg, a...)

	// 호출한 함수 정보 가져오기
	pc, file, line, ok := runtime.Caller(2)
	funcName := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			funcName = fn.Name()
			// 마지막 점(.) 이후 함수명만 사용
			if idx := lastDotIndex(funcName); idx != -1 {
				funcName = funcName[idx+1:]
			}
		}
	} else {
		file = "???"
	}
	fileName := filepath.Base(file)

	levelStr := ""
	switch level {
	case DEBUG:
		levelStr = "DEBUG"
	case INFO:
		levelStr = "INFO"
	case WARN:
		levelStr = "WARN"
	case ERROR:
		levelStr = "ERROR"
	}

	// 출력 포맷: [타임스탬프] [LEVEL] [파일:라인 함수()] 메시지
	fmt.Printf("%s [%s] [%s:%d %s()] %s\n",
		timestamp, levelStr, fileName, line, funcName, formattedMsg)
}

func lastDotIndex(s string) int {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			return i
		}
	}
	return -1
}

func Debugf(msg string, a ...interface{}) { logMessage(DEBUG, msg, a...) }
func Infof(msg string, a ...interface{})  { logMessage(INFO, msg, a...) }
func Warnf(msg string, a ...interface{})  { logMessage(WARN, msg, a...) }
func Errorf(msg string, a ...interface{}) { logMessage(ERROR, msg, a...) }

func SetLevel(level int16) { CurrentLevel = level }
