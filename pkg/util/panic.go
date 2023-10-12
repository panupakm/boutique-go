package util

import "runtime"

func IdentifyPanic() string {
	stackBuf := make([]byte, 1024)
	length := runtime.Stack(stackBuf, false)
	return string(stackBuf[:length])
}
