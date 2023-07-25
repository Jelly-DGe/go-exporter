package test

import (
	"go-exporter/counter"
	"go-exporter/gauge"
	"time"
)

//用于生成测试数据
func Test() {
	time.Sleep(10 * time.Second)
	for {
		//counter测试数据
		counter.Test()
		gauge.Test()
		time.Sleep(5 * time.Second)
	}
}
