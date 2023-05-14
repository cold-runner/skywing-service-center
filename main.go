package main

import (
	"github.com/cold-runner/skywing-service-center/internal"
)

func main() {
	app := internal.NewApp()

	app.PrepareRun().Run()
}
