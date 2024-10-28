package main

import (
	"log"
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

	menuRepo := repository.NewMenuItemRepository(db)
	categoryRepo := repository.NewCategoryRepository(db)
	searchRepo := repository.NewSearchRepository(db)

	_ = menuRepo
	_ = categoryRepo
	_ = searchRepo
	_ = timeout

	log.Println("Repositories initialized")
	log.Println("Backend service ready")

	// TODO Setup Routes in api/routes/ & Update UseCase and Repository with Gin Framework, then controllers, and middleware for admin dashboard
}
