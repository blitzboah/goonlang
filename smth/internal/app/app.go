package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"proj/internal/api"
	"proj/internal/store"
	"proj/migrations"
)

type Application struct{
  Logger *log.Logger
  WorkoutHandler *api.WorkoutHandler
  DB *sql.DB
}

func NewApplication() (*Application , error){
  pgDB, err := store.Open()
  if err != nil {
    return nil, err
  }

  err = store.MigrateFS(pgDB, migrations.FS, ".")
  if err != nil {
    panic(err)
  }

  logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

  // stores go here

  // handlers go here

  worktoutHandler := api.NewWorkoutHandler()

  app := &Application{
    Logger : logger,
    WorkoutHandler : worktoutHandler,
    DB : pgDB,
  }

  return app, nil
}


func (a *Application) HeatlhCheck(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "status is 200: OK")
}
