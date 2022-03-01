package msgcenter

import "zjh/Test/apitest"

type HandleFunc func(uint32, []byte) interface{}

type APIInfo struct {
	APIFunction HandleFunc
}

const (
	TestFunc = 0
)

var ApiMap = map[uint32]*APIInfo{
	TestFunc: &APIInfo{APIFunction: apitest.TestFunc},
}
