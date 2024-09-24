package htmx

import (
	"fmt"
	"net/http"
	"strings"

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
	handlers["/test/"] = server.Handler{
		Get: &server.HandlerFunc{
			Do: func(w http.ResponseWriter, r *http.Request) {
				// CHECK OUT WITH:  http://localhost:7001/test/1asdasdas?dupka=33&dupka=44

				id := strings.TrimPrefix(r.URL.Path, "/test/")
				queryParam := r.URL.Query()["dupka"] // if multiple "dupka" expected such thing returns array of strings, if we expect only one "dupka" it;d be Query().Get("dupka")
				fmt.Fprintf(w, "elo siema: "+id+": and query param: "+queryParam[0]+queryParam[1])
			},
		},
		Post: &server.HandlerFunc{
			Do: func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "elo siema wcale nie")
			},
		},
	}

}
