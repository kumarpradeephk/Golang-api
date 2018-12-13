package config

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDb() *gorm.DB{
  var db *gorm.DB
  var err error
  db, err = gorm.Open("mysql", "root:@/firstAppDb?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic("failed to connect database")
  }
  return db 
}
