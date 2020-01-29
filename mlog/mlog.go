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
	pc, file, line, _ := runtime.Caller(2)
	longFuncName := runtime.FuncForPC(pc).Name()
	funcname := longFuncName[strings.Index(longFuncName, ".")+1:] // skip package name
	_, filename := filepath.Split(file)                           // no need dir
	return fmt.Sprint(filename, ":", line, ":", funcname)
}

// LogError -
func LogError(a ...interface{}) string {
	res := FuncName() + ": err: " + fmt.Sprint(a...)
	log.Println(res)
	return res
}

// LogInfo -
func LogInfo(a ...interface{}) string {
	// It is easier to comment here if you don't need this output than commenting all thru all sourses.
	res := FuncName() + ": info: " + fmt.Sprint(a...)
	log.Print(res)
	return res
}

// LogIfErr -
func LogIfErr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	log.Println(FuncName() + ": err: " + err.Error() + fmt.Sprint(a...))
}

// LogFatalIfErr -
func LogFatalIfErr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	log.Println(FuncName() + ": err: " + err.Error() + fmt.Sprint(a...))
	os.Exit(1)
}
