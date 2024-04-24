package rl

/*
#include "stdlib.h"
#include "raylib.h"
#include "cgo_utils_log.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

var internalTraceLogCallbackFun TraceLogCallbackFun = func(int, string) {}

// SetTraceLogLevel - Set the current threshold (minimum) log level
func GetTraceLogLevel() TraceLogLevel {
	return (TraceLogLevel)(C.GetTraceLogLevel())
}

// SetTraceLogLevel - Set the current threshold (minimum) log level
func SetTraceLogLevel(logLevel TraceLogLevel) {
	clogLevel := (C.int)(logLevel)
	C.SetTraceLogLevel(clogLevel)
}

// TraceLog - Show trace log messages (LOG_DEBUG, LOG_INFO, LOG_WARNING, LOG_ERROR...)
func TraceLog(logLevel TraceLogLevel, text string, v ...interface{}) {
	ctext := textAlloc(fmt.Sprintf(text, v...))
	clogLevel := (C.int)(logLevel)
	C.TraceLogWrapper(clogLevel, ctext)
}

// SetTraceLogCallback - set a call-back function for trace log
func SetTraceLogCallback(fn TraceLogCallbackFun) {
	internalTraceLogCallbackFun = fn
	C.setLogCallbackWrapper()
}

//export internalTraceLogCallbackGo
func internalTraceLogCallbackGo(logType C.int, cstr unsafe.Pointer, length C.int) {
	str := string(C.GoBytes(cstr, length))
	lt := int(logType)
	internalTraceLogCallbackFun(lt, str)
}
