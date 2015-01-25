package controllers

import (
  "errors"
  "fmt"
  "github.com/gophergala/gotoolbox/models"
  "github.com/gorilla/mux"
  "github.com/markbates/goth/gothic"
  "net/http"
  "reflect"
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
  gothic.GetProviderName = GetProviderNameForMux

  u, err := gothic.CompleteUserAuth(controller.ResponseWriter, controller.Request)
  if err != nil {
    fmt.Fprintln(controller.ResponseWriter, err)
    return err
  }

  user := models.User{GitHubEmail: u.Email, GitHubName: u.Name, GitHubID: u.UserID, GitHubAvatarURL: u.AvatarURL, GitHubAccessToken: u.AccessToken}

  fmt.Fprintln(controller.ResponseWriter, "Found a user")
  fmt.Fprintln(controller.ResponseWriter, user)
  fmt.Fprintln(controller.ResponseWriter, u)
  fmt.Fprintln(controller.ResponseWriter, reflect.TypeOf(u))
  return nil
}
