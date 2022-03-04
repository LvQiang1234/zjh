package tool

import (
	"math/rand"
	"time"
	"zjh/log"
)

type (
	Rand struct {
		*rand.Rand
	}
)

func (this *Rand) RandI(i int, n int) int {
	if i > n {
		log.Debug("随机数获取失败，区间设置错误")
		return i
	}

	return int(i + this.Int()%(n-i+1))
}

func (this *Rand) RandF(i float32, n float32) float32 {
	if i > n {
		log.Debug("随机数获取失败，区间设置错误")
		return i
	}

	return (i + (n-i)*this.Float32())
}

var RAND = &Rand{rand.New(rand.NewSource(time.Now().UnixNano()))}
