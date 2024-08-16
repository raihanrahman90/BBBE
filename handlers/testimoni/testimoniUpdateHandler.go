// handlers/teacher.go
package testimoni

import (
	"net/http"

	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func UpdateTestimoni(c *gin.Context) {
	
	var requestData struct {
		Name        string `json:"name"`
		Testimoni 	string `json:"testimoni"`
		Image       string `json:"image"`
	}

	var testimoni models.Testimoni
	id := c.Param("id")
	if err := config.DB.Where("id=?", id).First(&testimoni).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	// Update fields
	testimoni.Name = requestData.Name;
	testimoni.Testimoni = requestData.Testimoni;

	// If image is provided, update it
	if requestData.Image != "" {
		imagePath, err := utils.SaveBase64Image(requestData.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save image"))
			return
		}
		testimoni.Image = imagePath
	}

	if err := config.DB.Save(&testimoni).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save changes"))
		return
	}

	responseData := response(testimoni)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
