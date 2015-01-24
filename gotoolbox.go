package main

import (
  "fmt"
  "net/http"
  // "github.com/codegangsta/controller"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  // "gotoolbox/controllers"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  //controller.Action((*controllers.ApplicationController).Index))

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
