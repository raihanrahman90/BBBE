package class

import (
	"net/http"
	"rumahbelajar/config"
	"rumahbelajar/models"
	"rumahbelajar/utils"

	"github.com/gin-gonic/gin"
)

func GetClass(c *gin.Context) {
	var class []models.Class
	if err := config.DB.Preload("Modules").Find(&class).Error; err != nil {
		c.JSON(http.StatusNotFound, utils.FailedResponse("Data Not Found"))
		return
	}

	responseData := responseList(class)
	c.JSON(http.StatusOK, utils.SuccessResponse(responseData))
}