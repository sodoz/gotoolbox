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

  // some nice seeding for the categories
  iab := models.Category{Name: "Images and Bitmaps", Description: "Images and Bitmaps"}
  db.Save(&iab)

}
