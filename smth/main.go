package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"

	"proj/internal/app"
	"proj/internal/routes"
)

func main(){

  var port int
  flag.IntVar(&port, "port", 8080, "go backend server")
  flag.Parse()

  app, err := app.NewApplication()

  if err != nil {
    panic(err)
  }

  defer app.DB.Close()

  r := routes.SetupRoutes(app)

  server := &http.Server{
    Addr : fmt.Sprintf(":%d", port),
    Handler: r,
    IdleTimeout : time.Minute,
    ReadTimeout : 10 * time.Second,
    WriteTimeout: 30 * time.Second,
  }

  app.Logger.Println("app is running yey on port :", port)

  err = server.ListenAndServe()
  
  if err != nil {
    app.Logger.Fatal(err)
  }
}

