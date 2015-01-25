package controllers

import (
  "github.com/gophergala/gotoolbox/models"
  . "github.com/gophergala/gotoolbox/services"
)

type ProjectsController struct {
  ApplicationController
}

func (controller *ProjectsController) New() error {
  scope := make(map[string]interface{})
  scope["Msg"] = "Test Test Test"

  if err := controller.Render("views/projects/new", scope); err != nil {
    return err
  }

  return nil
}

func (controller *ProjectsController) Create() error {
  project := models.Project{Name: controller.Request.FormValue("project_name"),
    Description: controller.Request.FormValue("project_description"),
    Link:        controller.Request.FormValue("project_link")}

  DB().Save(&project)

  controller.Redirect("/", 200)
  return nil
}
