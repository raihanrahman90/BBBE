package teacher

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func CreateTeacher(c *gin.Context) {
		var requestData struct {
			Name        string   `json:"name"`
			Username    string   `json:"username"`
			Image       string   `json:"image"`
			Description string   `json:"description"`
			Fields      []string `json:"fields"`
		}

		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
			return
		}

		// Save base64 image to disk
		imagePath, err := utils.SaveBase64Image(requestData.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed Save image"))
			return
		}

		// Create teacher object
		teacher := models.Teacher{
			Name		: requestData.Name,
			Image		: imagePath,
			Description	: requestData.Description,
			Fields		: generateFields(requestData.Fields),
		}

		// Save teacher to database
		if err := config.DB.Create(&teacher).Error; err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse(err.Error()))
			return
		}

		// Response
		c.JSON(http.StatusOK, utils.SuccessResponse(teacher))
}