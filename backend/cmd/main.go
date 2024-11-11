package main

import (
	"log"
	"net/http"

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

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World!")
	})
	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	port := env.BackendPort
	log.Printf("Listening on port %s", port)
	router.Run(port)

	/*
		TODO
		- Setup Routes in api/routes/
		- Update UseCase and Repository with Gin Framework
		- Add controllers
		- Add middleware for admin dashboard
	*/

}
