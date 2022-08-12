package routes

import (
	"live-resume/pkg/controllers"

	"github.com/gin-gonic/gin"
  )

// UserRoutes ... All User routing
func UserRoutes(router *gin.Engine) {
  // Create
  router.POST("/user/create", controllers.AddUser)
  // Read
  router.GET("/user/:id", controllers.GetUser)
  // Update
  router.PUT("/user/update/:id", controllers.UpdateUser)
  // Delete
  router.DELETE("/user/delete/:id", controllers.DeleteUser)
}

// ItemRoutes ... All Item routing
func ItemRoutes(router *gin.Engine) {
  endpoints := [6]string{"basic", "education", "skill", "job", "project", "list"}

  for _, item := range endpoints {
    router.GET("/user/:id/" + item, controllers.GetItem)
  }
}
