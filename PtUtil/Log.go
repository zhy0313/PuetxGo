package PtUtil

/**
 * Title:日志工具
 * User: iuoon
 * Date: 2016-9-22
 * Version: 1.0
 */

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	//TraceLevel Trace级别
	TraceLevel = iota
	//DebugLevel Debug级别
	DebugLevel
	//WarnLevel Warn级别
	WarnLevel
	//ErrorLevel Error级别
	ErrorLevel
	//InfoLevel Info级别
	InfoLevel
	//FatalLevel Fatal级别
	FatalLevel
)

var logger *log.Logger
var level = TraceLevel
var logFile *os.File
var isOutputScreen = true

//GetLevel 获取日志级别
func GetLevel() int {
	return level
}

//SetLevel 设置日志级别
func SetLevel(lev int) {
	if lev > FatalLevel || lev < TraceLevel {
		level = TraceLevel
	} else {
		level = lev
	}
}

//SetIsOutputScreen 是否输出到屏幕
func SetIsOutputScreen(isOutput bool) {
	isOutputScreen = isOutput
}

//InitLogger 日志模块初始化函数,程序启动时调用
func InitLogger(logFileName string) {
	var err error
	pID := os.Getpid()
	pIDStr := strconv.FormatInt(int64(pID), 10)
	if len(os.Args) == 1 {
		logFileName = "E:/" + logFileName + "_" + pIDStr + ".log"
	}else {
		logFileName = "log/" + logFileName + "_" + pIDStr + ".log"
	}
	logFile, err = os.OpenFile(logFileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Println(err.Error())
		return
	}

	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile)
	if isOutputScreen {
		Info("写日志初始化成功，日志输出路径：%s",logFileName)
	}
}

//Trace 不是很重要类型日志
func Trace(format string, v ...interface{}) {
	if level <= TraceLevel {
		var str string
		str = "[Trace] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			log.Println(str)
		}
	}
}

//Debug 调试类型日志
func Debug(format string, v ...interface{}) {
	if level <= DebugLevel {
		var str string
		str = "[Debug] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			log.Println(str)
		}
	}
}

//Warn 警告类型日志
func Warn(format string, v ...interface{}) {
	if level <= WarnLevel {
		var str string
		str = "[Warn] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			log.Println(str)
		}
	}
}

//Error 错误类型日志
func Error(format string, v ...interface{}) {
	if level <= ErrorLevel {
		var str string
		str = "[Error] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			log.Println(str)
		}
	}
}

//Info 程序信息类型日志
func Info(format string, v ...interface{}) {

	if level <= InfoLevel {
		var str string
		str = "[Info] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			log.Println(str)
		}
	}
}

//Fatal 致命错误类型日志
func Fatal(format string, v ...interface{}) {
	if level <= FatalLevel {
		var str string
		str = "[Fatal] " + format
		str = fmt.Sprintf(str, v...)
		logger.Output(2, str)

		if isOutputScreen {
			log.Println(str)
		}
	}
}
