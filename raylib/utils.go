package rl

/*
#include "stdlib.h"
#include "raylib.h"
#include "cgo_utils.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type LoadFileDataCallbackFn func(fileName string) []byte
type LoadFileTextCallbackFn func(fileName string) string

var internalTraceLogCallbackFn TraceLogCallbackFun
var internalLoadFileDataCallbackFn LoadFileDataCallbackFn
var internalLoadFileTextCallbackFn LoadFileTextCallbackFn

//export internalTraceLogCallback
func internalTraceLogCallback(logType C.int, ctext *C.char) {
	str := C.GoString(ctext)
	lt := int(logType)
	internalTraceLogCallbackFn(lt, str)
}

//export internalLoadFileDataCallback
func internalLoadFileDataCallback(cfileName *C.char, dataSize *C.int) unsafe.Pointer {
	fileName := C.GoString(cfileName)
	data := internalLoadFileDataCallbackFn(fileName)
	*dataSize = (C.int)(len(data))
	return C.CBytes(data)
}

//export internalLoadFileTextCallback
func internalLoadFileTextCallback(cfileName *C.char) *C.char {
	fileName := C.GoString(cfileName)
	data := internalLoadFileTextCallbackFn(fileName)
	return C.CString(data)
}

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
	internalTraceLogCallbackFn = fn
	C.SetTraceLogCallbackWrapper()
}

func SetLoadFileDataCallback(fn LoadFileDataCallbackFn) {
	internalLoadFileDataCallbackFn = fn
	C.SetLoadFileDataCallbackWrapper()
}

// RLAPI void SetSaveFileDataCallback(SaveFileDataCallback callback); // Set custom file binary data saver
func SetLoadFileTextCallback(fn LoadFileTextCallbackFn) {
	internalLoadFileTextCallbackFn = fn
	C.SetLoadFileTextCallbackWrapper()
}

//RLAPI void SetSaveFileTextCallback(SaveFileTextCallback callback); // Set custom file text data saver
