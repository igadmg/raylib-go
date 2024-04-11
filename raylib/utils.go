//go:build !android
// +build !android

package rl

/*
#include "stdlib.h"
#include "raylib.h"
void TraceLogWrapper(int logLevel, const char *text)
{
	TraceLog(logLevel, text);
}
*/
import "C"

import (
	"fmt"
	"os"
)

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

// HomeDir - Returns user home directory
// NOTE: On Android this returns internal data path and must be called after InitWindow
func HomeDir() string {
	if homeDir, err := os.UserHomeDir(); err == nil {
		return homeDir
	}
	return ""
}
