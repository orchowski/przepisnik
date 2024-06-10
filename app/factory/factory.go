package app

import "fit.synapse/przepisnik/server"
import "fit.synapse/przepisnik/logger"
import "net/http"
import "fmt"

type ApplicationBuilder interface {
	WithBaseStoragePath(path string) ApplicationBuilder
	WithPort(port int) ApplicationBuilder
	Start() ApplicationBuilder
}

type application struct {
	port     int
	basePath string
	started  bool
}

func (a *application) WithBaseStoragePath(path string) ApplicationBuilder {
	a.basePath = path
	return a
}

func (a *application) WithPort(port int) ApplicationBuilder {
	a.port = port
	return a
}

func BuildApp() ApplicationBuilder {
	return &application{}
}

func (a *application) Start() ApplicationBuilder {
	if a.started {
		return a
	}
	logger.NewLogger(logger.INFO)
	logger.NewLogger(logger.WARNING)
	logger.NewLogger(logger.ERROR)
	server.Start(
		a.port,
		map[string]server.Handler{
			"/test": {
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
			},
		},
	)

	a.started = true
	return a
}
