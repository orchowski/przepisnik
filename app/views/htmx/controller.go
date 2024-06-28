package htmx

import (
	"fmt"
	"net/http"

	"fit.synapse/przepisnik/server"
)

func RegisterHtmxControllers(handlers map[string]server.Handler) {
	handlers["/test"] = server.Handler{
		Get: &server.HandlerFunc{
			Do: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "elo siema")
			},
		},
		Post: &server.HandlerFunc{
			Do: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "elo siema wcale nie")
			},
		},
	}

}
