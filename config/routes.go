package config

import (
  "github.com/gin-gonic/gin"
  "myworkspace/firstApp/app/controllers"
)

func migrateDb() {
  controllers.Db.AutoMigrate(&controllers.User{})
}

func Start() {
  router := gin.Default()
  controllers.Db = ConnectDb()
  migrateDb()
  
  v1 := router.Group("/api/v1/users")
  {
    v1.GET("/", controllers.GetAllUser)
    v1.GET("/:id", controllers.GetProfile)
    v1.PUT("/:id", controllers.UpdateUser)
    v1.DELETE("/:id", controllers.DeleteUser)
  }

  v2 := router.Group("/api/v2/auth")
  {
    v2.POST("signup/", controllers.SignUp)
    v2.POST("login/", controllers.LogIn)
    v2.DELETE("logout/", controllers.LogOut)
  }
  router.Run(":4000")
}