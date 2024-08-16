package class

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func GetClassById(c *gin.Context) {
	id := c.Param("id")

	var class models.Class

	if err := config.DB.Preload("Modules").Where("id = ?", id).First(&class).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := response(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}