package rl

/*
#include "raylib.h"
*/
import "C"

// SetGesturesEnabled - Enable a set of gestures using flags
func SetGesturesEnabled(gestureFlags uint32) {
	cgestureFlags := (C.uint)(gestureFlags)
	C.SetGesturesEnabled(cgestureFlags)
}

// IsGestureDetected - Check if a gesture have been detected
func IsGestureDetected(gesture Gestures) bool {
	cgesture := (C.uint)(gesture)
	ret := C.IsGestureDetected(cgesture)
	return bool(ret)
}

// GetGestureDetected - Get latest detected gesture
func GetGestureDetected() Gestures {
	ret := C.GetGestureDetected()
	return (Gestures)(ret)
}

func GetGestureTapPosition() Vector2 {
	ret := C.GetGestureTapPosition()
	return *govec2ptr(&ret)
}

// GetGestureHoldDuration - Get gesture hold time in milliseconds
func GetGestureHoldDuration() float32 {
	ret := C.GetGestureHoldDuration()
	return (float32)(ret)
}

// GetGestureSwipeDistance - Get gesture swipe angle
func GetGestureSwipeDistance() float32 {
	ret := C.GetGestureSwipeDistance()
	return (float32)(ret)
}

// GetGestureSwipeIntensity - Get gesture swipe angle
func GetGestureSwipeIntensity() float32 {
	ret := C.GetGestureSwipeIntensity()
	return (float32)(ret)
}

// GetGestureSwipeAngle - Get gesture swipe angle
func GetGestureSwipeAngle() float32 {
	ret := C.GetGestureSwipeAngle()
	return (float32)(ret)
}

// GetGestureDragVector - Get gesture drag vector
func GetGestureDragVector() Vector2 {
	ret := C.GetGestureDragVector()
	return *govec2ptr(&ret)
}

// GetGestureDragAngle - Get gesture drag vector
func GetGestureDragAngle() float32 {
	ret := C.GetGestureDragAngle()
	return (float32)(ret)
}

// GetGesturePinchVector - Get gesture pinch delta
func GetGesturePinchVector() Vector2 {
	ret := C.GetGesturePinchVector()
	return *govec2ptr(&ret)
}

// GetGesturePinchAngle - Get gesture pinch angle
func GetGesturePinchAngle() float32 {
	ret := C.GetGesturePinchAngle()
	v := (float32)(ret)
	return v
}
