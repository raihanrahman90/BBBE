package user

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/enums"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserCreate struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Access enums.Access `json:"access"`
}

func CreateUser(c *gin.Context) {
	var requestData UserCreate
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to encrypt password"))
		return
	}

	user := models.User{
		Username: requestData.Username,
		Password: string(hashedPassword),
		Access: requestData.Access,
		RefreshToken: utils.RandomString(20),
	}

	if err := config.DB.Create(&user).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(user))
}