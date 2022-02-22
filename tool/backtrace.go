package tool

import (
	"fmt"
	"runtime"
)

//输出错误，跟踪代码
func TraceCode(code ...interface{}) {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	data := ""
	for _, v := range code {
		data += fmt.Sprintf("%v", v)
	}
	data += string(buf[:n])
}
