package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestAdminAPI(c *gin.Context) {
	c.JSON(http.StatusOK, "Ok")
}
