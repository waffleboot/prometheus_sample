package main

import (
	"context"
	"io"
	"net/http"
	_ "net/http/pprof"

	gk "github.com/go-kit/kit/transport/http"
	service "github.com/waffleboot/prometheus_sample/internal/app/go_kit_sample"
)

func main() {
	handler := gk.NewServer(service.Handler(), requestDecoder, responseEncoder)
	http.Handle("/service", handler)
	http.ListenAndServe(":8000", nil)
}

func requestDecoder(_ context.Context, r *http.Request) (request interface{}, err error) {
	return r.URL.Query().Get("s"), nil
}

func responseEncoder(_ context.Context, w http.ResponseWriter, r interface{}) error {
	io.WriteString(w, r.(string))
	return nil
}
