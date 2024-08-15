package main

import (
	factory "fit.synapse/przepisnik/app/factory"
)

func main() {
	factory.
		BuildApp().
		WithPort(7001).
		Start()
}
