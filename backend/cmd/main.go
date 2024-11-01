package main

import (
	"github.com/gin-gonic/gin"
	"github.com/sammig6i/sydneys-sourdough-co/bootstrap"
	"github.com/sammig6i/sydneys-sourdough-co/repository"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	// db := app.Postgres
	defer app.CloseDBConnection()

	repository.InitEmbeddingClient(app.EmbeddingClient)

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	gin.Run(env.BackendPort)

	// TODO Setup Routes in api/routes/ & Update UseCase and Repository with Gin Framework, then controllers, and middleware for admin dashboard
}
