package user

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Response struct{
	AccessToken string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func Login(c *gin.Context) {
	var userData models.User
	var input Credentials

	if err := c.ShouldBindBodyWithJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	if err := config.DB.Where("username = ?", input.Username).First(&userData).Error; err != nil{
		c.JSON(http.StatusUnauthorized, utils.FailedResponse("Username not found"))
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, utils.FailedResponse("Incorrect Pasword"))
		return
	}

	token, err := utils.GenerateJWT(userData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Could not generate token"))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(gin.H{"token": token, "refreshToken": userData.RefreshToken}))
}
