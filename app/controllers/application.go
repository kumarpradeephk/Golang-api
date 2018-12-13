package controllers
import (
  "github.com/jinzhu/gorm"
  "myworkspace/firstApp/app/models"
  "github.com/kjk/betterguid"
)

var Db *gorm.DB
type User models.User
type transformedUser models.TransformedUser

func generateToken() string {
  token := betterguid.New()
  return token
}