package app

import (
	"fit.synapse/przepisnik/app"
	"fit.synapse/przepisnik/app/views/htmx"
	"fit.synapse/przepisnik/logger"
	"fit.synapse/przepisnik/server"
)

type ApplicationBuilder interface {
	WithBaseStoragePath(path string) ApplicationBuilder
	WithPort(port int) ApplicationBuilder
	Start() ApplicationBuilder
}

type applicationFactory struct {
	port     int
	basePath string
	started  bool
}

func (a *applicationFactory) WithBaseStoragePath(path string) ApplicationBuilder {
	a.basePath = path
	return a
}

func (a *applicationFactory) WithPort(port int) ApplicationBuilder {
	a.port = port
	return a
}

func BuildApp() ApplicationBuilder {
	return &applicationFactory{}
}

func (a *applicationFactory) Start() ApplicationBuilder {
	if a.started {
		return a
	}
	logger.NewLogger(logger.INFO)
	logger.NewLogger(logger.WARNING)
	logger.NewLogger(logger.ERROR)

	app := app.NewApp(a.basePath)
	app.Users.GetAll()
	handlers := make(map[string]server.Handler)
	htmx.RegisterHtmxControllers(handlers)
	server.Start(a.port, handlers)

	a.started = true
	return a
}
