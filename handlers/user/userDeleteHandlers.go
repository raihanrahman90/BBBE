package user

import (
	"net/http"
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func DeleteUser(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	if err := config.DB.Where("id=?", id).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Id not found"))
		return;
	}

	if err := config.DB.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to delete user"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(true))
}