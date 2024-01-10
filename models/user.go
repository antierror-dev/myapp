package models

import (
  "time"
)

type User struct {
  ID       uint `gorm:"prymaryKey" json:"id"`
  Fname    string `json:"first_name"`
  Lname    string `json:"last_name"`
  Email    string `json:"email" gorm:"unique;not null;"`
  Phone    string `json:"-" gorm:"unique;default:null;"`
  Password string `json:"-"`
  Token    string `json:"-"`
  Expires time.Time `json:"expires"`
}