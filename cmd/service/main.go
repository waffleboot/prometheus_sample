package main

import (
	"net/http"

	"github.com/waffleboot/prometheus_sample/internal/app/service"
)

func main() {
	http.HandleFunc("/service", service.NewService())
	http.ListenAndServe(":8000", nil)
}
