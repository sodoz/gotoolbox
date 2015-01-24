package main

import (
  "fmt"
  "github.com/codegangsta/controller"
  "github.com/codegangsta/negroni"
  "github.com/gophergala/gotoolbox/controllers"
  "github.com/gorilla/mux"
  "os"
)

func main() {
  wd, err := os.Getwd()
  fmt.Println("err:", err)
  fmt.Println("wd:", wd)

  r := mux.NewRouter()
  r.Handle("/", controller.Action((*controllers.ApplicationController).Index))

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
