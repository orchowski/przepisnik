package main

import (
	factory "fit.synapse/przepisnik/app/factory"
)

func main() {
	factory.BuildApp().WithPort(7000).Start()
}
