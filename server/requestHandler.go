package server

import (
	"fit.synapse/przepisnik/logger"
	"net/http"
)

type HandlerFunc struct {
	Do http.HandlerFunc
}

type Handler struct {
	Get    *HandlerFunc
	Post   *HandlerFunc
	Put    *HandlerFunc
	Delete *HandlerFunc
}

func (h *Handler) build() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Log(logger.INFO, "Handling request : %s %s", r.Method, r.URL.Path)
		switch r.Method {
		case http.MethodGet:
			checkAndPass(h.Get, r.Method, r.URL.Path).ServeHTTP(w, r)
		case http.MethodPost:
			checkAndPass(h.Post, r.Method, r.URL.Path).ServeHTTP(w, r)
		case http.MethodPut:
			checkAndPass(h.Put, r.Method, r.URL.Path).ServeHTTP(w, r)
		case http.MethodDelete:
			checkAndPass(h.Delete, r.Method, r.URL.Path).ServeHTTP(w, r)
		default:
			logger.Log(logger.ERROR, "Handling request failure: %s", r.URL.Path)
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

		}
	}
}

func checkAndPass(handlerFunc *HandlerFunc, method string, path string) http.HandlerFunc {
	if handlerFunc != nil {
		return handlerFunc.Do
	}
	logger.Log(logger.WARNING, "wrong path requested for method %s %s", method, path)
	return func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Not found, sorry :(", http.StatusNotFound)
	}
}
