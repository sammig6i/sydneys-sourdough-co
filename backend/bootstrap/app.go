package bootstrap

import (
	"github.com/sammig6i/sydneys-sourdough-co/database"
	"github.com/sammig6i/sydneys-sourdough-co/pkg/embedding"
)

type Application struct {
	Env             *Env
	Postgres        database.Database
	EmbeddingClient *embedding.Client
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDB(app.Env)
	app.EmbeddingClient = embedding.NewClient(app.Env.EmbeddingServiceURL)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}
