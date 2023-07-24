package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-exporter/counter"
	"net/http"
)

func main() {
	go counter.Inc()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
