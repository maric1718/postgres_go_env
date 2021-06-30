package main

import (
	"airways/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	adminAPI := r.Group("/admin")
	{
		adminAPI.GET("/test", api.TestAdminAPI)
	}

	r.Run(":3000")
}
