package class

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func CreateClass(c *gin.Context) {
	var requestData struct {
		Title       string 		`json:"title"`
		Price       int    		`json:"price"`
		Description string 		`json:"description"`
		Image       string 		`json:"image"`
		Modules 	[]string	`json:"modules"`
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
	class := models.Class{
		Title:       requestData.Title,
		Price:       requestData.Price,
		Description: requestData.Description,
		Image:       imagePath,
		Modules: 	 generateModules(requestData.Modules),
	}

	if err := config.DB.Create(&class).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// Response
	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
