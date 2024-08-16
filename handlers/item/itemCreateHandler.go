package item

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var requestData struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Description string `json:"description"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	class := models.Item{
		Name:        requestData.Name,
		Price:       requestData.Price,
		Description: requestData.Description,
	}

	if err := config.DB.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Response
	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
