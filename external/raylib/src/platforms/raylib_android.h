#pragma once

#include "android_native_app_glue.h"    // Required for: android_app struct and activity management
#if defined(PLATFORM_ANDROID_GOLANG)
RLAPI void RayLibANativeActivity_onCreate(ANativeActivity* activity, void* savedState, size_t savedStateSize);
#endif

RLAPI struct android_app *GetAndroidApp(void);
