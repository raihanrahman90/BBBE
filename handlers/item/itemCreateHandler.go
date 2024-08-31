package item

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateItem(c *gin.Context) {
	var requestData struct {
		Name        string `json:"name"`
		Price       int    `json:"price"`
		Description string `json:"description"`
		Image       string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	imagePath, err := utils.SaveBase64Image(requestData.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
		return
	}

	class := models.Item{
		Name:        requestData.Name,
		Price:       requestData.Price,
		Description: requestData.Description,
		Image:       imagePath,
	}

	if err := config.DB.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Response
	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
