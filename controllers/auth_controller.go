package controllers

import (
  "errors"
  "fmt"
  "github.com/gorilla/mux"
  "github.com/markbates/goth/gothic"
  "net/http"
)

type AuthController struct {
  ApplicationController
}

func GetProviderNameForMux(request *http.Request) (string, error) {
  vars := mux.Vars(request)
  provider := vars["provider"]

  if provider == "" {
    return provider, errors.New("you must select a provider")
  }
  return provider, nil
}

func (controller *AuthController) Create() error {
  gothic.GetProviderName = controllers.GetProviderNameForMux

  user, err := gothic.CompleteUserAuth(controller.ResponseWriter, controller.Request)
  if err != nil {
    fmt.Fprintln(controller.ResponseWriter, err)
    return err
  }

  fmt.Fprintln(controller.ResponseWriter, "Found a user")
  fmt.Fprintln(controller.ResponseWriter, user)
  return nil
}
