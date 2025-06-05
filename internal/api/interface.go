package api

import "net/http"

type Api interface {
	Routes() http.Handler
	Register(mux *http.ServeMux)
}
