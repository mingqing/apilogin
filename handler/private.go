package handler

import (
	"net/http"

	"google.golang.org/grpc"
)

func (m *Microservice) privateExtended() error {
	return nil
}

func (m *Microservice) privateUnaryServerInterceptor() []grpc.UnaryServerInterceptor {
	return nil
}

func (m *Microservice) privateHTTPHandle(mux *http.ServeMux) {
	/*
		mux.HandleFunc("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, _ = fmt.Fprintf(w, "Hello Admin")
		})
	*/
}