package address

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetListAddress(c *gin.Context) {
	var address models.Address
	userId,_ := c.Get("userId")
	username,_ := c.Get("username")
	query := config.DB.Model(&models.Address{})
	query = query.Where("user_id = ?", userId).Preload("User")

	if err := query.First(&address).Error; err != nil {
		c.JSON(http.StatusOK, utils.SuccessResponse(gin.H{"username":username}))
		return
	}
	response := responseConvert(address);
	c.JSON(http.StatusOK, utils.SuccessResponse(response))

}
