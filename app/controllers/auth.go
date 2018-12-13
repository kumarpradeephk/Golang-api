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
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Yooooo! Signup success", "data": user})
}

func LogIn(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! Login is dialema"})
}

func LogOut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! You will be always logged out"})
}
