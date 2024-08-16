package testimoni

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func CreateTestimoni(c *gin.Context) {
	var requestData struct {
		Name        string `json:"name"`
		Testimoni 	string `json:"testimoni"`
		Image       string `json:"image"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Save base64 image to disk
	imagePath, err := utils.SaveBase64Image(requestData.Image)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed save image"})
		return
	}

	// Create teacher object
	testimoni := models.Testimoni{
		Name: requestData.Name,
		Testimoni: requestData.Testimoni,
		Image:       imagePath,
	}

	if err := config.DB.Create(&testimoni).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Response
	responseData := response(testimoni)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
