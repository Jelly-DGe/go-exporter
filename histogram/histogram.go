package histogram

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"time"
)

//固定指标写法
var histogram = prometheus.NewHistogram(
	prometheus.HistogramOpts{
		Name:        "jelly_test_histogram",
		Help:        "测试histogram",
		Buckets:     prometheus.LinearBuckets(0, 3, 6),
		ConstLabels: prometheus.Labels{"label1": "test1"}},
)

func init() {
	prometheus.MustRegister(histogram)
}

func Test() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Float64() * 17
	histogram.Observe(randomNumber)
}
