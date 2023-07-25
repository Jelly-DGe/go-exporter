package main

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-exporter/test"
	"net/http"
)

func main() {
	go test.Test()
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":8080", nil)
}
