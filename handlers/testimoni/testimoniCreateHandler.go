package testimoni

import (
	"bbbe/config"
	"bbbe/models"
	"bbbe/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTestimoni(c *gin.Context) {
	userId, _ := c.Get("userId")
	var requestData struct {
		Name        string `json:"name"`
		Testimoni 	string `json:"testimoni"`
		Image       string `json:"image"`
		ItemID		string `json:"itemId"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

		// Create teacher object
	testimoni := models.Testimoni{
		Name: requestData.Name,
		Testimoni: requestData.Testimoni,
		ItemID: requestData.ItemID,
		UserID:    userId.(string),
	}


	if err := config.DB.Create(&testimoni).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Response
	responseData := response(testimoni)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
