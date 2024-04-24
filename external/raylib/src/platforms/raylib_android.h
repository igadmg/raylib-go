#pragma once

#if !defined(PLATFORM_ANDROID_GOLANG)
#include <android_native_app_glue.h>    // Required for: android_app struct and activity management
#else
#include <android/configuration.h>
#include <android/input.h>
#include <android/native_activity.h>

struct android_app {
        ANativeActivity* activity;
        AConfiguration* config;
};

void AndroidInitWindow(ANativeWindow* window);
void AndroidGainedFocus();
void AndroidLostFocus();
void AndroidTerminateWindow();
#endif
