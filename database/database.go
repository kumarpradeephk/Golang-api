package database

import (
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func Init() *gorm.DB {
  //open a db connection
  var err error
  db, err = gorm.Open("mysql", "root:@/firstAppDb?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    fmt.Println(err)
    // panic("failed to connect database")
  }
  //Migrate the schema
  db.AutoMigrate(&user{})
  return db 
}

type (
  // userModel describes a userModel type
  user struct {
    gorm.Model
    FirstName     string `json:"first_name"`
    LastName      string `json:"last_name"`
    Phone         string `json:"phone"`
    Email         string `json:"email"`
    Otp           string `json:"otp"`
    AuthToken     string `json:"auth_token"`
  }

  // transformedUser represents a formatted user
  transformedUser struct {
    ID            uint   `json:"id"`
    FirstName     string `json:"first_name"`
    LastName      string `json:"last_name"`
    Phone         string `json:"phone"`
    Email         string `json:"email"`
  }
)
