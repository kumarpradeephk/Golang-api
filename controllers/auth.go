package controllers

import(
  "net/http"
  "github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Yooooo! Signup success"})
}

func LogIn(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! Login is dialema"})
}

func LogOut(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! You will be always logged out"})
}