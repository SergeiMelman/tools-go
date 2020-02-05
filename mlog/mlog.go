package mlog

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

// Verbose - false = do not print out log information.
var Verbose = true

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
	if Verbose {
		log.Print(res)
	}
	return res
}

// LogError -
func LogError(a ...interface{}) string {
	res := funcName(2) + ":err:" + fmt.Sprint(a...)
	if Verbose {
		log.Println(res)
	}
	return res
}

// LogIfErr -
func LogIfErr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	res := funcName(2) + ":err:" + err.Error() + ":" + fmt.Sprint(a...)
	if Verbose {
		log.Println(res)
	}
}

// LogFatalIfErr -
func LogFatalIfErr(err error, a ...interface{}) {
	if err == nil {
		return
	}
	res := funcName(2) + ":err:" + err.Error() + ":" + fmt.Sprint(a...)
	if Verbose {
		log.Println(res)
	}
	os.Exit(1)
}
