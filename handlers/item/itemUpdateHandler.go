package item

import (
	"net/http"

	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func UpdateItem(c *gin.Context) {

	var requestData struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Description string `json:"description"`
	}
	var item models.Item
	id := c.Param("id")
	if err := config.DB.Where("id=?", id).First(&item).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	// Update fields
	item.Name = requestData.Name
	item.Price = requestData.Price
	item.Description = requestData.Description

	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save changes"))
		return
	}

	responseData := response(item)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
