package msgcenter

import "zjh/Test/apitest"

type HandleFunc func(uint64, []byte) interface{}

type APIInfo struct {
	APIFunction HandleFunc
}

const (
	TestFunc = 0
)

var ApiMap = map[uint16]*APIInfo{
	TestFunc: &APIInfo{APIFunction: apitest.TestFunc},
}
