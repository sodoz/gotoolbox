package controllers

import (
  "github.com/codegangsta/controller"
  "github.com/gophergala/gotoolbox/helpers"
  "github.com/gophergala/gotoolbox/models"
  . "github.com/gophergala/gotoolbox/services"
  "github.com/gorilla/sessions"
  "github.com/yosssi/ace"
  "net/http"
)

var AppSecret = "XXHJHjhjFJKLfdsdfkjsdfkljljlfsjkKLFDJFHUFHg77GFgg77g72gghjfa8z94qbk"
var sessionStore sessions.Store = sessions.NewCookieStore([]byte(AppSecret))

type ApplicationController struct {
  controller.Base
  Session *sessions.Session
}

func (controller *ApplicationController) GetCurrentUser() *models.User {
  currentUser := new(models.User)
  if currentUserId, ok := controller.Session.Values["currentUserId"]; ok {
    DB().First(currentUser, currentUserId)
    // @TODO Whyyyyyyyyy???!?!?!?!?!?
    if currentUser.Id == 0 {
      return nil
    } else {
      return currentUser
    }
  } else {
    return nil
  }
}

func (controller *ApplicationController) Init(rw http.ResponseWriter, r *http.Request) error {
  var err error
  controller.Session, err = sessionStore.Get(r, "gotoolbox")
  if err != nil {
    controller.Base.Init(rw, r)
    return err
  }
  return controller.Base.Init(rw, r)
}

func (controller *ApplicationController) saveSession() error {
  if err := controller.Session.Save(controller.Request, controller.ResponseWriter); err != nil {
    return err
  } else {
    return nil
  }
}

func (controller *ApplicationController) Render(template string, scope map[string]interface{}) error {
  return controller.RenderWithLayout(template, "views/layout", scope)
}

func (controller *ApplicationController) RenderWithLayout(template string, layout string, scope map[string]interface{}) error {
  // better put this into action handler
  controller.saveSession()

  tpl, err := ace.Load(layout, template, &ace.Options{
    DelimLeft:  "<%",
    DelimRight: "%>",
    FuncMap:    helpers.FuncMap})
  if err != nil {
    return err
  }
  if err := tpl.Execute(controller.ResponseWriter, scope); err != nil {
    return err
  }

  return nil
}

func (controller *ApplicationController) Redirect(url string, status int) {
  // better put this into action handler
  controller.saveSession()

  http.Redirect(controller.ResponseWriter, controller.Request, url, http.StatusFound)
}
