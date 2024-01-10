package database

import (
  "github.com/glebarez/sqlite"
  "gorm.io/gorm"
  "fmt"
  "myapp/models"
)

var db *gorm.DB

func init() {
  var err error
  db, err = gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
  if err != nil || db == nil {
    panic("failed to connect database")
  }
  fmt.Println("connected to db")
  db.AutoMigrate(&models.User{})
  fmt.Println("fields of db is created")
}


func GetDB() *gorm.DB {
  return db
}