// handlers/teacher.go
package class

import (
	"net/http"

	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func UpdateClass(c *gin.Context) {
	
	var requestData struct {
		Title       string 	`json:"title"`
		Price       int    	`json:"price"`
		Description string 	`json:"description"`
		Image       string 	`json:"image"`
		Modules 	[]string`json:"modules"`
	}
	var class models.Class
	id := c.Param("id")
	if err := config.DB.Where("id=?", id).First(&class).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, utils.FailedResponse(err.Error()))
		return
	}

	// Update fields
	class.Title = requestData.Title
	class.Price = requestData.Price
	class.Description = requestData.Description

	// If image is provided, update it
	if requestData.Image != "" {
		imagePath, err := utils.SaveBase64Image(requestData.Image)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save image"))
			return
		}
		class.Image = imagePath
	}

	class.Modules = updateAllModules(id, requestData.Modules)
	if err := config.DB.Save(&class).Error; err != nil{
		c.JSON(http.StatusInternalServerError, utils.FailedResponse("Failed to save changes"))
		return
	}

	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}
