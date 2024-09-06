package address

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAddress(c *gin.Context) {
	userId,_ := c.Get("userId")
	username,_ := c.Get("username")

	var requestData struct {
		Address  	string      `json:"address"`
		City		string	  	`json:"city"`
		Province	string		`json:"province"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}
	var address models.Address;
	if err := config.DB.Where("user_id = ?", userId).Find(&address).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()));
		return;
	}

	address.UserID = userId.(string)
	address.Address= requestData.Address
	address.City= requestData.City
	address.Province= requestData.Province

	if err := config.DB.Save(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	response := responseConvert(address)
	response.Username = username.(string);
	c.JSON(http.StatusOK, utils.SuccessResponse(response))
}
