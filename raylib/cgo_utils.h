#if defined(__cplusplus)
extern "C" {            // Prevents name mangling of functions
#endif


void SetTraceLogCallbackWrapper(void);                 // enable the call-back
void TraceLogWrapper(int logLevel, const char *text);
void SetLoadFileDataCallbackWrapper(void);
void SetLoadFileTextCallbackWrapper(void);

#if defined(__cplusplus)
}
#endif
