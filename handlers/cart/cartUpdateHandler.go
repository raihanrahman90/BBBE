package cart

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateCart(c *gin.Context) {
	id := c.Param("id")
	var requestData struct {
		Amount int `json:"amount"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse("Bad Request"))
		return
	}

	var cart models.Cart
	if err:= config.DB.Where("id = ?", id).Find(&cart).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	cart.Amount = requestData.Amount

	if err := config.DB.Save(&cart).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save changes"))
		return
	}

	responseData := response(cart)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
