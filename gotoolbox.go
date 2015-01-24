package main

import (
  "fmt"
  "github.com/codegangsta/negroni"
  "github.com/gorilla/mux"
  "net/http"
)

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "Welcome to the home page!")
  })

  n := negroni.Classic()
  n.UseHandler(r)
  n.Run(":8080")
}
