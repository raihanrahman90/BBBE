package user

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RefreshToken(c *gin.Context) {
	var userData models.User
	refreshToken := c.DefaultQuery("refreshToken", "")
	if err := config.DB.Where("refresh_token = ?", refreshToken).First(&userData).Error; err!=nil{
		c.JSON(http.StatusNotFound, utils.FailedResponse(err.Error()))
		return
	}
	token, err := utils.GenerateJWT(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Could not generate token"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(gin.H{"token": token, "refreshToken": userData.RefreshToken}))
}