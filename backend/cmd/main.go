package main

import (
	"time"

	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
	"github.com/sammig6i/sydneys-sourdough-co/repository"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Postgres
	defer app.CloseDBConnection()

	repository.InitEmbeddingClient(app.EmbeddingClient)

	timeout := time.Duration(env.ContextTimeout) * time.Second

	// TODO Update with Gin API endpoints (routes), controllers, and middleware
	// TODO update domains, repositories, and usecases with Gin framework
}
