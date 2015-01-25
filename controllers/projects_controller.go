package controllers

import (
  "errors"
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
  scope := make(map[string]interface{})

  currentUser := controller.GetCurrentUser()
  if currentUser == nil {
    return errors.New("You need to be authentictad to create a link")
  }

  var categories []models.Category
  DB().Order("name asc").Find(&categories)

  scope["Categories"] = categories
  scope["CurrentUser"] = currentUser

  if err := controller.Render("views/projects/new", scope); err != nil {
    return err
  }

  return nil
}

func (controller *ProjectsController) Create() error {
  currentUser := controller.GetCurrentUser()

  if currentUser == nil {
    return errors.New("You need to be authentictad to create a link")
  }

  if err := controller.Request.ParseForm(); err != nil {
    return err
  }

  project := new(models.Project)
  if err := formDecoder.Decode(project, controller.Request.PostForm); err != nil {
    return err
  }
  project.UserId = currentUser.Id

  DB().Save(project)

  controller.Redirect("/", 200)
  return nil
}
