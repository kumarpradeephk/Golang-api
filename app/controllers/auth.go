package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func SignUp(c *gin.Context) {
  var u User
  c.BindJSON(&u)
  user := User{FirstName: u.FirstName, LastName: u.LastName, Phone: u.Phone, Email: u.Email}
	Db.Create(&user)
  if user.ID == 0 {
    c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Unable to SignUp"})
    return
  }
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Yooooo! Signup success", "data": user})
}

func LogIn(c *gin.Context) {
  var u, user User
  c.BindJSON(&u)
  Db.Where("phone = ?", u.Phone).First(&user)
  if user.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No user found!"})
    return
  }
  user.AuthToken = generateToken()
  Db.Save(&user)
  c.JSON(http.StatusOK, gin.H{"status": true, "message": "Yooooo! logged in success", "token": user.AuthToken})
}

func LogOut(c *gin.Context) {
  var u, user User
  c.BindJSON(&u)
  Db.Where("auth_token = ?", u.AuthToken).First(&user)
  if user.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": false, "message": "No user found!"})
    return
  }
  user.AuthToken = ""
  Db.Save(&user)
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! successfully logged out"})
}
