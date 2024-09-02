package cart

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMyCount(c *gin.Context) {
	query := config.DB.Model(&models.Cart{})

	userId, _ := c.Get("userId")
	query = query.Where("user_id = ?", userId);

	var totalItems int64
	query.Count(&totalItems)

	c.JSON(http.StatusOK, utils.SuccessResponse(totalItems))

}
