package util

import "runtime"

func GetFuncName() string {
	pc, _, _, ok := runtime.Caller(1)
	if ok {
		return runtime.FuncForPC(pc).Name()
	}
	return ""
}
