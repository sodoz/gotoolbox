package main

import (
  "fmt"
  "github.com/gophergala/gotoolbox/models"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

func main() {
  fmt.Println("Seeding Database")
  db, _ := gorm.Open("sqlite3", "gotoolbox.db")
  db.LogMode(true)

  db.DB()

  db.DropTableIfExists(&models.User{})
  db.DropTableIfExists(&models.Project{})

  // create database
  db.AutoMigrate(&models.User{})
  db.AutoMigrate(&models.Project{})
}
