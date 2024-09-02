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
	var requestData struct {
		Address  	string      `json:"address"`
		City		string	  	`json:"city"`
		Province	string		`json:"province"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}

	address := models.Address{
		UserID: userId.(string),
		Address: requestData.Address,
		City: requestData.City,
		Province: requestData.Province,
	}

	if err := config.DB.Create(&address).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, utils.SuccessResponse(address))
}
