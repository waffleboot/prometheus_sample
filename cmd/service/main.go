package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/waffleboot/prometheus_sample/internal/app/service"
)

func main() {
	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/service", service.Handler())
	http.ListenAndServe(":8000", nil)
}
