package counter

import (
	"github.com/prometheus/client_golang/prometheus"
)

// label1, label2是两个非固定指标指标,但是必须同时增加
var counter1 = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "jelly_test_counter1",
		Help: "测试counter1"},
	[]string{"label1", "label2"})

var counter2 = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "jelly_test_counter2",
		Help: "测试counter2"},
	[]string{"label3", "label4"})

func init() {
	prometheus.MustRegister(counter1)
	prometheus.MustRegister(counter2)
}

func Test() {
	counter1.With(map[string]string{"label1": "test1", "label2": "test2"}).Inc()
	counter2.WithLabelValues("test3", "test4").Add(2)
}
