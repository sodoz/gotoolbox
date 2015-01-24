package controllers

import (
  "github.com/codegangsta/controller"
  "github.com/yosssi/ace"
  "net/http"
)

type ApplicationController struct {
  controller.Base
}

func (controller *ApplicationController) Init(rw http.ResponseWriter, r *http.Request) error {
  return controller.Base.Init(rw, r)
}

func (controller *ApplicationController) Index() error {
  scope := make(map[string]interface{})
  scope["Msg"] = "Test Test Test"

  controller.Render("views/inner", scope)

  return nil
}

func (controller *ApplicationController) Render(template string, scope map[string]interface{}) {
  tpl, err := ace.Load("views/base", template, &ace.Options{
    DelimLeft:  "<%",
    DelimRight: "%>",
  })
  if err != nil {
    http.Error(controller.ResponseWriter, err.Error(), http.StatusInternalServerError)
  }
  if err := tpl.Execute(controller.ResponseWriter, scope); err != nil {
    http.Error(controller.ResponseWriter, err.Error(), http.StatusInternalServerError)
  }
}

func (controller *ApplicationController) Redirect(url string, status int) {
  http.Redirect(controller.ResponseWriter, controller.Request, url, http.StatusFound)
}
