//go:build android
// +build android

package rl

/*
#cgo android LDFLAGS: -llog -landroid -lEGL -lGLESv2 -lOpenSLES -lm
#cgo android CFLAGS: -I${SRCDIR}/../external/raylib/projects/VS2019-Android/raylib_android/raylib_android.NativeActivity
#cgo android CFLAGS: -DPLATFORM_ANDROID -DPLATFORM_ANDROID_GOLANG -DGRAPHICS_API_OPENGL_ES2 -Iexternal/android/native_app_glue -Wno-implicit-const-int-float-conversion

#cgo android,arm CFLAGS: -march=armv7-a -mfloat-abi=softfp -mfpu=vfpv3-d16

//#include "android_native_app_glue.c"
*/
import "C"
