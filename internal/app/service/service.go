package service

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Handler() http.HandlerFunc {
	opsProcessed := promauto.NewCounterVec(prometheus.CounterOpts{
		Name: "service_processed_ops_total",
		Help: "the total number of processed events",
	}, []string{"method"})
	return promhttp.InstrumentHandlerCounter(opsProcessed, http.HandlerFunc(handler))
}

func handler(w http.ResponseWriter, r *http.Request) {
}
