package item

import (
	"net/http"
	"strings"

	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"

	"github.com/gin-gonic/gin"
)

func UpdateItem(c *gin.Context) {

	var requestData struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		Category	*string `json:"category"`
		Image		string `json:"image"`
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
	item.Category = requestData.Category

	if len(strings.TrimSpace(requestData.Image)) >0 {
		imagePath, err := utils.SaveBase64Image(requestData.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
			return
		}
		item.Image = imagePath
	}

	if err := config.DB.Save(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save changes"))
		return
	}

	responseData := response(item)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
