package main

import (
  "live-resume/pkg/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
  router:= gin.New()
  router.SetTrustedProxies([]string{"localhost"})
  router.Use(gin.Logger())
  router.Use(cors.Default())

  // ENDPOINTS - C.R.U.D
  // USER
  routes.UserRoutes(router)
  // ITEM
  routes.ItemRoutes(router)

  // Run Serer
  router.Run("localhost:8880")
}

