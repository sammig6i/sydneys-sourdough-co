package main

import (
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	// TODO Update with API endpoints

}
