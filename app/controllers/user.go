package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllUser(c *gin.Context) {
  var users []User
  var _users []transformedUser

  Db.Find(&users)
  if len(users) <= 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No users found!"})
    return
  }
  for _, item := range users {
    _users = append(_users, transformedUser{ID: item.ID, FirstName: item.FirstName, LastName: item.LastName, Phone: item.Phone, Email: item.Email})
  }
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Yooooo, all user listed ", "data": _users})
}

func GetProfile(c *gin.Context) {
  var user User
  Db.First(&user, c.Param("id"))

  if user.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
    return
  }
  _user := transformedUser{ID: user.ID, FirstName: user.FirstName, LastName: user.LastName, Phone: user.Phone, Email: user.Email}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Yooooo! found user", "data": _user})
}

func UpdateUser(c *gin.Context) {

  var user User
  var u User
  c.BindJSON(&u)

  Db.First(&user, c.Param("id"))

  if user.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
    return
  }
  Db.Model(&user).Updates(map[string]interface{}{"FirstName": u.FirstName, "LastName": u.LastName, "Phone": u.Phone, "Email": u.Email})
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Yooooo! user updated successfully", "data": user})
}

func DeleteUser(c *gin.Context) {
  var user User

  Db.First(&user, c.Param("id"))

  if user.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No user found!"})
    return
  }

  Db.Delete(&user)
  c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "User deleted successfully!"})
}
