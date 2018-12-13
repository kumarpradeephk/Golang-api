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
  
  r1 := router.Group("/api/v1/users")
  {
    r1.GET("/", controllers.GetAllUser)
    r1.GET("/:id", controllers.GetProfile)
    r1.PUT("/:id", controllers.UpdateUser)
    r1.DELETE("/:id", controllers.DeleteUser)
  }

  r2 := router.Group("/api/v1/auth")
  {
    r2.POST("signup/", controllers.SignUp)
    r2.POST("login/", controllers.LogIn)
    r2.DELETE("logout/", controllers.LogOut)
  }
  router.Run(":4000")
}