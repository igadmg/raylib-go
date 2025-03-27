#include "../external/raylib/src/utils.c"

#include "cgo_utils.h"
#include <stdio.h>                      // Required for: vprintf()
#include <string.h>                     // Required for: strcpy(), strcat()

#define MAX_TRACELOG_BUFFER_SIZE   128  // As defined in utils.c from raylib

extern void internalTraceLogCallback(int logType, void *text);
extern unsigned char *internalLoadFileDataCallback(const char *fileName, int *dataSize);    // FileIO: Load binary data
//extern bool internalSaveFileDataCallback(const char *fileName, void *data, int dataSize); // FileIO: Save binary data
extern char *internalLoadFileTextCallback(const char *fileName);                            // FileIO: Load text data
//extern bool internalSaveFileTextCallback(const char *fileName, char *text);               // FileIO: Save text data

void rayLogWrapperCallback(int logType, const char *text, va_list args) {
	char buffer[MAX_TRACELOG_BUFFER_SIZE] = { 0 };

	vsprintf(buffer, text, args);

	internalTraceLogCallback(logType, buffer);
}

void SetTraceLogCallbackWrapper(void) {
	SetTraceLogCallback(rayLogWrapperCallback);
}

void TraceLogWrapper(int logLevel, const char *text)
{
	TraceLog(logLevel, text);
}

void SetLoadFileDataCallbackWrapper() {
	SetLoadFileDataCallback(internalLoadFileDataCallback);
}

void SetLoadFileTextCallbackWrapper() {
	SetLoadFileTextCallback(internalLoadFileTextCallback);
}