package main

import (
  "fmt"
  "github.com/gophergala/gotoolbox/models"
  "github.com/jinzhu/gorm"
  _ "github.com/lib/pq"
  "os"
)

func main() {
  fmt.Println("Seeding Database")
  db, _ := gorm.Open("postgres", os.Getenv("GOTOOLBOX_POSTGRES_URL"))
  db.LogMode(true)

  db.DB()

  db.DropTableIfExists(&models.User{})
  db.DropTableIfExists(&models.Project{})
  db.DropTableIfExists(&models.Category{})

  // create database
  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Project{})
  db.AutoMigrate(&models.Category{})

  // some nice seeding for the categories
  router := models.Category{Name: "Router", Description: "net/http router"}
  db.Save(&router)

  webframeworks := models.Category{Name: "Web Frameworks", Description: "Webframeworks available"}
  db.Save(&webframeworks)

  json := models.Category{Name: "JSON", Description: "JSON Parsing"}
  db.Save(&json)
}
