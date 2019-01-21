package supports

import (
	"github.com/dotuancd/ezserve/app/models"
	"github.com/gin-gonic/gin"
)

func GetLoggedInUser(c *gin.Context) models.User {
	return c.MustGet("user").(models.User)
}
