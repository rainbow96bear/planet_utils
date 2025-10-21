package logger

import (
	"fmt"
	"path/filepath"
	"runtime"
	"time"
)

// 로그 레벨 정의
const (
	DEBUG int16 = iota
	INFO
	WARN
	ERROR
)

// 현재 로그 레벨 (이 레벨 이상만 출력)
var CurrentLevel int16 = DEBUG

// 로그 메시지를 출력하는 내부 함수
func logMessage(level int16, prefix string, msg string, a ...interface{}) {
	if level < CurrentLevel {
		return
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	formattedMsg := fmt.Sprintf(msg, a...)

	// 호출한 함수 정보 가져오기
	pc, file, line, ok := runtime.Caller(2)
	funcName := "unknown"
	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			funcName = fn.Name()
			// 패키지명 제외하고 함수명만 남기기
			if lastSlash := filepath.Base(funcName); lastSlash != "" {
				funcName = lastSlash
			}
		}
	} else {
		file = "???"
	}

	fileName := filepath.Base(file)

	// 로그 레벨별 색상
	color := ""
	reset := "\033[0m" // 색상 초기화
	switch level {
	case DEBUG:
		color = "\033[36m" // 청록색
	case INFO:
		color = "\033[32m" // 초록색
	case WARN:
		color = "\033[33m" // 노랑색
	case ERROR:
		color = "\033[31m" // 빨강색
	}

	// 함수 이름까지 출력
	fmt.Printf("%s%s [%s] [%s:%d %s()] %s%s\n",
		color, timestamp, prefix, fileName, line, funcName, formattedMsg, reset)
}

// 로그 레벨별 함수
func Debugf(msg string, a ...interface{}) { logMessage(DEBUG, "DEBUG", msg, a...) }
func Infof(msg string, a ...interface{})  { logMessage(INFO, "INFO", msg, a...) }
func Warnf(msg string, a ...interface{})  { logMessage(WARN, "WARN", msg, a...) }
func Errorf(msg string, a ...interface{}) { logMessage(ERROR, "ERROR", msg, a...) }

// 로그 레벨 설정 함수
func SetLevel(level int16) {
	CurrentLevel = level
}
