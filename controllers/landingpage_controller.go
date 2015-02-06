package controllers

import (
  "github.com/gophergala/gotoolbox/models"
  . "github.com/gophergala/gotoolbox/services"
)

type LandingpageController struct {
  ApplicationController
}

func (controller *LandingpageController) Index() error {
  var categories []models.Category
  DB().Order("name asc").Find(&categories)

  var projects []models.Project
  DB().Order("created_at desc").Find(&projects)
  for index, _ := range projects {
    var category models.Category
    DB().Model(&projects[index]).Related(&category, "CategoryId")
    projects[index].Category = category

    var user models.User
    DB().Model(&projects[index]).Related(&user, "UserId")
    projects[index].User = user
  }

  scope := make(map[string]interface{})
  currentUser := controller.GetCurrentUser()
  if currentUser != nil {
    scope["CurrentUser"] = currentUser
  }

  scope["Categories"] = categories
  scope["Projects"] = projects

  if err := controller.RenderWithLayout("views/landingpage/index", "views/landingpage/layout", scope); err != nil {
    return err
  }

  return nil
}
