package service

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

func Handler() http.HandlerFunc {
	opsProcessed := promauto.NewCounter(prometheus.CounterOpts{
		Name: "service_processed_ops_total",
		Help: "the total number of processed events",
	})
	return func(w http.ResponseWriter, r *http.Request) {
		opsProcessed.Inc()
	}
}
