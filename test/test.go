package test

import (
	"go-exporter/counter"
	"go-exporter/gauge"
	"go-exporter/histogram"
	"time"
)

//用于生成测试数据
func Test() {
	time.Sleep(10 * time.Second)
	for {
		//counter测试数据
		counter.Test()
		gauge.Test()
		histogram.Test()
		time.Sleep(5 * time.Second)
	}
}
