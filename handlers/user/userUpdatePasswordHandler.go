package user

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func UserUpdatePassword(c *gin.Context) {
	var requestData struct {
		Username	string	`json:"username"`
		Password 	string `json:"password"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to encrypt password"))
		return
	}

	var user models.User

	if err := config.DB.Where("username=?", requestData.Username).First(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	user.Password = string(hashedPassword)

	if err:=config.DB.Save(&user).Error;err!=nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to update password"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}