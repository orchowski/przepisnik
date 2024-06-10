package server

import (
	"fit.synapse/przepisnik/logger"
	"net/http"
	"strconv"
)

func Start(port int, endpoints map[string]Handler) {
	for path, handler := range endpoints {
		http.HandleFunc(path, handler.build())
	}

	logger.Log(logger.INFO, "Starting server on :%d", port)
	if err := http.ListenAndServe(":"+strconv.Itoa(port), nil); err != nil {
		logger.Log(logger.ERROR, "Could not start server: %s\n", err)
	}
}
