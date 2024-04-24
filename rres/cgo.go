package rres

/*
#cgo CFLAGS: -I${SRCDIR}/../external/raylib/src -I${SRCDIR}/../external/rres/src -std=gnu99 -Wno-unused-result -Wno-implicit-function-declaration -Wno-deprecated-declarations

#include "raylib.h"
*/
import "C"
import "unsafe"

func textAlloc(text string) *C.char {
	ctext := (*C.char)(unsafe.Pointer(unsafe.StringData(text)))
	clen := (C.int)(len(text))
	return C.TextAlloc(ctext, &clen)
}
