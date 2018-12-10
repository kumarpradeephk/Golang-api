package main

import (
  "myworkspace/firstApp/controllers"
  "myworkspace/firstApp/database"
  "github.com/gin-gonic/gin"
)

func main() {
  database.Init()
  router := gin.Default()

  v1 := router.Group("/api/v1/users")
  {
    v1.GET("/", controllers.GetAllUser)
    v1.GET("/:id", controllers.GetProfile)
    v1.PUT("/:id", controllers.UpdateUser)
  }

  v2 := router.Group("/api/v1/auth")
  {
    v2.POST("signup/", controllers.SignUp)
    v2.POST("login/", controllers.LogIn)
    v2.DELETE("logout/", controllers.LogOut)
  }

  router.Run()
}