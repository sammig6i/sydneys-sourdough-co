package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	route "github.com/sammig6i/sydneys-sourdough-co/api/route"
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

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	port := env.BackendPort
	log.Printf("Listening on port %s", port)
	gin.Run(port)

	// TODO Test menu routes for CRUD and others
	// TODO configure to skip migrations when on local supabase
}

/*
Domain
+----------------------------------------------------------+
|                                                          |
| Controller --> Usecase --> Repository --> DB             |
|                                                          |
+----------------------------------------------------------+
*/
