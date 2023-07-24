package counter

import (
	"github.com/prometheus/client_golang/prometheus"
	"time"
)

// label1, label2是两个指标
var counter1 = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gdd_test_counter1",
		Help: "测试counter1"},
	[]string{"label1", "label2"})

var counter2 = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "gdd_test_counter2",
		Help: "测试counter2"},
	[]string{"label3", "label4"})

func init() {
	prometheus.MustRegister(counter1)
	prometheus.MustRegister(counter2)
}

func Inc() {
	time.Sleep(10 * time.Second)
	for {
		counter1.With(map[string]string{"label1": "test1", "label2": "test2"}).Inc()
		counter2.With(map[string]string{"label3": "test3", "label4": "test4"}).Inc()
		time.Sleep(1 * time.Second)
	}
}
