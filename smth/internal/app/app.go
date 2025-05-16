package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"proj/internal/api"
)

type Application struct{
  Logger *log.Logger
  WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application , error){
  logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

  // stores go here

  // handlers go here

  worktoutHandler := api.NewWorkoutHandler()

  app := &Application{
    Logger : logger,
    WorkoutHandler : worktoutHandler,
  }

  return app, nil
}


func (a *Application) HeatlhCheck(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w, "status is 200: OK")
}
