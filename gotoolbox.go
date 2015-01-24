package main

import (
  "github.com/codegangsta/controller"
  "github.com/codegangsta/negroni"
  "github.com/gophergala/gotoolbox/controllers"
  "github.com/gorilla/mux"
)

func main() {
  r := mux.NewRouter()
  r.Handle("/", controller.Action((*controllers.ApplicationController).Index))

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
