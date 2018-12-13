package controllers
import (
  "github.com/jinzhu/gorm"
  "myworkspace/firstApp/app/models"
)

var Db *gorm.DB
type User models.User
type transformedUser models.TransformedUser