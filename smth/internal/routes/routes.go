package routes

import (
  "proj/internal/app"
  "github.com/go-chi/chi/v5"   
)

func SetupRoutes(app *app.Application) *chi.Mux {
  r := chi.NewRouter()

  r.Get("/health", app.HeatlhCheck)
  r.Get("/workouts/{id}", app.WorkoutHandler.HandleGetWorkoutById)

  r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
  return r
}
