// handlers/teacher.go
package teacher

import (
	"net/http"

	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

type requestData struct {
	Name   		string   `json:"name"`
	Description string	`json:"description"`
	Image  		string   `json:"image"`
	Fields 		[]string `json:"fields"`
	Grades 		[]string `json:"grades"`
}

func UpdateTeacher(c *gin.Context) {
		var teacher models.Teacher
		var requestData requestData
		id := c.Param("id")
		if err := config.DB.Where("id = ?", id).First(&teacher).Error; err != nil {
			c.JSON(http.StatusNotFound, utils.FailedResponse("Data not found"))
			return
		}

		if err := c.ShouldBindJSON(&requestData); err != nil {
			c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
			return
		}

		// Update fields
		teacher.Name = requestData.Name
		teacher.Description = requestData.Description

		// If image is provided, update it
		if requestData.Image != "" {
			imagePath, err := utils.SaveBase64Image(requestData.Image)
			if err != nil {
				c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed Save Image"))
				return
			}
			teacher.Image = imagePath
		}
		teacher.Fields = updateAllFields(id, requestData.Fields);

		config.DB.Save(&teacher)
		
		c.JSON(http.StatusOK, utils.SuccessResponse(teacher))
}


