package gauge

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

//无指标写法
var gauge = prometheus.NewGauge(
	prometheus.GaugeOpts{
		Name: "jelly_test_gauge",
		Help: "测试gauge"},
)

func init() {
	prometheus.MustRegister(gauge)
}

func Test() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64()
	gauge.Set(randomNumber)
}
