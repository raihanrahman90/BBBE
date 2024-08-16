package user

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	var user []models.User
	if err := config.DB.Find(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
	return ;
}