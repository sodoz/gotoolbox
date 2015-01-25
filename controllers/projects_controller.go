package controllers

import (
  "fmt"
  "github.com/gophergala/gotoolbox/models"
  . "github.com/gophergala/gotoolbox/services"
  "github.com/gorilla/schema"
)

// @TODO is this one thread safe?
var formDecoder = schema.NewDecoder()

type ProjectsController struct {
  ApplicationController
}

func (controller *ProjectsController) New() error {
  currentUser := models.User{}
  if currentUserId, ok := controller.Session.Values["currentUserId"]; ok {
    DB().First(&currentUser, currentUserId)
    fmt.Println("YO YO YO")
    fmt.Println(currentUser.GitHubEmail)
  }

  scope := make(map[string]interface{})
  scope["Msg"] = "Test Test Test"

  if err := controller.Render("views/projects/new", scope); err != nil {
    return err
  }

  return nil
}

func (controller *ProjectsController) Create() error {
  if err := controller.Request.ParseForm(); err != nil {
    return err
  }

  project := new(models.Project)
  if err := formDecoder.Decode(project, controller.Request.PostForm); err != nil {
    return err
  }

  DB().Save(project)

  controller.Redirect("/", 200)
  return nil
}
