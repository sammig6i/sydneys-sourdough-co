package bootstrap

import "github.com/sammig6i/sydneys-sourdough-co/database"

type Application struct {
	Env      *Env
	Postgres database.Database
}

func App() Application {
	app := &Application{}
	app.Env = NewEnv()
	app.Postgres = NewPostgresDB(app.Env)
	return *app
}

func (app *Application) CloseDBConnection() {
	ClosePostgresDBConnection(app.Postgres)
}
