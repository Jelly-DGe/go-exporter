package summary

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

//不带指标写法
var summary = prometheus.NewSummary(
	prometheus.SummaryOpts{
		Name: "jelly_test_summary",
		Help: "测试histogram",
		//Objectives: map[float64]float64{0.5: 0.05, 0.90: 0.01, 0.99: 0.001},
	})

func init() {
	prometheus.MustRegister(summary)
}

func Test() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64() * 10
	summary.Observe(randomNumber)
}
