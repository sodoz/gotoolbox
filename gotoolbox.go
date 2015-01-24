package main

import (
  "github.com/codegangsta/controller"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "gotoolbox/controllers"
)

func main() {
  r := mux.NewRouter()
  r.Handle("/", controller.Action((*controllers.ApplicationController).Index))

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
