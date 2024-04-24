#if defined(__cplusplus)
extern "C" {            // Prevents name mangling of functions
#endif


void setLogCallbackWrapper(void);                 // enable the call-back
void TraceLogWrapper(int logLevel, const char *text);

#if defined(__cplusplus)
}
#endif
