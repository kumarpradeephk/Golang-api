package controllers

import(
  "net/http"
  "github.com/gin-gonic/gin"
)

func GetAllUser(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo, oops! Nothing is found "})
}

func GetProfile(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! Profile is empty."})
}

func UpdateUser(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": "Yooooo! There is no user to update"})
}