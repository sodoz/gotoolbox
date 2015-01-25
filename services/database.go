package database

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/mattn/go-sqlite3"
)

var db *gorm.DB = nil

func DB() *gorm.DB {
  if db == nil {
    connection, err := gorm.Open("sqlite3", "gotoolbox.db")
    if err != nil {
      panic(fmt.Sprintf("Could not connect to database"))
    }
    connection.LogMode(true)
    db = &connection
  }
  return db
}
