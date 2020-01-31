package mlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// FuncName -
func FuncName() string {
	return funcName(2)
}

func funcName(l int) string {
	pc, file, line, _ := runtime.Caller(l)
	longFuncName := runtime.FuncForPC(pc).Name()
	funcname := longFuncName[strings.Index(longFuncName, ".")+1:] // skip package name
	_, filename := filepath.Split(file)                           // no need dir
	return fmt.Sprint(filename, ":", line, ":", funcname)
}

// LogInfo -
func LogInfo(a ...interface{}) string {
	// It is easier to comment here if you don't need this output than commenting all thru all sourses.
	res := funcName(2) + ": info: " + fmt.Sprint(a...)
	log.Print(res)
	return res
}

func logErr(a ...interface{}) string {
	res := funcName(3) + ": err: " + fmt.Sprint(a...)
	log.Println(res)
	return res
}

// LogError -
func LogError(a ...interface{}) string {
	return logErr(a...)
}

// LogIfErr -
func LogIfErr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	logErr(err, a)
}

// LogFatalIfErr -
func LogFatalIfErr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	logErr(err, a)
	os.Exit(1)
}
